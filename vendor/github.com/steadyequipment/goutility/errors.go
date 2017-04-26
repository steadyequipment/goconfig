package goutility

import (
	"fmt"
	"reflect"
)

const (
	CodeContext = "goutility"

	UnableToFindExternalIPAddressErrorCode ErrorCode = 1 << iota

	MarshalErrorCode
	UnmarshalErrorCode
)

// region UnableToFindExternalIPAddressError
type UnableToFindExternalIPAddressError struct {
	ErrorType
}

func MakeUnableToFindExternalIPAddressError() UnableToFindExternalIPAddressError {
	return UnableToFindExternalIPAddressError{
		ErrorType: MakeErrorWithCode("Unable to find external ip address", UnableToFindExternalIPAddressErrorCode, CodeContext),
	}
}

// endregion

// region MashalError
type MashalError struct {
	ErrorType

	object interface{}
	error  error
}

func MakeMashalError(object interface{}, error error) MashalError {
	message := fmt.Sprintf("Error while mashalling an object of type '%s': %s", reflect.TypeOf(object).String(), error.Error())
	return MashalError{
		ErrorType: MakeErrorWithCode(message, MarshalErrorCode, CodeContext),
		object:    object,
		error:     error,
	}
}

// endregion

// region UnmarshalError
type UnmarshalError struct {
	ErrorType

	object interface{}
	error  error
}

func MakeUnmashalError(object interface{}, error error) UnmarshalError {
	message := fmt.Sprintf("Error while unmarshalling an object of type '%s' from data: %s", reflect.TypeOf(object).String(), error.Error())
	return UnmarshalError{
		ErrorType: MakeErrorWithCode(message, UnmarshalErrorCode, CodeContext),
		object:    object,
		error:     error,
	}
}

// endregion

// region ReadFileError
type ReadFileError struct {
	ErrorType

	filePath string
	error    error
}

func MakeReadFileError(filePath string, error error) ReadFileError {
	message := fmt.Sprintf("Error while reading file '%s': %s", filePath, error.Error())
	return ReadFileError{
		ErrorType: MakeErrorWithCode(message, MarshalErrorCode, CodeContext),
		filePath:  filePath,
		error:     error,
	}
}

// endregion

// region WriteFileError
type WriteFileError struct {
	ErrorType

	filePath string
	error    error
}

func MakeWriteFileError(filePath string, error error) WriteFileError {
	message := fmt.Sprintf("Error while writing file '%s': %s", filePath, error.Error())
	return WriteFileError{
		ErrorType: MakeErrorWithCode(message, UnmarshalErrorCode, CodeContext),
		filePath:  filePath,
		error:     error,
	}
}

// endregion
