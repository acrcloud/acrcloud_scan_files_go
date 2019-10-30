package models

type MediaFile struct {
	Filename     string // the file path
	DurationInMs int    // the duration of this file (MillionSecond)
	Source       string // the source of the file (if a file from the network address)
}
