package graph

type Contact struct {
	Id                   string        `json:"id,omitempty"`
	CreatedDateTime      flatTime      `json:"createdDateTime,omitempty"`
	LastModifiedDateTime flatTime      `json:"lastModifiedDateTime,omitempty"`
	ChangeKey            string        `json:"changeKey,omitempty"`
	GivenName            string        `json:"givenName,omitempty"`
	Surname              string        `json:"surname,omitempty"`
	EmailAddresses       []nameAddress `json:"emailAddresses"`
	HomePhones           []string      `json:"homePhones"`
	MobilePhone          string        `json:"mobilePhone"`
	BusinessPhones       []string      `json:"businessPhones"`
	CompanyName          string        `json:"companyName"`
	HomeAddress          address       `json:"homeAddress"`
	BusinessAddress      address       `json:"businessAddress"`
	OtherAddress         address       `json:"otherAddress"`
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

type address struct {
	Street          string `json:"street"`
	City            string `json:"city"`
	State           string `json:"state"`
	CountryOrRegion string `json:"countryOrRegion"`
	PostalCode      string `json:"postalCode"`
}

func NewAddress(street, city, state, country, postalCode string) address {
	return address{
		Street:          street,
		City:            city,
		State:           state,
		CountryOrRegion: country,
		PostalCode:      postalCode,
	}
}
