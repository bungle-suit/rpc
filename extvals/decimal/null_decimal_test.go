package decimal_test

import (
	"encoding/json"
	"testing"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNullDecimalJSON(t *testing.T) {
	d, err := decimal.FromString("3.33")
	assert.NoError(t, err)

	n := decimal.NullDecimal{d, true}
	buf, err := json.Marshal(n)
	assert.NoError(t, err)
	assert.Equal(t, "3.33", string(buf))

	var back decimal.NullDecimal
	assert.NoError(t, json.Unmarshal(buf, &back))
	assert.Equal(t, decimal.NullDecimal{d, true}, back)

	buf, err = json.Marshal(decimal.NullDecimal{})
	assert.NoError(t, err)
	assert.Equal(t, "null", string(buf))

	assert.NoError(t, json.Unmarshal(buf, &back))
	assert.Equal(t, decimal.NullDecimal{}, back)
}
