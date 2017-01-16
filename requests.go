package graph

import (
	"fmt"
	"bytes"
	"encoding/json"
	"errors"
)

func (c *Client) GetGeneric(path string) (GenericGraphResult, error) {
	res := make(GenericGraphResult)
	err := c.getRequest(path, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetMe() (*GraphMe, error) {
	res := new(GraphMe)
	err := c.getRequest("me", res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetMeCalendar() ([]*GraphCalendar, error) {
	wrp := new(valueWrapper)
	wrp.Value = new([]*GraphCalendar)
	err := c.getRequest("me/calendars", wrp)
	if err != nil {
		return nil, err
	}

	calendars, ok := wrp.Value.(*[]*GraphCalendar)
	if !ok {
		return nil, errors.New("Graph Calendar request has invalid type")
	}

	return *calendars, nil
}

func (c *Client) GetCalendarEvents(calendarId string) ([]*GraphEvent, error) {
	path := fmt.Sprintf("me/calendars/%s/events", calendarId)
	wrp := new(valueWrapper)
	wrp.Value = new([]*GraphEvent)
	err := c.getRequest(path, wrp)
	if err != nil {
		return nil, err
	}

	events, ok := wrp.Value.(*[]*GraphEvent)
	if !ok {
		return nil, errors.New("Graph Calendar request has invalid type")
	}

	return *events, nil
}

func (c *Client) CreateCalendarEvent(calendarId string, event *GraphEvent) (*GraphEvent, error) {
	path := fmt.Sprintf("me/calendars/%s/events", calendarId)
	bts, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respEvent := new(GraphEvent)
	err = c.doRequest("POST", path, bodyReader, respEvent)
	if err != nil {
		return nil, err
	}

	return respEvent, nil
}

func (c *Client) UpdateCalendarEvent(calendarId string, event *GraphEvent) (*GraphEvent, error) {
	path := fmt.Sprintf("me/calendars/%s/events/%s", calendarId, event.Id)
	bts, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bts)

	respEvent := new(GraphEvent)
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
