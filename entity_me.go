package graph

type Me struct {
	Id                string `json:"id"`
	UserPrincipalName string `json:"userPrincipalName"`
	GivenName         string `json:"givenName"`
	Surname           string `json:"surname"`
	DisplayName       string `json:"displayName"`
}
