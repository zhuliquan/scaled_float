package scaled_float

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaledFloat(t *testing.T) {
	a, err := NewFromString("92233720368547758.07", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, true, a.IsInf(0))
	assert.Equal(t, true, a.IsInf(1))

	a, err = NewFromString("92233720368547758.09", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, true, a.IsInf(0))
	assert.Equal(t, true, a.IsInf(1))

	a, err = NewFromString("92233720368547758.08", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, true, a.IsInf(0))
	assert.Equal(t, true, a.IsInf(1))

	a, err = NewFromString("92233720368547758.06", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, false, a.IsInf(0))
	assert.Equal(t, false, a.IsInf(1))

	a, err = NewFromString("92233720368547758.05", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, false, a.IsInf(0))
	assert.Equal(t, false, a.IsInf(1))

	a, err = NewFromString("-92233720368547758.08", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, true, a.IsInf(0))
	assert.Equal(t, true, a.IsInf(-1))

	a, err = NewFromString("-92233720368547758.10", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, true, a.IsInf(0))
	assert.Equal(t, true, a.IsInf(-1))

	a, err = NewFromString("-92233720368547758.09", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, true, a.IsInf(0))
	assert.Equal(t, true, a.IsInf(-1))

	a, err = NewFromString("-92233720368547758.07", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, false, a.IsInf(0))
	assert.Equal(t, false, a.IsInf(-1))

	a, err = NewFromString("-92233720368547758.06", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, false, a.IsInf(0))
	assert.Equal(t, false, a.IsInf(-1))

	var d1 = New(13.56, 100.0)
	var d2 = New(15.56, 100.0)

	assert.Equal(t, 13.56, d1.RawFloat())
	assert.Equal(t, int64(1356), d1.Int64())
	assert.Equal(t, -1, d1.Compare(d2))
	assert.Equal(t, 1, d2.Compare(d1))
	assert.Equal(t, 0, d1.Compare(d1))

	d3, err := NewFromString("-92233720368547758.07", 100.0)
	assert.Nil(t, err)
	d4, err := NewFromString("-92233720368547758.08", 100.0)
	assert.Nil(t, err)
	d5, err := NewFromString("-92233720368547758.09", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, int64(-9223372036854775807), d3.Int64())
	assert.Equal(t, -92233720368547758.07, d3.RawFloat())
	assert.Equal(t, int64(-9223372036854775808), d4.Int64())
	assert.Equal(t, -92233720368547758.08, d4.RawFloat())
	assert.Equal(t, int64(-9223372036854775808), d5.Int64())
	assert.Equal(t, -92233720368547758.09, d5.RawFloat())

	assert.Equal(t, 0, d5.Compare(d4))
	assert.Equal(t, -1, d5.Compare(d3))

	d6, err := NewFromString("92233720368547758.06", 100.0)
	assert.Nil(t, err)
	d7, err := NewFromString("92233720368547758.07", 100.0)
	assert.Nil(t, err)
	d8, err := NewFromString("92233720368547758.08", 100.0)
	assert.Nil(t, err)
	assert.Equal(t, int64(9223372036854775806), d6.Int64())
	assert.Equal(t, 92233720368547758.06, d7.RawFloat())
	assert.Equal(t, int64(9223372036854775807), d7.Int64())
	assert.Equal(t, 92233720368547758.07, d7.RawFloat())
	assert.Equal(t, int64(9223372036854775807), d8.Int64())
	assert.Equal(t, 92233720368547758.08, d8.RawFloat())

	assert.Equal(t, 0, d8.Compare(d7))
	assert.Equal(t, 1, d8.Compare(d6))

}
