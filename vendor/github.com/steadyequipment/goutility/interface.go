package goutility

import (
	"fmt"

	"reflect"

	"encoding/json"
	"io/ioutil"
)

// SprintfObject print information about an object to a string, optionally include its contents
func SprintfObject(i interface{}, contents bool) (result string) {
	result = SprintfObjectInstance(i)

	if contents == true {
		// TODO: handle error
		contents, _ := SprintfObjectContents(i)

		result = StringAppendWithJoin(result, "\n", contents)
	}

	return
}

// SprintfObjectInstance print information about an object to a string
func SprintfObjectInstance(i interface{}) string {
	return fmt.Sprintf("%s <%p>", reflect.TypeOf(i), i)
}

// SprintfObjectContents print contents of an object to a string
func SprintfObjectContents(i interface{}) (result string, error error) {
	marshalResult, marshalError := json.Marshal(i)

	error = marshalError
	if marshalResult != nil && len(marshalResult) > 0 {
		result = string(marshalResult)
	}

	return
}

// ReadObjectFromJSONFile read an object from a JSON file
func ReadObjectFromJSONFile(object interface{}, fileName string) (error error) {

	fileContents, readFileError := ioutil.ReadFile(fileName)
	if readFileError != nil {

		error = readFileError
	} else {

		marshalError := json.Unmarshal(fileContents, object)
		if marshalError != nil {
			error = marshalError
		}
	}

	return
}

// WriteObjectToJSONFile write an object to a JSON file
func WriteObjectToJSONFile(object interface{}, fileName string, pretty bool) (result error) {

	var fileContents []byte
	var marshalError error

	if pretty == true {
		fileContents, marshalError = json.MarshalIndent(object, "", "    ")
	} else {
		fileContents, marshalError = json.Marshal(object)
	}

	if marshalError != nil {
		result = marshalError
	} else {

		writeError := ioutil.WriteFile(fileName, fileContents, 0644)
		if writeError != nil {
			result = writeError
		}
	}

	return
}
