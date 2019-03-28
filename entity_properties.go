package graph

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jaytaylor/html2text"
)

const (
	FORMAT_DATETIME = "2006-01-02T15:04:05"
	FORMAT_DATE     = "2006-01-02"
)

// Time provided as flat string from API
type flatTime struct {
	time.Time
}

func NewGraphFlatTime(t time.Time) *flatTime {
	return &flatTime{t}
}

func (t flatTime) MarshalJSON() ([]byte, error) {
	strT := t.Native().UTC().Format(time.RFC3339)
	return json.Marshal(strT)
}

func (t *flatTime) UnmarshalJSON(bts []byte) (err error) {
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

func (b *body) GetText() (string, error) {
	if b.ContentType == "html" {
		return html2text.FromString(b.Content)
	} else {
		return b.Content, nil
	}
}

// Graph date time with timezone provided by API
type timeTimezone struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

func NewGraphTime(t time.Time) timeTimezone {
	dateTime := time.Time(t).Format(FORMAT_DATETIME)
	timeZone := time.Time(t).Location().String()

	return timeTimezone{
		DateTime: dateTime,
		TimeZone: timeZone,
	}
}

func (t *timeTimezone) Native() (time.Time, error) {
	loc, _ := time.LoadLocation(t.TimeZone)
	tm, err := time.ParseInLocation(FORMAT_DATETIME, t.DateTime, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("Error while parsing Graph Time: %s\n", err)
	}

	return tm, nil
}

// Graph location provided by API
type location struct {
	DisplayName string `json:"displayName"`
}

func NewGraphLocation(displayName string) *location {
	return &location{
		DisplayName: displayName,
	}
}

// Often object provided by API are wrapped in value property
type ValueWrapper struct {
	Value    interface{} `json:"value"`
	NextLink string      `json:"@odata.nextLink,omitempty"`
}

type date struct {
	time.Time
}

func NewDate(t time.Time) date {
	return date{t}
}

func (t date) MarshalJSON() ([]byte, error) {
	strT := t.Native().UTC().Format(FORMAT_DATE)
	return json.Marshal(strT)
}

func (t *date) UnmarshalJSON(bts []byte) (err error) {
	var strTime string
	err = json.Unmarshal(bts, &strTime)
	if err != nil {
		return
	}

	t.Time, err = time.Parse(FORMAT_DATE, strTime)
	return
}

func (t date) Native() time.Time {
	return t.Time
}

type RecurrenceFrequency string

const (
	DayRecurrenceFrequency   RecurrenceFrequency = "daily"
	WeekRecurrenceFrequency  RecurrenceFrequency = "weekly"
	MonthRecurrenceFrequency RecurrenceFrequency = "absoluteMonthly"
	YearRecurrenceFrequency  RecurrenceFrequency = "absoluteYearly"
)

type recurrence struct {
	Pattern recurrencePattern `json:"pattern"`
	Range   recurrenceRange   `json:"range"`
}

type recurrencePattern struct {
	Type           RecurrenceFrequency `json:"type"`
	Interval       int                 `json:"interval"`
	Month          int                 `json:"month,omitempty"`
	DayOfMonth     int                 `json:"dayOfMonth,omitempty"`
	DaysOfWeek     []string            `json:"daysOfWeek,omitempty"`
	FirstDayOfWeek string              `json:"firstDayOfWeek,omitempty"`
	Index          string              `json:"index,omitempty"`
}

type recurrenceRange struct {
	Type                string `json:"type"`
	StartDate           date   `json:"startDate,omitempty"`
	EndDate             date   `json:"endDate,omitempty"`
	RecurrenceTimeZone  string `json:"recurrenceTimeZone,omitempty"`
	NumberOfOccurrences int    `json:"numberOfOccurrences,omitempty"`
}

func NewRecurrence(recurrenceFrequency RecurrenceFrequency, startDate time.Time, endDate *time.Time) *recurrence {
	rec := new(recurrence)
	rec.Pattern.Type = recurrenceFrequency
	rec.Pattern.Interval = 1
	rec.Range.StartDate = NewDate(startDate)

	switch recurrenceFrequency {
	case WeekRecurrenceFrequency:
		weekday := startDate.Weekday().String()
		rec.Pattern.DaysOfWeek = []string{strings.ToLower(weekday)}
		break
	case MonthRecurrenceFrequency:
		rec.Pattern.DayOfMonth = startDate.Day()
		break
	case YearRecurrenceFrequency:
		rec.Pattern.DayOfMonth = startDate.Day()
		rec.Pattern.Month = int(startDate.Month())
		break
	}

	if endDate == nil {
		rec.Range.Type = "noEnd"
	} else {
		rec.Range.Type = "endDate"

		rec.Range.EndDate = NewDate(*endDate)
	}

	return rec
}

type attendee struct {
	Type         string         `json:"type"`
	Status       attendeeStatus `json:"status,omitempty"`
	EmailAddress nameAddress    `json:"emailAddress,omitempty"`
}

type attendeeStatus struct {
	Response string `json:"response,omitempty"`
	Time     string `json:"time,omitempty"`
}

type organizer struct {
	EmailAddress nameAddress `json:"emailAddress,omitempty"`
}
