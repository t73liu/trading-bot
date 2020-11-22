package analyze

import (
	"encoding/json"
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

type ValidMicro struct {
	Valid bool
	Value int64
}

func (vm *ValidMicro) MarshalJSON() ([]byte, error) {
	if vm.Valid {
		return json.Marshal(utils.RoundToTwoDecimals(candle.MicrosToDollars(vm.Value)))
	}
	return json.Marshal(nil)
}

type ValidFloat struct {
	Valid bool
	Value float64
}

func (vf *ValidFloat) MarshalJSON() ([]byte, error) {
	if vf.Valid {
		return json.Marshal(vf.Value)
	}
	return json.Marshal(nil)
}

type ValidBool struct {
	Valid bool
	Value bool
}

func (vb *ValidBool) MarshalJSON() ([]byte, error) {
	if vb.Valid {
		return json.Marshal(vb.Value)
	}
	return json.Marshal(nil)
}

type ValidMicroRange struct {
	Valid bool
	High  int64
	Mid   int64
	Low   int64
}
