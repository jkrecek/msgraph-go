package graph_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jkrecek/msgraph-go"
	. "gopkg.in/check.v1"
)

const (
	EVENT_EXAMPLE = `{
	"@odata.context":"https://graph.microsoft.com/v1.0/$metadata#users('easycore.sync.bridge%40gmail.com')/calendars('AQMkADAwATNiZmYAZC0wMTMyLTNiMTAtMDACLTAwCgBGAAADdTs04axG5U6OWbe7-Ph5zAcAuHn7VrZh0EKN1yFgjvJMAwAAAgEGAAAAuHn7VrZh0EKN1yFgjvJMAwAAAhEWAAAA')/events",
	"value":[
		{
			"@odata.etag":"W/\"uHn7VrZh0EKN1yFgjvJMAwAAR9tjfw==\"",
			"id":"AQMkADAwATNiZmYAZC0wMTMyLTNiMTAtMDACLTAwCgBGAAADdTs04axG5U6OWbe7-Ph5zAcAuHn7VrZh0EKN1yFgjvJMAwAAAgENAAAAuHn7VrZh0EKN1yFgjvJMAwAAAEe65O0AAAA=",
			"createdDateTime":"2017-02-27T15:55:26.7970745Z",
			"lastModifiedDateTime":"2017-02-27T15:55:26.8283256Z",
			"changeKey":"uHn7VrZh0EKN1yFgjvJMAwAAR9tjfw==",
			"categories":[
			],
			"originalStartTimeZone":"Central Europe Standard Time",
			"originalEndTimeZone":"Central Europe Standard Time",
			"iCalUId":"040000008200E00074C5B7101A82E0080000000016EDADE91191D201000000000000000010000000B57FB4D117EF6E4E90468724CEB47709",
			"reminderMinutesBeforeStart":15,
			"isReminderOn":true,
			"hasAttachments":false,
			"subject":"Daily meeting",
			"bodyPreview":"",
			"importance":"normal",
			"sensitivity":"normal",
			"isAllDay":false,
			"isCancelled":false,
			"isOrganizer":true,
			"responseRequested":true,
			"seriesMasterId":null,
			"showAs":"busy",
			"type":"seriesMaster",
			"webLink":"https://outlook.live.com/owa/?itemid=AQMkADAwATNiZmYAZC0wMTMyLTNiMTAtMDACLTAwCgBGAAADdTs04axG5U6OWbe7%2FPh5zAcAuHn7VrZh0EKN1yFgjvJMAwAAAgENAAAAuHn7VrZh0EKN1yFgjvJMAwAAAEe65O0AAAA%3D&exvsurl=1&path=/calendar/item",
			"onlineMeetingUrl":null,
			"responseStatus":{
				"response":"organizer",
				"time":"0001-01-01T00:00:00Z"
			},
			"body":{
				"contentType":"html",
				"content":"<html>\r\n<head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\r\n<meta content=\"text/html; charset=us-ascii\">\r\n<style type=\"text/css\" style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n-->\r\n</style>\r\n</head>\r\n<body dir=\"ltr\">\r\n<div id=\"divtagdefaultwrapper\" dir=\"ltr\" style=\"font-size:12pt; color:#000000; font-family:Calibri,Arial,Helvetica,sans-serif\">\r\n<p><br>\r\n</p>\r\n</div>\r\n</body>\r\n</html>\r\n"
			},
			"start":{
				"dateTime":"2017-03-02T16:00:00.0000000",
				"timeZone":"UTC"
			},
			"end":{
				"dateTime":"2017-03-02T16:30:00.0000000",
				"timeZone":"UTC"
			},
			"location":{
				"displayName":"",
				"address":{
				}
			},
			"recurrence":{
				"pattern":{
					"type":"daily",
					"interval":1,
					"month":0,
					"dayOfMonth":0,
					"firstDayOfWeek":"sunday",
					"index":"first"
				},
				"range":{
					"type":"endDate",
					"startDate":"2017-03-02",
					"endDate":"2017-03-16",
					"recurrenceTimeZone":"Central Europe Standard Time",
					"numberOfOccurrences":0
				}
			},
			"attendees":[
			],
			"organizer":{
				"emailAddress":{
					"name":"Easy Core",
					"address":"outlook_3775E5FB23D97756@outlook.com"
				}
			}
		}
	]
}`
)

type event_suite struct{}

var _ = Suite(new(event_suite))

func TestEvent(t *testing.T) { TestingT(t) }

func (s *event_suite) TestUnmarshal(c *C) {
	wrp := new(graph.ValueWrapper)
	var events []*graph.Event
	wrp.Value = &events
	err := json.Unmarshal([]byte(EVENT_EXAMPLE), &wrp)
	c.Assert(err, IsNil)
	c.Assert(len(events), Equals, 1)
	c.Assert(events[0].Subject, Equals, "Daily meeting")
	c.Assert(events[0].Recurrence.Pattern.Type, Equals, graph.DayRecurrenceFrequency)
}

func (s *event_suite) TestMarshalRecurring(c *C) {
	cnt := &graph.Event{
		Subject: "Weekly meeting",
		Start:   graph.NewGraphTime(time.Date(2010, 10, 10, 10, 10, 10, 0, time.UTC)),
		End:     graph.NewGraphTime(time.Date(2010, 10, 10, 11, 10, 10, 0, time.UTC)),
	}

	start, err := cnt.Start.Native()
	c.Assert(err, IsNil)

	end := time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC)
	cnt.Recurrence = graph.NewRecurrence(graph.WeekRecurrenceFrequency, start, &end)

	res, err := json.Marshal(cnt)
	c.Assert(err, IsNil)

	c.Assert(string(res), Equals, `{"subject":"Weekly meeting","body":{"contentType":"","content":""},"start":{"dateTime":"2010-10-10T10:10:10","timeZone":"UTC"},"end":{"dateTime":"2010-10-10T11:10:10","timeZone":"UTC"},"recurrence":{"pattern":{"type":"weekly","interval":1,"daysOfWeek":["sunday"]},"range":{"type":"endDate","startDate":"2010-10-10","endDate":"2020-10-10"}}}`)
}
