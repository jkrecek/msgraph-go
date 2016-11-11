package graph

type GraphEvent struct {
	Id                   string        `json:"id,omitempty"`
	CreatedDateTime      graphFlatTime `json:"createdDateTime,omitempty"`
	LastModifiedDateTime graphFlatTime `json:"lastModifiedDateTime,omitempty"`
	ChangeKey            string        `json:"changeKey,omitempty"`
	Subject              string        `json:"subject,omitempty"`
	Body                 graphBody     `json:"body,omitempty"`
	Start                graphTime     `json:"start,omitempty"`
	End                  graphTime     `json:"end,omitempty"`
	Location             graphLocation `json:"location,omitempty"`
	// TODO more properties
}
