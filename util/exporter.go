package util

import (
	"acrcloud-scan-tool-golang/logger"
	"acrcloud-scan-tool-golang/models"
	"fmt"
	"github.com/jszwec/csvutil"
	"os"
	"path/filepath"
	"time"
)

func ExportToCsv(filename string, records []models.Result) {

	dir, fileBaseName := filepath.Split(filename)

	if fileBaseName == "" {
		ct := time.Now().Format("20060102")

		filename = filename+fmt.Sprintf("acrcloud_report_%s.csv", ct)
	}

	err := os.MkdirAll(dir, os.ModePerm)

	b, err := csvutil.Marshal(records)

	if err != nil {
		logger.LogError(componentName, "Export to csv error ", err)
	}

	f, err := os.Create(filename)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // write UTF-8 BOM, to compatible with Microsoft Excel
	f.Write(b)

	defer f.Sync()

	logger.LogInfo(componentName, "Export the report to "+f.Name())
}
