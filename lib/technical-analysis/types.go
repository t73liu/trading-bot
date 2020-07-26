package analyze

import "encoding/json"

type ValidMicro struct {
	Micro int64
	Valid bool
}

func (vm ValidMicro) MarshalJson() ([]byte, error) {
	if vm.Valid {
		return json.Marshal(MicrosToDollars(vm.Micro))
	}
	return json.Marshal(nil)
}

type ValidFloat struct {
	Value float64
	Valid bool
}

func (vf ValidFloat) MarshalJson() ([]byte, error) {
	if vf.Valid {
		return json.Marshal(vf.Value)
	}
	return json.Marshal(nil)
}
