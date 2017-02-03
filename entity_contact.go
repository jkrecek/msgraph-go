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

func (c *Contact) AddHomePhone(phone string) {
	if phone != "" {
		c.HomePhones = append(c.HomePhones, phone)
	}
}

func (c *Contact) AddMobilePhone(phone string) {
	if phone != "" {
		c.MobilePhone = phone
	}
}

func (c *Contact) AddBusinessPhone(phone string) {
	if phone != "" {
		c.BusinessPhones = append(c.BusinessPhones, phone)
	}
}

type nameAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewNameAddresses(addresses ...string) []nameAddress {
	nameAddresses := make([]nameAddress, len(addresses))
	for i, address := range addresses {
		if address != "" {
			nameAddresses[i] = nameAddress{address, address}
		}
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
