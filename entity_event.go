package graph

type Event struct {
	Id                   string       `json:"id,omitempty"`
	CreatedDateTime      flatTime     `json:"createdDateTime,omitempty"`
	LastModifiedDateTime flatTime     `json:"lastModifiedDateTime,omitempty"`
	ChangeKey            string       `json:"changeKey,omitempty"`
	Subject              string       `json:"subject,omitempty"`
	Body                 body         `json:"body,omitempty"`
	Start                timeTimezone `json:"start,omitempty"`
	End                  timeTimezone `json:"end,omitempty"`
	Location             location     `json:"location,omitempty"`
	// TODO more properties
}
