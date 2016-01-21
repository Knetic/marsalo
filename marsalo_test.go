package marsalo

import (
  "io/ioutil"
  "bytes"
  "testing"
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

  content := `{"name": "Bob", "age": 50, "hometown": "testingTown"}`
  runParser(parseJSON, content, test)
}

/*
  Tests that xml parsing function *actually* decodes XML.
*/
func TestXMLParse(test *testing.T) {

  content := `<dummy><name>Bob</name> <age>50</age> <hometown>testingTown</hometown></dummy>`
  runParser(parseXML, content, test)
}

/*
  Tests that http responses' mime types are used to properly switch their parsers
*/
func TestResponseParse(test *testing.T) {

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
