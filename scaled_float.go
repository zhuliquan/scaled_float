package scaled_float

import (
	"math"

	"github.com/shopspring/decimal"
)

// store_value = value * scaling_factor
type ScaledFloat struct {
	storeValue    int64
	value         decimal.Decimal
	scalingFactor decimal.Decimal
}

var (
	NegativeInf = &ScaledFloat{storeValue: math.MinInt64}
	PositiveInf = &ScaledFloat{storeValue: math.MaxInt64}
)

func new(v, s decimal.Decimal) *ScaledFloat {
	var storeValue = v.Mul(s)
	var storeSign = v.Sign() * s.Sign()
	var storeInt64 = storeValue.BigInt().Int64()
	if storeValue.Sign() == storeSign {
		if storeSign > 0 && storeInt64 < 0 {
			storeInt64 = PositiveInf.storeValue
		} else if storeSign < 0 && storeInt64 > 0 {
			storeInt64 = NegativeInf.storeValue
		}
	} else if storeSign > 0 {
		storeInt64 = PositiveInf.storeValue
	} else {
		storeInt64 = NegativeInf.storeValue
	}

	return &ScaledFloat{
		storeValue:    storeInt64,
		value:         v,
		scalingFactor: s,
	}
}

func New(value, scalingFactor float64) *ScaledFloat {
	var fValue = decimal.NewFromFloat(value)
	var sValue = decimal.NewFromFloat(scalingFactor)
	return new(fValue, sValue)

}

func NewFromString(value string, scalingFactor float64) (*ScaledFloat, error) {
	var fValue, err = decimal.NewFromString(value)
	if err != nil {
		return nil, err
	}
	var sValue = decimal.NewFromFloat(scalingFactor)
	return new(fValue, sValue), nil
}

func (s *ScaledFloat) RawFloat() float64 {
	var value, _ = s.value.Float64()
	return value
}

func (s *ScaledFloat) Int64() int64 {
	return s.storeValue
}

func (s *ScaledFloat) IsInf(sign int) bool {
	if sign == 0 {
		return s.storeValue == NegativeInf.storeValue || s.storeValue == PositiveInf.storeValue
	} else if sign < 0 {
		return s.storeValue == NegativeInf.storeValue
	} else {
		return s.storeValue == PositiveInf.storeValue
	}
}

func (a *ScaledFloat) Compare(b *ScaledFloat) int {
	if a.storeValue < b.storeValue {
		return -1
	} else if a.storeValue > b.storeValue {
		return 1
	} else {
		return 0
	}
}
