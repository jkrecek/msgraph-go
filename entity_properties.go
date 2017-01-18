package graph

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	FORMAT_NATIVE   = "2006-01-02T15:04:05 MST"
	FORMAT_DATETIME = "2006-01-02T15:04:05"
	FORMAT_TIMEZONE = "MST"
)

// Time provided as flat string from API
type flatTime struct {
	time.Time
}

func NewGraphFlatTime(t time.Time) flatTime {
	return flatTime{t}
}

func (t flatTime) MarshalJSON() ([]byte, error) {
	strT := t.Native().UTC().Format(time.RFC3339)
	return json.Marshal(strT)
}

func (t flatTime) UnmarshalJSON(bts []byte) (err error) {
	var strTime string
	err = json.Unmarshal(bts, &strTime)
	if err != nil {
		return
	}

	t.Time, err = time.Parse(time.RFC3339, strTime)
	return
}

func (t flatTime) Native() time.Time {
	return t.Time
}

// Body of event, provided by API
type body struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

// TODO add support for multiple body types
func NewGraphBody(content string) body {
	return body{
		ContentType: "text",
		Content:     content,
	}
}

// Graph date time with timezone provided by API
type timeTimezone struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

func NewGraphTime(t time.Time) timeTimezone {
	dateTime := time.Time(t).Format(FORMAT_DATETIME)
	timeZone := time.Time(t).Format(FORMAT_TIMEZONE)

	return timeTimezone{
		DateTime: dateTime,
		TimeZone: timeZone,
	}
}

func (t *timeTimezone) Native() (time.Time, error) {
	tm, err := time.Parse(FORMAT_NATIVE, fmt.Sprintf("%s %s", t.DateTime, t.TimeZone))
	if err != nil {
		return time.Time{}, fmt.Errorf("Error while parsing Graph Time: %s\n", err)
	}

	return tm, nil
}

// Graph location provided by API
type location struct {
	DisplayName string `json:"displayName"`
}

func NewGraphLocation(displayName string) location {
	return location{
		DisplayName: displayName,
	}
}

// Often object provided by API are wrapped in value property
type valueWrapper struct {
	Value interface{} `json:"value"`
}
