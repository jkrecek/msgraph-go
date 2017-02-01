package graph

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	wrp := new(ValueWrapper)
	wrp.Value = new([]*Calendar)
	err := c.getRequest("me/calendars", wrp)
	if err != nil {
		return nil, err
	}

	calendars, ok := wrp.Value.(*[]*Calendar)
	if !ok {
		return nil, errors.New("GraphCalendar request has invalid type")
	}

	return *calendars, nil
}

func (c *Client) GetCalendarEvents(calendarId string) ([]*Event, error) {
	path := fmt.Sprintf("me/calendars/%s/events", calendarId)
	wrp := new(ValueWrapper)
	wrp.Value = new([]*Event)
	err := c.getRequest(path, wrp)
	if err != nil {
		return nil, err
	}

	events, ok := wrp.Value.(*[]*Event)
	if !ok {
		return nil, errors.New("GraphEvents request has invalid type")
	}

	return *events, nil
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

func (c *Client) GetContacts() ([]*Contact, error) {
	wrp := new(ValueWrapper)
	wrp.Value = new([]*Contact)
	err := c.getRequest("me/contacts", wrp)
	if err != nil {
		return nil, err
	}

	contacts, ok := wrp.Value.(*[]*Contact)
	if !ok {
		return nil, errors.New("GraphContact request has invalid type")
	}

	return *contacts, nil
}

func (c *Client) CreateContact(contact *Contact) (*Contact, error) {
	path := fmt.Sprintf("me/contacts/%s", contact.Id)
	bts, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respContact := new(Contact)
	err = c.doRequest("POST", path, bodyReader, respContact)
	if err != nil {
		return nil, err
	}

	return respContact, nil
}

func (c *Client) UpdateContact(contact *Contact) (*Contact, error) {
	path := fmt.Sprintf("me/contacts/%s", contact.Id)
	bts, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respContact := new(Contact)
	err = c.doRequest("PATCH", path, bodyReader, respContact)
	if err != nil {
		return nil, err
	}

	return respContact, nil
}

func (c *Client) DeleteContact(contactId string) error {
	path := fmt.Sprintf("me/contacts/%s", contactId)
	err := c.doRequest("DELETE", path, nil, nil)
	return err
}
