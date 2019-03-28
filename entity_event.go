package graph

type Event struct {
	Id                   string       `json:"id,omitempty"`
	CreatedDateTime      *flatTime    `json:"createdDateTime,omitempty"`
	LastModifiedDateTime *flatTime    `json:"lastModifiedDateTime,omitempty"`
	ChangeKey            string       `json:"changeKey,omitempty"`
	Subject              string       `json:"subject,required"`
	Body                 body         `json:"body,omitempty"`
	BodyPreview          string       `json:"bodyPreview,omitempty"`
	Start                timeTimezone `json:"start,required"`
	End                  timeTimezone `json:"end,required"`
	Location             *location    `json:"location,omitempty"`
	Recurrence           *recurrence  `json:"recurrence,omitempty"`
	IsAllDay             bool         `json:"isAllDay"`
	Attendees            []*attendee  `json:"attendees,omitempty"`
	IsOrganizer          bool         `json:"isOrganizer,omitempty"`
	Organizer            organizer    `json:"organizer,omitempty"`
	ShowAs               string       `json:"showAs,omitempty"`
	Importance           string       `json:"importance,omitempty"`
	// TODO more properties
}
