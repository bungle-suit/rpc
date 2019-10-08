package types

import (
	"bytes"
	"fmt"

	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals/table"
)

type tableType struct {
	*Parser
}

func (tt tableType) Marshal(w *json.Writer, v interface{}) error {
	t := v.(*table.Table)
	colTypes := make([]Type, t.NumCols())
	cols := make([]*table.Column, t.NumCols())
	for i, l := 0, t.NumCols(); i < l; i++ {
		c := t.Col(i)
		cols[i] = c
		var err error
		if colTypes[i], err = tt.Parse(c.TypeString()); err != nil {
			return err
		}
	}

	w.BeginObject()
	if t.NumCols() != 0 {
		w.WriteName("cols")
		w.BeginArray()
		for _, m := range cols {
			w.BeginObject()
			w.WriteName("name")
			w.WriteString(m.Name())
			w.WriteName("type")
			w.WriteString(m.TypeString())
			w.EndObject()
		}
		w.EndArray()
	}
	if t.NumRows() != 0 {
		w.WriteName("rows")
		w.BeginArray()
		for idx := 0; idx < t.NumRows(); idx++ {
			if err := tt.writeRow(w, t.Row(idx), colTypes, false); err != nil {
				return err
			}
		}
		w.EndArray()
	}
	if t.HasSumRow() {
		w.WriteName("sumrow")
		tt.writeRow(w, t.SumRow(), colTypes, true)
	}
	w.EndObject()
	return nil
}

func (tt tableType) writeRow(w *json.Writer, row table.Row, colTypes []Type, isSumRow bool) error {
	w.BeginArray()
	for i, ct := range colTypes {
		v := row.Cell(i)
		if v == nil {
			w.WriteNull()
		} else {
			if err := ct.Marshal(w, row.Cell(i)); err != nil {
				return err
			}
		}
	}
	w.EndArray()
	return nil
}

func (t tableType) Unmarshal(r *json.Reader) (interface{}, error) {
	if err := r.Expect(json.BEGIN_OBJECT); err != nil {
		return nil, err
	}

	result := table.New()
	tt, err := r.Next()
	if err != nil {
		return nil, err
	}

	switch tt {
	case json.END_OBJECT:
		return result, nil

	case json.PROPERTY_NAME:
		if err := t.parseMeta(r, result); err != nil {
			return nil, err
		}
		if err := t.afterMeta(r, result); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, fmt.Errorf("[%s] Unexpected token while unmarshal table", tag)
	}
}

func (t tableType) parseMeta(r *json.Reader, table *table.Table) error {
	if !bytes.Equal(r.Str, []byte("cols")) {
		return fmt.Errorf("[%s] 'cols' should be table type first property", tag)
	}

	if err := r.Expect(json.BEGIN_ARRAY); err != nil {
		return err
	}

	for {
		tt, err := r.Next()
		if err != nil {
			return err
		}

		switch tt {
		case json.END_ARRAY:
			return nil

		case json.BEGIN_OBJECT:
			if err := t.parseColumn(r, table); err != nil {
				return err
			}

		default:
			return fmt.Errorf("[%s] should not happen 2", tag)
		}
	}
}

func (t tableType) parseColumn(r *json.Reader, table *table.Table) error {
	var name, colType string

	for {
		tt, err := r.Next()
		if err != nil {
			return err
		}

		switch tt {
		case json.PROPERTY_NAME:
			switch string(r.Str) {
			case "name":
				var err error
				if name, err = r.ReadString(); err != nil {
					return err
				}
			case "type":
				var err error
				if colType, err = r.ReadString(); err != nil {
					return err
				}
			default:
				return fmt.Errorf("[%s] unexpected table column field: %s", tag, string(r.Str))
			}

		case json.END_OBJECT:
			if name == "" {
				return fmt.Errorf("[%s] `table` meta requires `name` field", tag)
			}
			if colType == "" {
				return fmt.Errorf("[%s] `table` meta requires `type` field", tag)
			}
			table.NewCol(name, colType)
			return nil

		default:
			return fmt.Errorf("[%s] should not happen 1", tag)
		}
	}
}

func (t tableType) afterMeta(r *json.Reader, table *table.Table) error {
	colTypes, err := t.toColumnTypes(table)
	if err != nil {
		return err
	}

	for {
		tt, err := r.Next()
		if err != nil {
			return err
		}

		switch tt {
		case json.END_OBJECT:
			return nil
		case json.PROPERTY_NAME:
			if bytes.Equal(r.Str, []byte("rows")) {
				if err := t.parseRows(r, table, colTypes); err != nil {
					return err
				}
			} else if bytes.Equal(r.Str, []byte("sumrow")) {
				if err := r.Expect(json.BEGIN_ARRAY); err != nil {
					return err
				}
				if err := t.parseRow(r, table.EnsureSumRow(), colTypes); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("[%s] Unexpected field '%s' when restore table", tag, string(r.Str))
			}
		default:
			return fmt.Errorf("[%s] should not happen 3", tag)
		}
	}
}

func (t tableType) toColumnTypes(tbl *table.Table) ([]Type, error) {
	result := make([]Type, tbl.NumCols())
	for i := 0; i < tbl.NumCols(); i++ {
		col := tbl.Col(i)
		ty, err := t.Parse(col.TypeString())
		if err != nil {
			return nil, fmt.Errorf("[%s] Unknown column type: %v,\n%w", tag, col.TypeString(), err)
		}

		result[i] = ty
	}
	return result, nil
}

func (t tableType) parseRows(r *json.Reader, table *table.Table, colTypes []Type) error {
	if err := r.Expect(json.BEGIN_ARRAY); err != nil {
		return err
	}

	for {
		tt, err := r.Next()
		if err != nil {
			return err
		}
		switch tt {
		case json.BEGIN_ARRAY:
			row := table.NewRow()
			if err := t.parseRow(r, row, colTypes); err != nil {
				return err
			}
		case json.END_ARRAY:
			return nil
		default:
			return fmt.Errorf("[%s] should not happen 4", tag)
		}
	}
}

func (t tableType) parseRow(r *json.Reader, row table.Row, colTypes []Type) error {
	for idx, col := range colTypes {
		tt, err := r.Next()
		if err != nil {
			return err
		}
		if tt == json.NULL {
			continue
		} else {
			r.Undo()
		}

		cellVal, err := col.Unmarshal(r)
		if err != nil {
			return err
		} else {
			row.SetCell(idx, cellVal)
		}
	}
	return r.Expect(json.END_ARRAY)
}
