package utils

func ToNullString(any interface{}) (ns NullString) {
	value, ok := any.(string)
	ns.Valid = ok
	ns.String = value
	return ns
}

func ToNullBool(any interface{}) (nb NullBool) {
	value, ok := any.(bool)
	nb.Valid = ok
	nb.Bool = value
	return nb
}

func ToNullFloat64(any interface{}) (nf NullFloat64) {
	value, ok := any.(float64)
	nf.Valid = ok
	nf.Float64 = value
	return nf
}

func ToNullInt64(any interface{}) (ni NullInt64) {
	value, ok := any.(int64)
	ni.Valid = ok
	ni.Int64 = value
	return ni
}
