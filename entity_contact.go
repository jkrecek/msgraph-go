package graph

type Contact struct {
	Id                   string        `json:"id,omitempty"`
	Path                 string        `json:"-"`
	CreatedDateTime      *flatTime     `json:"createdDateTime,omitempty"`
	LastModifiedDateTime *flatTime     `json:"lastModifiedDateTime,omitempty"`
	ChangeKey            string        `json:"changeKey,omitempty"`
	GivenName            string        `json:"givenName,omitempty"`
	Surname              string        `json:"surname,omitempty"`
	EmailAddresses       []nameAddress `json:"emailAddresses,omitempty"`
	HomePhones           []string      `json:"homePhones,omitempty"`
	MobilePhone          string        `json:"mobilePhone,omitempty"`
	BusinessPhones       []string      `json:"businessPhones,omitempty"`
	CompanyName          string        `json:"companyName,omitempty"`
	HomeAddress          *address      `json:"homeAddress,omitempty"`
	BusinessAddress      *address      `json:"businessAddress,omitempty"`
	OtherAddress         *address      `json:"otherAddress,omitempty"`
}

func (c *Contact) Out() *Contact {
	c.CreatedDateTime = nil
	c.LastModifiedDateTime = nil

	return c
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
