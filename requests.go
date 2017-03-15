package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func (c *Client) GetGeneric(path string) (GenericGraphResult, error) {
	res := make(GenericGraphResult)
	err := c.getRequest(path, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetMe() (*Me, error) {
	res := new(Me)
	err := c.getRequest("me", res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetMeCalendar() ([]*Calendar, error) {
	var calendars []*Calendar

	err := c.readGetIntoFunc("me/calendars", &Calendar{}, func(c interface{}) {
		if cal, ok := c.(*Calendar); ok {
			calendars = append(calendars, cal)
		} else {
			log.Println("Expected Calendar ptr")
		}
	})

	return calendars, err
}

func (c *Client) GetCalendarEvents(calendarId string) ([]*Event, error) {
	var events []*Event

	path := fmt.Sprintf("me/calendars/%s/events", calendarId)
	err := c.readGetIntoFunc(path, &Event{}, func(c interface{}) {
		if ev, ok := c.(*Event); ok {
			events = append(events, ev)
		} else {
			log.Println("Expected Event ptr")
		}
	})

	return events, err
}

func (c *Client) CreateCalendarEvent(calendarId string, event *Event) (*Event, error) {
	path := fmt.Sprintf("me/calendars/%s/events", calendarId)
	bts, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respEvent := new(Event)
	err = c.doRequest("POST", path, bodyReader, respEvent)
	if err != nil {
		return nil, err
	}

	return respEvent, nil
}

func (c *Client) UpdateCalendarEvent(calendarId string, event *Event) (*Event, error) {
	path := fmt.Sprintf("me/calendars/%s/events/%s", calendarId, event.Id)
	bts, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respEvent := new(Event)
	err = c.doRequest("PATCH", path, bodyReader, respEvent)
	if err != nil {
		return nil, err
	}

	return respEvent, nil
}

func (c *Client) DeleteCalendarEvent(calendarId string, eventId string) error {
	path := fmt.Sprintf("me/calendars/%s/events/%s", calendarId, eventId)
	err := c.doRequest("DELETE", path, nil, nil)
	return err
}

func (c *Client) GetDefaultContacts() ([]*Contact, error) {
	return c.GetContacts("me/contacts")
}

func (c *Client) GetContactsInFolder(folderId string) ([]*Contact, error) {
	return c.GetContacts(fmt.Sprintf("me/contactFolders/%s/contacts", folderId))
}

func (c *Client) GetContacts(path string) ([]*Contact, error) {
	var contacts []*Contact

	err := c.readGetIntoFunc(path, &Contact{}, func(c interface{}) {
		if cnt, ok := c.(*Contact); ok {
			cnt.Path = fmt.Sprintf("%s/%s", path, cnt.Id)
			contacts = append(contacts, cnt)
		} else {
			log.Println("Expected Contact ptr")
		}
	})

	return contacts, err
}

func (c *Client) ListContactFolders() ([]*ContactFolder, error) {
	var contactFolders []*ContactFolder
	err := c.readGetIntoFunc("me/contactFolders", &ContactFolder{}, func(c interface{}) {
		if cnt, ok := c.(*ContactFolder); ok {
			contactFolders = append(contactFolders, cnt)
		} else {
			log.Println("Expected ContactFolder ptr")
		}
	})

	return contactFolders, err
}

func (c *Client) CreateContactFolder(folder *ContactFolder) (*ContactFolder, error) {
	bts, err := json.Marshal(folder)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respContactFolder := new(ContactFolder)
	err = c.doRequest("POST", "me/contactFolders", bodyReader, respContactFolder)
	if err != nil {
		return nil, err
	}

	return respContactFolder, nil
}

func (c *Client) CreateDefaultContact(contact *Contact) (*Contact, error) {
	return c.CreateContact("me/contacts", contact)
}

func (c *Client) CreateContactInFolder(folderId string, contact *Contact) (*Contact, error) {
	path := fmt.Sprintf("me/contactFolders/%s/contacts", folderId)
	return c.CreateContact(path, contact)
}

func (c *Client) CreateContact(path string, contact *Contact) (*Contact, error) {
	bts, err := json.Marshal(contact.Out())
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respContact := new(Contact)
	err = c.doRequest("POST", path, bodyReader, respContact)
	if err != nil {
		return nil, err
	}

	respContact.Path = fmt.Sprintf("%s/%s", path, respContact.Id)
	return respContact, nil
}

func (c *Client) UpdateContact(contact *Contact) (*Contact, error) {
	bts, err := json.Marshal(contact.Out())
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respContact := new(Contact)
	err = c.doRequest("PATCH", contact.Path, bodyReader, respContact)
	if err != nil {
		return nil, err
	}

	respContact.Path = contact.Path
	return respContact, nil
}

func (c *Client) DeleteContact(path string) error {
	err := c.doRequest("DELETE", path, nil, nil)
	return err
}
