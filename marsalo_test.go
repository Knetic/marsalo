package marsalo

import (
  "net/http"
  "io/ioutil"
  "bytes"
  "testing"
)

const (
  JSON_CONTENT = `{"name": "Bob", "age": 50, "hometown": "testingTown"}`
  XML_CONTENT = `<dummy><name>Bob</name> <age>50</age> <hometown>testingTown</hometown></dummy>`
)

/*
  Something to parse into.
*/
type Dummy struct {

  Name string `json:"name" xml:"name"`
  Age int `json:"age" xml:"age"`
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

  var response *http.Response
  var dummy Dummy
  var err error

  response = new(http.Response)
  response.Header = make(http.Header)
  response.Header.Set("Content-Type", "application/json")
  response.Body = ioutil.NopCloser(bytes.NewReader([]byte(JSON_CONTENT)))

  err = UnmarshalBody(response, &dummy)
  if(err != nil) {
    test.Logf("Failed to parse response body: %v\n", err)
    test.Fail()
    return
  }
}

/*
  Tests that parsing returns errors for mismatched Body and content types.
*/
func TestMismatchResponseParse(test *testing.T) {


}

func runParser(parser bodyParser, content string, test *testing.T) {

  var dummy Dummy
  var err error

  reader := ioutil.NopCloser(bytes.NewReader([]byte(content)))

  err = parser(reader, &dummy)
  if(err != nil) {
    test.Logf("Error while parsing directly: %v\n", err)
    test.Fail()
    return
  }

  // make sure the values are right
  if(dummy.Name != "Bob" || dummy.Age != 50 || dummy.Hometown != "testingTown") {
    test.Logf("Parsing did not actually decode into the given struct: %v", &dummy)
    test.Fail()
    return
  }
}
