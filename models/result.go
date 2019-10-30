package models

import "fmt"

type Result struct {
	Filename          string `csv:"filename,omitempty"`                     // file path
	StatusCode        int    `csv:"status_code,omitempty"`                  // acrcloud status code
	Status            string `csv:"status,omitempty"`                  // acrcloud status code
	Source            string `csv:"source,omitempty"`                       // the source of the file (if a file from the network address)
	FileDurationInMs  int    `csv:"file_duration_in_millisecond,omitempty"` //被扫描文件的总共的时间 单位：毫秒
	StartTime         string `csv:"start_time,omitempty"`                   //扫描在文件中开始时间
	EndTime           string `csv:"end_time,omitempty"`                     //扫描在文件中结束时间
	PlayedDurationInS int    `csv:"played_duration_in_second,omitempty"`    //播放的时间,单位: 秒
	Title             string `csv:"music_title,omitempty"`                  //歌曲名
	Artists           string `csv:"music_artists,omitempty"`                //歌曲作者
	AlbumName         string `csv:"music_album_name,omitempty"`             //专辑名
	Label             string `csv:"label,omitempty"`                        //唱片公司
	Isrc              string `csv:"isrc,omitempty"`                         //ISRC
	Upc               string `csv:"upc,omitempty"`                          //UPC
	Acrid             string `csv:"acr_id,omitempty"`                       //ACRCloud 的 id
	DeezerId          string `csv:"deezer_id,omitempty"`                    //Deezer 对应的 id
	SpotifyId         string `csv:"spotify_id,omitempty"`                   //Spotify 对应的 id
	YoutubeId         string `csv:"youtube_id,omitempty"`                   //Youtube 对应的 id
	CustomFilesResult
}
type CustomFilesResult struct {
	CustomFilesTitle             string             `csv:"custom_files_title,omitempty"`                     //用户文件的名字
	CustomFilesBucketID          string             `csv:"custom_files_bucket_id,omitempty"`                 //用户文件的 Bucket id
	CustomFilesAcrid             string             `csv:"custom_files_acr_id,omitempty"`                    //用户文件的 ACRCloud 的 id
	CustomFilesPlayedDurationInS int                `csv:"custom_files_played_duration_in_second,omitempty"` //用户文件播放的时间
	CustomFilesUserDefinedFields StringInterfaceMap `csv:",omitempty"`                                       //用户自定义字段
}

type StringInterfaceMap map[string]interface{}

func (sim StringInterfaceMap) MarshalCSV() ([]byte, error) {
	return []byte(fmt.Sprint(sim)), nil
}
