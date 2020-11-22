package utils

import (
	"encoding/json"
	"io"
)

func ToJSON(intf interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(intf)
}

func FromJSON(intf interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(intf)
}
