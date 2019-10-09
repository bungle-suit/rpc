package types

import (
	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals"
)

type objectType struct {
	*Parser
}

func (o objectType) Marshal(w *json.Writer, val interface{}) error {
	v := val.(extvals.Object)
	valType, err := o.Parse(v.T)
	if err != nil {
		return err
	}

	w.BeginObject()
	w.WriteName(`t`)
	w.WriteString(v.T)
	w.WriteName(`v`)
	if err := valType.Marshal(w, v.V); err != nil {
		return err
	}
	w.EndObject()
	return nil
}

func (o objectType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if err := r.Expect(json.BeginObject); err != nil {
		return nil, err
	}

	if err := r.ExpectName("t"); err != nil {
		return nil, err
	}
	ts, err := r.ReadString()
	if err != nil {
		return nil, err
	}
	valueType, err := o.Parse(ts)
	if err != nil {
		return nil, err
	}

	if err := r.ExpectName("v"); err != nil {
		return nil, err
	}
	result, err := valueType.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	err = r.Expect(json.EndObject)
	return extvals.Object{
		T: ts,
		V: result,
	}, err
}
