package util

import (
	"acrcloud-scan-tool-golang/acrcloud"
	"acrcloud-scan-tool-golang/logger"
	"acrcloud-scan-tool-golang/models"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ParseResponseToResult(file models.MediaFile, startTime string, endTime string, response models.Response) models.Result {

	var record models.Result

	record.StatusCode = response.Status.Code
	record.Filename = file.Filename
	record.Status = response.Status.Msg

	if response.Status.Code == acrcloud.ACR_ERR_CODE_OK {
		musicInfo := response.Metadata.Music[0]

		record.Source = file.Source
		record.FileDurationInMs = file.DurationInMs
		record.StartTime = startTime
		record.EndTime = endTime
		record.Title = musicInfo.Title
		record.SpotifyId = musicInfo.ExternalMetadata.Spotify.Track.ID
		record.YoutubeId = musicInfo.ExternalMetadata.Youtube.Vid

		if musicInfo.ExternalMetadata.Deezer.Track.ID != nil {
			// this is a hack, cause the api may return a int or a string value
			record.DeezerId = fmt.Sprintf("%v", musicInfo.ExternalMetadata.Deezer.Track.ID)
		}

		record.Upc = musicInfo.ExternalIds.Upc
		record.Isrc = musicInfo.ExternalIds.Isrc
		record.Acrid = musicInfo.Acrid
		record.Artists = ParseArtists(musicInfo.Artists)
		record.AlbumName = musicInfo.Album.Name
		record.Label = musicInfo.Label
		record.PlayedDurationInS = musicInfo.DbEndTimeOffsetMs - musicInfo.DbBeginTimeOffsetMs

	}

	return record
}

func ParseArtists(artists interface{}) string {
	var artistsSlice []string

	switch reflect.TypeOf(artists).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(artists)

		for i := 0; i < s.Len(); i++ {
			artistsSlice = append(artistsSlice, s.Index(i).Field(0).String())
			// using reflect to get the value of the artists
		}
		return strings.Join(artistsSlice, ",")
	}

	return ""
}

// Deserialize the custom file's fields to a map[string]interface{}
func ParseCustomFile(message *json.RawMessage) (map[string]interface{}, error) {

	var m []map[string]interface{}
	var customFields map[string]interface{}

	err := json.Unmarshal(*message, &m)

	if err != nil {
		logger.LogError(componentName, "Custom Json data structure Error", err)
	}

	if len(m) > 0 {

		customFields = m[0] // Only need the first result
		return customFields, nil
	}

	return nil, err

}
