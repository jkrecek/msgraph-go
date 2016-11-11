package graph

import (
	"encoding/json"
	"fmt"
	"github.com/spkg/bom"
	"io"
)

type jsonGraphError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type jsonGraphWrapper struct {
	Error jsonGraphError `json:"error"`
}

type GraphError struct {
	url string
	err jsonGraphError
}

func newGraphError(url string, rc io.ReadCloser) (*GraphError, error) {
	r := bom.NewReader(rc)
	dec := json.NewDecoder(r)

	wrp := new(jsonGraphWrapper)
	err := dec.Decode(&wrp)
	if err != nil {
		return nil, fmt.Errorf("Graph error parsing went wrong. Error: %s", err)
	}

	return &GraphError{
		url: url,
		err: wrp.Error,
	}, nil
}

func graphError(url string, rc io.ReadCloser) error {
	graphErr, err := newGraphError(url, rc)
	if err != nil {
		return err
	}

	return graphErr
}

func (e *GraphError) Error() string {
	return fmt.Sprintf("Request to url '%s' returned error.\n    Code: %s\n    Message: %s", e.url, e.err.Code, e.err.Message)
}
