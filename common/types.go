package common

import "time"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type Code int

const (
	OK   Code = iota
	Fail Code = iota
)

type GuidData struct {
	ID     uint64    `json:"id"`
	WorkId uint64    `json:"workId"`
	Base32 string    `json:"base32"`
	Time   time.Time `json:"time"`
}
