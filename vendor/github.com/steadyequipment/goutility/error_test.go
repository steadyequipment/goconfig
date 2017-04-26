package goutility

import (
	"strings"
	"testing"
)

func testErrorContainsString(t *testing.T, error error, contains string) {
	errorString := error.Error()
	if !strings.Contains(errorString, contains) {
		t.Fatalf("Error's Error string should contain '%s', Error string is '%s'", contains, errorString)
	}
}

func TestMakeError(t *testing.T) {

	testMessage := "This is a test"

	testError := MakeError(testMessage)
	if testError.Message() != testMessage {
		t.Fatalf("Error's Message should be '%s', is '%s'", testMessage, testError.Message())
	}

	if len(testError.Error()) == 0 {
		t.Fatal("Error's Error string should not be empty")
	} else {
		testErrorContainsString(t, testError, "Code:")
		testErrorContainsString(t, testError, "Context:")
		testErrorContainsString(t, testError, testMessage)
	}

	if testError.Code() != UnspecifiedCode {
		t.Fatalf("Error's Code should be '%d', is '%d'", UnspecifiedCode, testError.Code())
	}

	if testError.CodeContext() != UnspecifiedCodeContext {
		t.Fatalf("Error's Code Context should be '%s', is '%s'", UnspecifiedCodeContext, testError.CodeContext())
	}

	if testError.HasValidCode() {
		t.Fatal("Error should have invalid code")
	}
}
