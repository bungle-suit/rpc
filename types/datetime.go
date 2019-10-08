package types

import (
	"reflect"
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

func (datetimeType) Unmarshal(r *json.Reader, v reflect.Value) error {
	fv, err := r.ReadNumber()
	if err != nil {
		return err
	}

	t := time.Unix(int64(fv), 0)
	v.Elem().Set(reflect.ValueOf(t))
	return nil
}

func (datetimeType) New() interface{} {
	return time.Time{}
}
