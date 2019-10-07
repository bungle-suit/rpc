package table_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/table"
	"github.com/stretchr/testify/assert"
)

type columnPair struct {
	name string
	ts   string
}

func validCols(t *testing.T, tbl *table.Table, expected ...columnPair) {
	assert.Equal(t, len(expected), tbl.NumCols())

	for i, exp := range expected {
		col := tbl.Col(i)
		assert.Equal(t, exp.name, col.Name())
		assert.Equal(t, exp.ts, col.TypeString())
	}
}

func validNewRow(t *testing.T, tbl *table.Table, data ...interface{}) {
	lenBefore := tbl.NumRows()
	r := tbl.NewRow()
	assert.Equal(t, r, tbl.Row(lenBefore), "Expected NewRow() to be the last row")
	assert.Equal(t, lenBefore+1, tbl.NumRows())

	assert.Equal(t, tbl.NumCols(), len(data))
	for i, v := range data {
		assert.Nil(t, r.Cell(i))
		r.SetCell(i, v)
		if v == nil {
			assert.Nil(t, v)
		} else {
			assert.Equal(t, v, r.Cell(i))
		}
	}
}

func TestNoColsTable(t *testing.T) {
	tbl := table.New()
	validCols(t, tbl)
	assert.Equal(t, 0, tbl.NumRows())
	assert.False(t, tbl.HasSumRow())
	assert.Panics(t, func() {
		tbl.SumRow()
	})

	validNewRow(t, tbl)
	validNewRow(t, tbl)

	r := tbl.EnsureSumRow()
	assert.True(t, tbl.HasSumRow())
	assert.Equal(t, r, tbl.SumRow())
	assert.Equal(t, r, tbl.EnsureSumRow())

	tbl.RemoveSumRow()
	assert.False(t, tbl.HasSumRow())
	assert.Panics(t, func() {
		tbl.SumRow()
	})
}

func TestTable(t *testing.T) {
	tbl := table.New()
	tbl.NewCol("a", "int")
	tbl.NewCol(`b`, "long")

	validCols(t, tbl, columnPair{`a`, "int"},
		columnPair{`b`, "long"})

	validNewRow(t, tbl, int32(1), int64(2))
	validNewRow(t, tbl, int32(3), nil)
}
