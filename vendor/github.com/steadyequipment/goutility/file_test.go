package goutility

import (
	"testing"

	"os"
)

func testCreateEmptyFile(location string) error {
	file, createError := os.Create(location)
	if createError != nil {
		return createError
	}
	closeError := file.Close()
	if closeError != nil {
		return closeError
	}
	return nil
}

func testDoCopyFile(source string, createSource bool, destination string, createDestination bool) error {

	if createSource == true {
		testCreateEmptyFile(source)
	} else {

		exists, existsError := FileExists(source)
		if existsError != nil {
			return existsError
		}
		if exists == true {
			removeError := os.Remove(source)
			if removeError != nil {
				return removeError
			}
		}
	}

	if createDestination == true {
		testCreateEmptyFile(destination)
	} else {

		exists, existsError := FileExists(destination)
		if existsError != nil {
			return existsError
		}
		if exists == true {
			removeError := os.Remove(destination)
			if removeError != nil {
				return removeError
			}
		}

	}

	return CopyFile(source, destination)
}

func TestCopyFile(t *testing.T) {

	invalidSource := "this is a garbage file path that cannot exist!@#$%^&*()_+-=,./<>?;'\":][[{}]]"
	invalidDestination := "this is a garbage file path that cannot exist!@#$%^&*()_+-=,./<>?;'\":][[{}]]"
	validSource := "/tmp/goutility_TestCopyFile_source"
	validDestination := "/tmp/goutility_TestCopyFile_destination"

	invalidSourceError := testDoCopyFile(invalidSource, false, validDestination, true)
	if invalidSourceError == nil {
		t.Fatalf("Invalid source '%s' should have resulted in an error", invalidSource)
	}

	invalidDestinationError := testDoCopyFile(validSource, true, invalidDestination, false)
	if invalidDestinationError == nil {
		t.Fatalf("Invalid destination '%s' should have resulted in an error", invalidDestination)
	}

	sourceExists := testDoCopyFile(validSource, true, validDestination, false)
	if sourceExists != nil {
		t.Fatalf("Source exists shouldn't have errored: %s", sourceExists)
	}
	destinationExists := testDoCopyFile(validSource, true, validDestination, true)
	if destinationExists != nil {
		t.Fatalf("Destination exists shouldn't have errored: %s", destinationExists)
	}

	sourceDoesNotExist := testDoCopyFile(validSource, false, validDestination, false)
	if sourceDoesNotExist == nil {
		t.Fatalf("Source does not exist should have errored: %s", sourceDoesNotExist)
	}
}

func TestFileExists(t *testing.T) {

	validFilePath := "/tmp/goutility_TestFileExists"
	createError := testCreateEmptyFile(validFilePath)
	if createError != nil {
		t.Fatalf("Test Error: unable to create file '%s' to test 'FileExists' with: %s", validFilePath, createError)
	}

	validResult, validError := FileExists(validFilePath)
	if validResult == false {
		t.Fatalf("'%s' should have been considered valid", validFilePath)
	}
	if validError != nil {
		t.Fatalf("'%s' should not have cause an error: %s", validFilePath, validError)
	}

	invalidFilePath := "this is a garbage file path that cannot exist!@#$%^&*()_+-=,./<>?;'\":][[{}]]"
	invalidResult, invalidError := FileExists(invalidFilePath)
	if invalidResult == true {
		t.Fatalf("'%s' should not have existed", invalidFilePath)
	}
	if invalidError != nil {
		t.Fatalf("'%s' should not have cause an error: %s", invalidFilePath, invalidError)
	}
}
