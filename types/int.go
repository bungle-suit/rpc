package types

import "encoding/json"

type int32Type struct{}

func (int32Type) FromJson(decoder *json.Decoder) (interface{}, error) {
	panic("not implemented")
}

func (int32Type) ToJson(encoder *json.Encoder, v interface{}) error {
	panic("not implemented")
}

func (int32Type) IsDefault(v interface{}) bool {
	panic("not implemented")
}

func (int32Type) DefaultValue() interface{} {
	panic("not implemented")
}
