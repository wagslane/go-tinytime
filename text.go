package tinytime

import "time"

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (tt TinyTime) MarshalText() ([]byte, error) {
	return []byte(tt.Format(time.RFC3339)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (tt *TinyTime) UnmarshalText(data []byte) error {
	newTD, err := Parse(time.RFC3339, string(data))
	if err != nil {
		return err
	}
	tt.unix = newTD.unix
	return nil
}
