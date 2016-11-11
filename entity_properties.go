package graph

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	GRAPH_FORMAT_NATIVE   = "2006-01-02T15:04:05 MST"
	GRAPH_FORMAT_DATETIME = "2006-01-02T15:04:05"
	GRAPH_FORMAT_TIMEZONE = "MST"
)

// Time provided as flat string from API
type graphFlatTime struct {
	time.Time
}

func NewGraphFlatTime(t time.Time) graphFlatTime {
	return graphFlatTime{t}
}

func (t graphFlatTime) MarshalJSON() ([]byte, error) {
	strT := t.Native().UTC().Format(time.RFC3339)
	return json.Marshal(strT)
}

func (t graphFlatTime) UnmarshalJSON(bts []byte) (err error) {
	var strTime string
	err = json.Unmarshal(bts, &strTime)
	if err != nil {
		return
	}

	t.Time, err = time.Parse(time.RFC3339, strTime)
	return
}

func (t graphFlatTime) Native() time.Time {
	return t.Time
}

// Body of event, provided by API
type graphBody struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

// TODO add support for multiple body types
func NewGraphBody(content string) graphBody {
	return graphBody{
		ContentType: "text",
		Content:     content,
	}
}

// Graph date time with timezone provided by API
type graphTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

func NewGraphTime(t time.Time) graphTime {
	dateTime := time.Time(t).Format(GRAPH_FORMAT_DATETIME)
	timeZone := time.Time(t).Format(GRAPH_FORMAT_TIMEZONE)

	return graphTime{
		DateTime: dateTime,
		TimeZone: timeZone,
	}
}

func (t *graphTime) Native() (time.Time, error) {
	tm, err := time.Parse(GRAPH_FORMAT_NATIVE, fmt.Sprintf("%s %s", t.DateTime, t.TimeZone))
	if err != nil {
		return time.Time{}, fmt.Errorf("Error while parsing Graph Time: %s\n", err)
	}

	return tm, nil
}

// Graph location provided by API
type graphLocation struct {
	DisplayName string `json:"displayName"`
}

func NewGraphLocation(displayName string) graphLocation {
	return graphLocation{
		DisplayName: displayName,
	}
}

// Often object provided by API are wrapped in value property
type valueWrapper struct {
	Value interface{} `json:"value"`
}
