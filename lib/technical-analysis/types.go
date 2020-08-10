package analyze

import (
	"encoding/json"
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

type ValidMicro struct {
	Micro int64
	Valid bool
}

func (vm ValidMicro) MarshalJSON() ([]byte, error) {
	if vm.Valid {
		return json.Marshal(utils.RoundToTwoDecimals(candle.MicrosToDollars(vm.Micro)))
	}
	return json.Marshal(nil)
}

type ValidFloat struct {
	Value float64
	Valid bool
}

func (vf ValidFloat) MarshalJSON() ([]byte, error) {
	if vf.Valid {
		return json.Marshal(vf.Value)
	}
	return json.Marshal(nil)
}
