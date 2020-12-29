package utils

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) Value() string {
	return ns.String
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	ns.Valid = value != nil
	if value != nil {
		ns.String = *value
	}
	return nil
}

type NullBool struct {
	sql.NullBool
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	}
	return json.Marshal(nil)
}

func (nb *NullBool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	nb.Valid = value != nil
	if value != nil {
		nb.Bool = *value
	}
	return nil
}

func (nb *NullBool) Value() bool {
	return nb.Bool
}

func NewNullBool(value bool) (nb NullBool) {
	nb.Valid = true
	nb.Bool = value
	return nb
}

func EqNullBoolSlice(expected, actual []NullBool) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}
	return true
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	}
	return json.Marshal(nil)
}

func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	nf.Valid = value != nil
	if value != nil {
		nf.Float64 = *value
	}
	return nil
}

func (nf *NullFloat64) Value() float64 {
	return nf.Float64
}

func NewNullFloat64(value float64) (nf NullFloat64) {
	nf.Valid = true
	nf.Float64 = value
	return nf
}

func EqNullFloat64Slice(expected, actual []NullFloat64) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}
	return true
}

type NullInt64 struct {
	sql.NullInt64
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return json.Marshal(nil)
}

func (ni *NullInt64) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	ni.Valid = value != nil
	if value != nil {
		ni.Int64 = *value
	}
	return nil
}

func (ni *NullInt64) Value() int64 {
	return ni.Int64
}

func NewNullInt64(value int64) (ni NullInt64) {
	ni.Valid = true
	ni.Int64 = value
	return ni
}

func EqNullInt64Slice(expected, actual []NullInt64) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}
	return true
}

// One millionth of a dollar
type MicroDollar struct {
	sql.NullInt64
}

func (md *MicroDollar) Dollar() float64 {
	return RoundToTwoDecimals(MicrosToDollars(md.Int64))
}

func (md *MicroDollar) MarshalJSON() ([]byte, error) {
	if md.Valid {
		// Rounding introduces display issues for low values (e.g. MACD)
		return json.Marshal(MicrosToDollars(md.Int64))
	}
	return json.Marshal(nil)
}

func (md *MicroDollar) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	md.Valid = value != nil
	if value != nil {
		md.Int64 = DollarsToMicros(*value)
	}
	return nil
}

func (md *MicroDollar) Value() int64 {
	return md.Int64
}

func NewMicroDollar(value int64) (ni MicroDollar) {
	ni.Valid = true
	ni.Int64 = value
	return ni
}

func EqMicroDollarSlice(expected, actual []MicroDollar) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}
	return true
}
