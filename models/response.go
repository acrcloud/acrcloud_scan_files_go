package models

import "encoding/json"

type Response struct {
	Status struct {
		Msg     string `json:"msg"`
		Code    int    `json:"code"`
		Version string `json:"version"`
	} `json:"status"`
	Metadata struct {
		TimestampUtc string           `json:"timestamp_utc,omitempty"`
		CustomFiles  *json.RawMessage `json:"custom_files,omitempty"`
		Music        []struct {
			ExternalIds struct {
				Isrc string `json:"isrc,omitempty"`
				Upc  string `json:"upc,omitempty"`
			} `json:"external_ids,omitempty"`
			SampleBeginTimeOffsetMs int    `json:"sample_begin_time_offset_ms,omitempty"`
			SampleEndTimeOffsetMs   int    `json:"sample_end_time_offset_ms,omitempty"`
			Label                   string `json:"label,omitempty"`
			DurationMs              int    `json:"duration_ms,omitempty"`
			Acrid                   string `json:"acrid,omitempty"`
			DbBeginTimeOffsetMs     int    `json:"db_begin_time_offset_ms,omitempty"`
			PlayOffsetMs            int    `json:"play_offset_ms,omitempty"`
			ReleaseDate             string `json:"release_date,omitempty"`
			Genres                  []struct {
				Name string `json:"name,omitempty"`
			} `json:"genres,omitempty"`
			Score            int    `json:"score,omitempty"`
			Title            string `json:"title,omitempty"`
			ExternalMetadata struct {
				Youtube struct {
					Vid string `json:"vid,omitempty"`
				} `json:"youtube,omitempty"`
				Spotify struct {
					Album struct {
						Name string `json:"name,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"album,omitempty"`
					Artists []struct {
						Name string `json:"name,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"artists,omitempty"`
					Track struct {
						Name string `json:"name,omitempty"`
						ID   string `json:"id,omitempty"`
					} `json:"track,omitempty"`
				} `json:"spotify,omitempty"`
				Deezer struct {
					Album struct {
						Name string      `json:"name,omitempty"`
						ID   interface{} `json:"id,omitempty"`
					} `json:"album,omitempty"`
					Artists []struct {
						Name string      `json:"name,omitempty"`
						ID   interface{} `json:"id,omitempty"`
					} `json:"artists,omitempty"`
					Track struct {
						Name string      `json:"name,omitempty"`
						ID   interface{} `json:"id,omitempty"`
					} `json:"track,omitempty"`
				} `json:"deezer,omitempty"`
			} `json:"external_metadata,omitempty"`
			Album struct {
				Name string `json:"name,omitempty"`
			} `json:"album,omitempty"`
			DbEndTimeOffsetMs int `json:"db_end_time_offset_ms,omitempty"`
			ResultFrom        int `json:"result_from,omitempty"`
			Artists           []struct {
				Name string `json:"name,omitempty"`
			} `json:"artists,omitempty"`
		} `json:"music,omitempty"`
	} `json:"metadata,omitempty"`
	CostTime   float64 `json:"cost_time,omitempty"`
	ResultType int     `json:"result_type,omitempty"`
}
