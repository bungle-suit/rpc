package types

import (
	"bytes"
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

// Marshal value to json
func Marshal(p *Parser, ts string, v interface{}) ([]byte, error) {
	t, err := p.Parse(ts)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	w := myjson.NewWriter(&buf)
	t.Marshal(w, v)
	err = w.Flush()
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
