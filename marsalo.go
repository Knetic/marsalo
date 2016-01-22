package marsalo

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type bodyParser func(io.ReadCloser, interface{}) error

var parsers map[string]bodyParser

func init() {

	parsers = make(map[string]bodyParser)
	parsers["application/json"] = parseJSON
	parsers["text/json"] = parseJSON
	parsers["text/xml"] = parseXML
}

/*
  Unmarshals the given [response] body from either XML or JSON, populating the given [target]
  with the results.
*/
func UnmarshalRequest(request *http.Request, target interface{}) error {

	return unmarshalStream(request.Body, request.Header.Get("Content-Type"), target)
}

/*
  Unmarshals the given [response] body from either XML or JSON, populating the given [target]
  with the results.
*/
func UnmarshalResponse(response *http.Response, target interface{}) error {

	return unmarshalStream(response.Body, response.Header.Get("Content-Type"), target)
}

func unmarshalStream(readCloser io.ReadCloser, contentType string, target interface{}) error {

	var parser bodyParser
	var found bool

	if contentType == "" {
		return errors.New("Unable to unmarshal - no 'Content-Type' header was specified")
	}

	parser, found = parsers[contentType]
	if !found {
		errorMsg := fmt.Sprintf("Unable to unmarshal - no parser found for MIME type '%s'", contentType)
		return errors.New(errorMsg)
	}

	return parser(readCloser, target)
}

/*
  Parses the given [stream] as JSON, unmarshalling into the given [target].
*/
func parseJSON(stream io.ReadCloser, target interface{}) error {

	var decoder *json.Decoder

	decoder = json.NewDecoder(stream)
	return decoder.Decode(target)
}

/*
  Parses the given [stream] as XML, unmarshalling into the given [target].
*/
func parseXML(stream io.ReadCloser, target interface{}) error {

	var decoder *xml.Decoder

	decoder = xml.NewDecoder(stream)
	return decoder.Decode(target)
}
