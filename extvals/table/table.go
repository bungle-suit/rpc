package table

import (
	"fmt"
)

// Table structure data made by rows and columns.
type Table struct {
	cols []*Column
	rows int

	// First row always to be summary, even hasSumRow is false.
	hasSumRow bool
}

// Column metadata of a table column.
type Column struct {
	data []interface{}
	name string
	ts   string
}

// Row to access data by row.
type Row struct {
	cols []*Column
	idx  int
}

// New create empty Table.
func New() *Table {
	return &Table{rows: 1} // use row 0 to store sum row
}

// NumCol returns number of columns.
func (t *Table) NumCols() int {
	return len(t.cols)
}

// NewCol create column and add to table.
//
// ts is type string of rpc type system, nullable type or non-nullable
// types are both okay, table rpc type marshaller handles nil/null value,
// marshal nil cell values to json null, use non-nullable has a little better
// performance.
func (t *Table) NewCol(name string, ts string) *Column {
	result := &Column{name: name, ts: ts, data: []interface{}{nil}}
	t.cols = append(t.cols, result)
	return result
}

// Col returns column at specific position.
func (t *Table) Col(colIdx int) *Column {
	return t.cols[colIdx]
}

// SumRow table summary row, panic if table do not have summary row.
func (t *Table) SumRow() Row {
	if !t.hasSumRow {
		panic(fmt.Errorf("[%s] No summary row", tag))
	}
	return Row{t.cols, 0}
}

// HasSumRow returns true if current table contains summary row.
func (t *Table) HasSumRow() bool {
	return t.hasSumRow
}

// GetOrCreateSumRow return summary row, create it if not exist.
func (t *Table) EnsureSumRow() Row {
	if !t.hasSumRow {
		t.hasSumRow = true
	}
	return Row{t.cols, 0}
}

// RemoveSumRow removes summary row.
func (t *Table) RemoveSumRow() {
	t.hasSumRow = false

	for _, col := range t.cols {
		col.data[0] = nil
	}
}

// NumRows number of rows.
func (t *Table) NumRows() int {
	return t.rows - 1
}

// NewRow create a new row, and append to table.
func (t *Table) NewRow() Row {
	for _, col := range t.cols {
		col.data = append(col.data, nil)
	}

	r := Row{t.cols, t.rows}
	t.rows++
	return r
}

// Row at specific row index, row index start from zero.
func (t *Table) Row(idx int) Row {
	return Row{t.cols, idx + 1}
}

// TypeString of current column
func (c *Column) TypeString() string {
	return c.ts
}

// Name of current column
func (c *Column) Name() string {
	return c.name
}

// cell at specific column index, column index start from zero.
func (r Row) Cell(idx int) interface{} {
	return r.cols[idx].data[r.idx]
}

// Setcell at specific column index, column index start from zero.
func (r Row) SetCell(idx int, val interface{}) {
	r.cols[idx].data[r.idx] = val
}
