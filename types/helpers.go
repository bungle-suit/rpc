package types

import (
	"bytes"
	"reflect"

	"github.com/bungle-suit/json"
)

// Marshal value to json
func Marshal(p *Parser, ts string, v interface{}) ([]byte, error) {
	t, err := p.Parse(ts)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	w := json.NewWriter(&buf)
	if err = t.Marshal(w, v); err != nil {
		return nil, err
	}
	err = w.Flush()
	return buf.Bytes(), err
}

// Unmarshal value from json
func Unmarshal(p *Parser, ts string, data []byte, v interface{}) error {
	t, err := p.Parse(ts)
	if err != nil {
		return err
	}

	r := json.NewReader(data)
	return t.Unmarshal(r, reflect.ValueOf(v))
}
