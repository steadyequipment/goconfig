package goutility

import (
	"testing"
)

func TestGetFirstNonLoopbackIP(t *testing.T) {

	ip, error := GetFirstNonLoopbackIP()
	if (len(ip) > 0 && error != nil) || (len(ip) == 0 && error == nil) {
		t.Fatalf("Should be given invalid ip and error or valid ip and no error: ip(\"%s\") error(\"%s\")", ip, error)
	}
}
