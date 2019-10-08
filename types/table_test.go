package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/table"
)

func TestTable(t *testing.T) {
	tbl := table.New()
	assertMarshal(t, "table", tbl, "{}")
	assertUnmarshal(t, "table", "{}", tbl)

	tbl.NewCol("ID", "long")
	tbl.NewCol("Name", "str")
	assertMarshal(t, "table", tbl, `{"cols":[{"name":"ID","type":"long"},{"name":"Name","type":"str"}]}`)
	assertUnmarshal(t, "table", `{"cols":[{"name":"ID","type":"long"},{"name":"Name","type":"str"}]}`, tbl)

	row := tbl.NewRow()
	row.SetCell(0, int64(123))
	row.SetCell(1, "foo")
	row = tbl.NewRow()
	row.SetCell(0, int64(456))
	assertMarshal(t, "table", tbl, `{
		"cols":[{"name":"ID","type":"long"},{"name":"Name","type":"str"}],
		"rows":[[123,"foo"],[456,null]]
	}`)
	assertUnmarshal(t, "table", `{
		"cols":[{"name":"ID","type":"long"},{"name":"Name","type":"str"}],
		"rows":[[123,"foo"],[456,null]]
	}`, tbl)

	row = tbl.EnsureSumRow()
	row.SetCell(0, int64(1000))
	assertMarshal(t, "table", tbl, `{
		"cols":[{"name":"ID","type":"long"},{"name":"Name","type":"str"}],
		"rows":[[123,"foo"],[456,null]],
		"sumrow":[1000,null]
	}`)
	assertUnmarshal(t, "table", `{
		"cols":[{"name":"ID","type":"long"},{"name":"Name","type":"str"}],
		"rows":[[123,"foo"],[456,null]],
		"sumrow":[1000,null]
	}`, tbl)
}
