package utils

func ToNullString(any interface{}) *NullString {
	value, ok := any.(string)
	var ns NullString
	ns.Valid = ok
	ns.String = value
	return &ns
}

func ToNullBool(any interface{}) *NullBool {
	value, ok := any.(bool)
	var nb NullBool
	nb.Valid = ok
	nb.Bool = value
	return &nb
}

func ToNullFloat64(any interface{}) *NullFloat64 {
	value, ok := any.(float64)
	var nf NullFloat64
	nf.Valid = ok
	nf.Float64 = value
	return &nf
}

func ToNullInt64(any interface{}) *NullInt64 {
	value, ok := any.(int64)
	var ni NullInt64
	ni.Valid = ok
	ni.Int64 = value
	return &ni
}
