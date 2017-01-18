package graph

type Contact struct {
	Id                   string        `json:"id,omitempty"`
	CreatedDateTime      flatTime      `json:"createdDateTime,omitempty"`
	LastModifiedDateTime flatTime      `json:"lastModifiedDateTime,omitempty"`
	ChangeKey            string        `json:"changeKey,omitempty"`
	GivenName            string        `json:"givenName,omitempty"`
	Surname              string        `json:"surname,omitempty"`
	EmailAddresses       []nameAddress `json:"emailAddresses"`
	// TODO more properties
}

type nameAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewNameAddresses(addresses ...string) []nameAddress {
	nameAddresses := make([]nameAddress, len(addresses))
	for i, address := range addresses {
		nameAddresses[i] = nameAddress{address, address}
	}

	return nameAddresses
}
