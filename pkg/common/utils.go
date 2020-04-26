package common

import (
	"bytes"
	"encoding/json"
	"io"
)

func MarshalJsonToIOReader(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}
