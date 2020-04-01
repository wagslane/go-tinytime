package tinytime

import "time"

// MarshalJSON implements the json.Marshaler interface
func (tt TinyTime) MarshalJSON() ([]byte, error) {
	return []byte(tt.Format(`"` + time.RFC3339 + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tt *TinyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	var err error
	*tt, err = Parse(`"`+time.RFC3339+`"`, string(data))
	return err
}
