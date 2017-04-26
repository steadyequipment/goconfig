package goutility

import (
	"runtime"
	"testing"
)

func GetStack() []byte {
	result := make([]byte, 50000)
	runtime.Stack(result, true)
	return result
}

func PrintBeforeAndAfterStacks(before []byte, after []byte, t *testing.T) {
	t.Log("-----------------------------------------------")
	t.Log("Before")
	t.Logf("%s", before)

	t.Log("-----------------------------------------------")
	t.Log("After")
	t.Logf("%s", after)
}
