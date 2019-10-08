package types

import (
	"time"

	"github.com/bungle-suit/json"
)

type datetimeType struct{}

func (datetimeType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(time.Time)
	secs := val.Unix()
	w.WriteNumber(float64(secs))
	return nil
}

func (datetimeType) Unmarshal(r *json.Reader) (interface{}, error) {
	fv, err := r.ReadNumber()
	if err != nil {
		return nil, err
	}

	return time.Unix(int64(fv), 0), nil
}
