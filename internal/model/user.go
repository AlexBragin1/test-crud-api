package model

type User struct {
	ID            string `json:"id"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Age           int    `json:"age"`
	RecordingDate int64  `json:"recordingdate"`
}
