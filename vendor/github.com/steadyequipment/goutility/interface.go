package goutility

import (
	"fmt"

	"reflect"
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
	marshalResult, marshalError := MarshalToJSON(i)

	error = marshalError
	if marshalResult != nil && len(marshalResult) > 0 {
		result = string(marshalResult)
	}

	return
}

// ReadObjectFromJSONFile read an object from a JSON file
func ReadObjectFromJSONFile(object interface{}, fileName string) ErrorTypeInterface {

	fileContents, readFileError := ReadFile(fileName)
	if readFileError != nil {
		return readFileError
	}

	return UnmarshalFromJSON(fileContents, object)
}

// WriteObjectToJSONFile write an object to a JSON file
func WriteObjectToJSONFile(object interface{}, fileName string, pretty bool) ErrorTypeInterface {

	var fileContents []byte
	var marshalError ErrorTypeInterface

	if pretty == true {
		fileContents, marshalError = MarshalIndentToJSON(object, "", "    ")
	} else {
		fileContents, marshalError = MarshalToJSON(object)
	}

	if marshalError != nil {
		return marshalError
	}

	return WriteFile(fileName, fileContents, 0644)
}
