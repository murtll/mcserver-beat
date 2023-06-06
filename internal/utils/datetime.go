package utils

import (
	"strconv"
	"time"
)

type JSONTime time.Time

func (t JSONTime) String() string {
    return strconv.FormatInt(time.Time(t).UnixMilli(), 10)
}

func (t JSONTime) RoundHour() JSONTime {
    th := time.Time(t).UnixMilli() / time.Hour.Milliseconds()
	th = th * time.Hour.Milliseconds()
	return JSONTime(time.UnixMilli(th))
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
    return []byte(t.String()), nil
}