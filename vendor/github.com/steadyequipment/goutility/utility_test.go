package goutility

import (
	"fmt"
	"testing"

	"os"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestEmitOnChanOneLength(t *testing.T) {

	oneLength := make(chan string, 1)
	emitOneLength := EmitOnChan(oneLength)

	if len(oneLength) != 0 {
		t.Fatal("Channel should be empty")
	}

	test := "test"
	error := emitOneLength(test)
	if len(oneLength) != 1 || error != nil {
		t.Fatal("Channel should be length one and unerroring")
	}

	value := <-oneLength
	if value != test {
		t.Fatalf("Channel value was unexpected: %s vs %s", value, test)
	}

	emitOneLength(test)
	error = emitOneLength(test)
	if len(oneLength) != 1 || error != nil {
		t.Fatal("Channel should be length one and unerroring")
	}
}

func TestEmitOnChanTwoLength(t *testing.T) {

	twoLength := make(chan string, 2)
	emitTwoLength := EmitOnChan(twoLength)

	if len(twoLength) != 0 {
		t.Fatal("Channel should be empty")
	}

	test := "test"
	error := emitTwoLength(test)
	if len(twoLength) != 1 || error != nil {
		t.Fatal("Channel should be length one and unerroring")
	}

	value := <-twoLength
	if value != test {
		t.Fatalf("Channel value was unexpected: %s vs %s", value, test)
	}

	emitTwoLength(test)
	error = emitTwoLength(test)
	if len(twoLength) != 2 || error != nil {
		t.Fatal("Channel should be length one and unerroring")
	}

	emitTwoLength(test)
	error = emitTwoLength(test)
	if len(twoLength) != 2 || error != nil {
		t.Fatal("Channel should be length one and unerroring")
	}
}

func TestNewError(t *testing.T) {
	result := NewError("asdf %s", "asdf")
	if fmt.Sprintf("%s", result) != fmt.Sprintf("asdf %s", "asdf") {
		t.Fatal("Unexpected result")
	}

	result = NewError("asdf %s %d", "asdf", 2)
	if fmt.Sprintf("%s", result) != fmt.Sprintf("asdf %s %d", "asdf", 2) {
		t.Fatal("Unexpected result")
	}
}

func TestCurrentExecutable(t *testing.T) {
	result := CurrentExecutable()

	if len(result) == 0 {
		t.Fatalf("CurrentExecutable result should not be empty")
	}
}
