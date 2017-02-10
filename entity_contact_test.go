package graph_test

import (
	"encoding/json"
	"github.com/jkrecek/msgraph-go"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

const (
	CONTACT_EXAMPLE = `{
    "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users('easycore.sync.bridge%40gmail.com')/contacts",
    "value": [
        {
            "@odata.etag": "W/\"EQAAABYAAAC4eftWtmHQQo3XIWCO8kwDAAA2Ss79\"",
            "id": "AQMkADAwATNiZmYAZC0wMTMyLTNiMTAtMDACLTAwCgBGAAADdTs04axG5U6OWbe7-Ph5zAcAuHn7VrZh0EKN1yFgjvJMAwAAAgEOAAAAuHn7VrZh0EKN1yFgjvJMAwAAADY1-JsAAAA=",
            "createdDateTime": "2017-02-01T09:40:30Z",
            "lastModifiedDateTime": "2017-02-01T09:43:06Z",
            "changeKey": "EQAAABYAAAC4eftWtmHQQo3XIWCO8kwDAAA2Ss79",
            "categories": [],
            "parentFolderId": "AQMkADAwATNiZmYAZC0wMTMyLTNiMTAtMDACLTAwCgAuAAADdTs04axG5U6OWbe7-Ph5zAEAuHn7VrZh0EKN1yFgjvJMAwAAAgEOAAAA",
            "birthday": null,
            "fileAs": "Doe, Jane",
            "displayName": "Jane Doe",
            "givenName": "Jane",
            "initials": null,
            "middleName": null,
            "nickName": null,
            "surname": "Doe",
            "title": null,
            "yomiGivenName": null,
            "yomiSurname": null,
            "yomiCompanyName": null,
            "generation": null,
            "imAddresses": [],
            "jobTitle": null,
            "companyName": "CMP ENTERPRISES",
            "department": null,
            "officeLocation": null,
            "profession": null,
            "businessHomePage": null,
            "assistantName": null,
            "manager": null,
            "homePhones": [
                "111 222 444"
            ],
            "mobilePhone": "111 222 333",
            "businessPhones": [
                "111 222 555"
            ],
            "spouseName": null,
            "personalNotes": null,
            "children": [],
            "emailAddresses": [
                {
                    "name": "jane.personal@doe.com",
                    "address": "jane.personal@doe.com"
                },
                {
                    "name": "jane.work@doe.com",
                    "address": "jane.work@doe.com"
                }
            ],
            "homeAddress": {
                "street": "Street 10",
                "city": "Prague",
                "state": "Czech",
                "countryOrRegion": "Czech Republic",
                "postalCode": "11000"
            },
            "businessAddress": {},
            "otherAddress": {}
        }
    ]
}`
)

type suite struct{}

var _ = Suite(new(suite))

func TestContact(t *testing.T) { TestingT(t) }

func (s *suite) TestUnmarshal(c *C) {
	wrp := new(graph.ValueWrapper)
	var contacts []*graph.Contact
	wrp.Value = &contacts
	err := json.Unmarshal([]byte(CONTACT_EXAMPLE), &wrp)
	c.Assert(err, IsNil)
	c.Assert(len(contacts), Equals, 1)
	c.Assert(len(contacts[0].HomePhones), Equals, 1)
	c.Assert(contacts[0].HomePhones[0], Equals, "111 222 444")
	c.Assert(contacts[0].MobilePhone, Equals, "111 222 333")
	c.Assert(len(contacts[0].BusinessPhones), Equals, 1)
	c.Assert(contacts[0].BusinessPhones[0], Equals, "111 222 555")
	c.Assert(contacts[0].HomeAddress.State, Equals, "Czech")
	c.Assert(contacts[0].HomeAddress.CountryOrRegion, Equals, "Czech Republic")
	c.Assert(contacts[0].HomeAddress.PostalCode, Equals, "11000")
	c.Assert(contacts[0].CompanyName, Equals, "CMP ENTERPRISES")
}

func (s *suite) TestMarshal(c *C) {
	cnt := &graph.Contact{
		GivenName:       "Jane",
		Surname:         "Doe",
		EmailAddresses:  graph.NewNameAddresses("jane.doe@example.com"),
		CreatedDateTime: graph.NewGraphFlatTime(time.Now()),
	}

	cnt.Out()

	res, err := json.Marshal(cnt)
	c.Assert(err, IsNil)

	c.Assert(string(res), Equals, `{"givenName":"Jane","surname":"Doe","emailAddresses":[{"name":"jane.doe@example.com","address":"jane.doe@example.com"}]}`)
}

func (s *suite) TestMarshalNoMail(c *C) {
	cnt := &graph.Contact{
		GivenName:       "Jane",
		Surname:         "Doe",
		CreatedDateTime: graph.NewGraphFlatTime(time.Now()),
	}

	cnt.Out()

	res, err := json.Marshal(cnt)
	c.Assert(err, IsNil)

	c.Assert(string(res), Equals, `{"givenName":"Jane","surname":"Doe"}`)
}
