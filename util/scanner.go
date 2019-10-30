package util

import (
	"acrcloud-scan-tool-golang/acrcloud"
	"acrcloud-scan-tool-golang/logger"
	"acrcloud-scan-tool-golang/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const componentName = "util"

func Scan(ctx *cli.Context) error {

	var scanMode, downloadUrl, targetType, filename, output, filters string
	if ctx.IsSet("mode") {
		scanMode = strings.ToLower(ctx.String("mode"))
		logger.LogInfo(componentName, "SCAN MODE: "+scanMode)

	}

	if scanMode == "network" {
		if ctx.IsSet("url") {
			downloadUrl = ctx.String("url")
		} else {

			err := errors.New("URL must be input")
			logger.LogFatal(componentName, "Scan error", err)
		}
	}

	if ctx.IsSet("type") {
		targetType = strings.ToLower(ctx.String("type"))
		logger.LogInfo(componentName, "TARGET TYPE: "+targetType)

	}

	if ctx.IsSet("filename") {
		filename = ctx.String("filename")
		logger.LogInfo(componentName, "FILENAME: "+filename)
	} else {

		err := errors.New("FILENAME must be input")
		logger.LogFatal(componentName, "Scan error", err)
	}

	if ctx.IsSet("output") {
		output = ctx.String("output")
		logger.LogInfo(componentName, "OUTPUT: "+output)
	}

	if ctx.IsSet("filter") {
		filters = strings.ToLower(ctx.String("filter"))
		logger.LogInfo(componentName, "FILTER: "+fmt.Sprintf("%v", filters))

	}

	file := InitFile(scanMode, filename, downloadUrl)

	models.Config.Custom.Report.ReportPath = output

	results := RecognizeFile(file)

	ExportToCsv(models.Config.Custom.Report.ReportPath, results)

	return nil

}

// Recognize a piece (fragment) of media file.
func DoRecognize(file models.MediaFile, startSeconds int, lenSeconds int) (models.Result, error) {

	var response models.Response
	var result models.Result

	fileName := file.Filename

	startTime := secondsToHuman(startSeconds)
	endTime := secondsToHuman(startSeconds + lenSeconds)

	// Call SDK to recognize
	acrResult := models.RecHandler.RecognizeByFile(fileName, startSeconds, lenSeconds, nil)
	// map json to Response struct
	err := json.Unmarshal([]byte(acrResult), &response)

	if err != nil {
		logger.LogError(componentName, "Parse SDK response error", err)
		return result, err
	}

	// map the response to the result struct
	result = ParseResponseToResult(file, startTime, endTime, response)

	msg := []string{fmt.Sprintf("From %s to %s, Status: %s", startTime, endTime, response.Status.Msg)}

	//  add info msg
	if response.Status.Code == acrcloud.ACR_ERR_CODE_OK {

		msg = append(msg, fmt.Sprintf("Music Title: %s", result.Title))
		msg = append(msg, fmt.Sprintf("Artists Name: %s", result.Artists))
		msg = append(msg, fmt.Sprintf("ACRID: %s", result.Acrid))
	}

	// parse custom files fields
	// if can be recognized & config require CustomFile & have custom files fields
	if response.Status.Code == acrcloud.ACR_ERR_CODE_OK && models.Config.Custom.IsCustomFile && response.Metadata.CustomFiles != nil {

		customFile, err := ParseCustomFile(response.Metadata.CustomFiles)
		if err != nil {
			logger.LogError(componentName, "Parse custom file json data error", err)
		}
		msg = append(msg, fmt.Sprintf("Custom File Title: %s", customFile["title"]))
	}

	logger.LogInfo(componentName, strings.Join(msg, ", "))

	return result, err
}

// Recognize a media file
func RecognizeFile(file models.MediaFile) []models.Result {
	logger.LogInfo(componentName, "Start Recognize File: "+file.Filename)

	var recognizeResults []models.Result
	totalDurationInS := file.DurationInMs / 1000 // Conversion of unit
	step := models.Config.Acrcloud.StepInSeconds
	for i := 0; i <= totalDurationInS; i += step {
		res, err := DoRecognize(file, i, models.Config.Acrcloud.RecognizeLength)
		if err != nil {
			logger.LogError(componentName, "Recognize Failed", err)
		}

		recognizeResults = append(recognizeResults, res)
	}
	return recognizeResults
}

// Initialize the file, it will call acrcloud sdk to get the duration of the file and save into the model
// If networkAddress isn't null it will call download function to download the file
func InitFile(mode string, filepath string, networkAddress string) models.MediaFile {

	if mode == "network" {
		if IsUrl(networkAddress) {
			filepath = DownloadFile(networkAddress, filepath)
		} else {
			err := errors.New(fmt.Sprintf("Invalid URL, please check it. %s", networkAddress))
			logger.LogFatal(componentName, "Init File error", err)
		}
	}
	duration, err := models.RecHandler.GetDurationMsByFile(filepath)

	if err != nil {
		logger.LogError(componentName, fmt.Sprintf("Cannot get the file's duration：%s", filepath), err)
	}

	mediaFile := models.MediaFile{Filename: filepath, DurationInMs: duration, Source: networkAddress}

	return mediaFile
}

func IsUrl(str string) bool {

	u, err := url.Parse(str)

	return err == nil && u.Scheme != "" && u.Host != ""
}

func DownloadFile(url string, dest string) string {

	file := path.Base(url)

	logger.LogInfo(componentName, fmt.Sprintf("Downloading file %s from %s", file, url))

	var path bytes.Buffer
	if dest == "" {
		dest = "."
	}
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(file)

	start := time.Now()

	out, err := os.Create(path.String())

	if err != nil {
		err := errors.New(fmt.Sprintf("Create download file error %s  %", path.String()))
		logger.LogError(componentName, "Download error", err)
		fmt.Println(path.String())
		panic(err)
	}

	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		panic(err)
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if err != nil {
		panic(err)
	}

	done := make(chan int64)

	go PrintDownloadPercent(done, path.String(), int64(size))

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	done <- n

	elapsed := time.Since(start)
	log.Printf("Download completed in %s", elapsed)

	return path.String()
}

func PrintDownloadPercent(done chan int64, path string, total int64) {

	var stop = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				logger.LogFatal(componentName, "Download error", err)
			}

			fi, err := file.Stat()
			if err != nil {
				logger.LogFatal(componentName, "Download error", err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent = float64(size) / float64(total) * 100

			logger.LogInfo(componentName, fmt.Sprintf("Download %.0f", percent))
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}
