package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/table"
)

func _TestTable(t *testing.T) {
	tbl := table.New()
	assertMarshal(t, "table", tbl, "{}")
}
