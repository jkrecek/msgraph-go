package graph

type ContactFolder struct {
	Id             string `json:"id,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	ParentFolderId string `json:"parentFolderId,omitempty"`
}
