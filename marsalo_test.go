package marsalo

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	JSON_CONTENT = `{"name": "Bob", "age": 50, "hometown": "testingTown"}`
	XML_CONTENT  = `<dummy><name>Bob</name> <age>50</age> <hometown>testingTown</hometown></dummy>`
)

/*
  Something to parse into.
*/
type Dummy struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Hometown string `json:"hometown" xml:"hometown"`
}

/*
  Tests that json parsing function *actually* decodes JSON.
*/
func TestJSONParse(test *testing.T) {

	runParser(parseJSON, JSON_CONTENT, test)
}

/*
  Tests that xml parsing function *actually* decodes XML.
*/
func TestXMLParse(test *testing.T) {

	runParser(parseXML, XML_CONTENT, test)
}

/*
  Tests that http responses' mime types are used to properly switch their parsers
*/
func TestResponseParse(test *testing.T) {

	var err error

	err = runResponseUnmarshalParser("application/json", JSON_CONTENT, test)
	if err != nil {
		test.Logf("Failed to parse json body: %v\n", err)
		test.Fail()
		return
	}

	err = runResponseUnmarshalParser("text/json", JSON_CONTENT, test)
	if err != nil {
		test.Logf("Failed to parse json body: %v\n", err)
		test.Fail()
		return
	}

	err = runResponseUnmarshalParser("text/xml", XML_CONTENT, test)
	if err != nil {
		test.Logf("Failed to parse xml body: %v\n", err)
		test.Fail()
		return
	}
}

/*
  Tests that parsing returns errors for mismatched Body and content types.
*/
func TestMismatchResponseParse(test *testing.T) {

	if runResponseUnmarshalParser("text/xml", JSON_CONTENT, test) == nil {
		test.Logf("Failed to correctly refuse to parse json content given as xml")
		test.Fail()
		return
	}

	if runResponseUnmarshalParser("application/json", XML_CONTENT, test) == nil {
		test.Logf("Failed to correctly refuse to parse xml content given as json")
		test.Fail()
		return
	}
}

func runParser(parser bodyParser, content string, test *testing.T) {

	var dummy Dummy
	var err error

	reader := ioutil.NopCloser(bytes.NewReader([]byte(content)))

	err = parser(reader, &dummy)
	if err != nil {
		test.Logf("Error while parsing directly: %v\n", err)
		test.Fail()
		return
	}

	// make sure the values are right
	if !checkDummyValues(dummy) {
		test.Logf("Parsing did not actually decode into the given struct: %v", &dummy)
		test.Fail()
		return
	}
}

func runResponseUnmarshalParser(contentType string, content string, test *testing.T) error {

	var response *http.Response
	var dummy Dummy
	var err error

	response = new(http.Response)
	response.Header = make(http.Header)
	response.Header.Set("Content-Type", contentType)
	response.Body = ioutil.NopCloser(bytes.NewReader([]byte(content)))

	err = UnmarshalResponse(response, &dummy)
	if err != nil {
		return err
	}

	if !checkDummyValues(dummy) {
		errorMsg := fmt.Sprintf("Parsing did not actually decode into the given struct: %v", dummy)
		return errors.New(errorMsg)
	}
	return nil
}

func runRequestUnmarshalParser(contentType string, content string, test *testing.T) error {

	var request *http.Request
	var dummy Dummy
	var err error

	request, _ = http.NewRequest("POST", "http://example.com/marsalo", bytes.NewReader([]byte(content)))
	request.Header = make(http.Header)
	request.Header.Set("Content-Type", contentType)

	err = UnmarshalRequest(request, &dummy)
	if err != nil {
		return err
	}

	if !checkDummyValues(dummy) {
		errorMsg := fmt.Sprintf("Parsing did not actually decode into the given struct: %v", dummy)
		return errors.New(errorMsg)
	}
	return nil
}

func checkDummyValues(dummy Dummy) bool {

	return dummy.Name == "Bob" && dummy.Age == 50 && dummy.Hometown == "testingTown"
}
