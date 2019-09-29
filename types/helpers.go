package types

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// Marshal value to json
func Marshal(p *Parser, ts string, v interface{}) ([]byte, error) {
	t, err := p.Parse(ts)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	encoder := json.NewEncoder(&buf)
	err = t.Marshal(encoder, v)
	return buf.Bytes(), err
}

// Unmarshal value from json
func Unmarshal(p *Parser, ts string, data []byte, v interface{}) error {
	t, err := p.Parse(ts)
	if err != nil {
		return err
	}

	buf := bytes.NewReader(data)
	decoder := json.NewDecoder(buf)
	return t.Unmarshal(decoder, reflect.ValueOf(v))
}
