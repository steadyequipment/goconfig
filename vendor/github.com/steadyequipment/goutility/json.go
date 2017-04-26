package goutility

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func MarshalToJSON(v interface{}) ([]byte, ErrorTypeInterface) {
	rawData, error := json.Marshal(v)
	if error != nil {
		return nil, MakeMashalError(v, error)
	}

	return rawData, nil
}

func MarshalIndentToJSON(v interface{}, prefix string, indent string) ([]byte, ErrorTypeInterface) {
	rawData, error := json.MarshalIndent(v, prefix, indent)
	if error != nil {
		return nil, MakeMashalError(v, error)
	}

	return rawData, nil
}

func UnmarshalFromJSON(data []byte, v interface{}) ErrorTypeInterface {
	error := json.Unmarshal(data, v)
	if error != nil {
		return MakeUnmashalError(v, error)
	}

	return nil
}

// region JSON Safe Time

type JSONSafeTime struct {
	time.Time
}

func (this *JSONSafeTime) Format() string {
	return time.RFC3339Nano
}

func (this *JSONSafeTime) UnmarshalJSON(b []byte) (err error) {
	string := strings.Trim(string(b), "\"")
	if string == "null" {
		this.Time = time.Time{}
		return
	}
	this.Time, err = time.Parse(this.Format(), string)
	return
}

func (this *JSONSafeTime) MarshalJSON() ([]byte, error) {
	if this.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", this.Time.Format(this.Format()))), nil
}

// endregion
