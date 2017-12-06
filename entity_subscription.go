package graph

type Subscription struct {
	Id                 string `json:"id,omitempty"`
	ChangeType         string `json:"changeType,omitempty"`
	NotificationUrl    string `json:"notificationUrl,omitempty"`
	Resource           string `json:"resource,omitempty"`
	ExpirationDateTime string `json:"expirationDateTime,omitempty"`
	ClientState        string `json:"clientState,omitempty"`
}
