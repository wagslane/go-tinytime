package tinytime

import "errors"

const tinyTimeBinaryVersion byte = 1

// MarshalBinary implements the encoding.BinaryMarshaler interface
func (tt TinyTime) MarshalBinary() ([]byte, error) {
	enc := []byte{
		tinyTimeBinaryVersion, // byte 0 : version
		byte(tt.unix >> 24),   // byte 1-4: unix
		byte(tt.unix >> 16),
		byte(tt.unix >> 8),
		byte(tt.unix),
	}
	return enc, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface
func (tt *TinyTime) UnmarshalBinary(data []byte) error {
	buf := data
	if len(buf) == 0 {
		return errors.New("tinytime: no data")
	}

	if buf[0] != tinyTimeBinaryVersion {
		return errors.New("tinytime: unsupported version")
	}

	if len(buf) != /*version*/ 1+ /*unix*/ 4 {
		return errors.New("tinytime: invalid length")
	}

	tt.unix = uint32(buf[4]) |
		uint32(buf[3])<<8 |
		uint32(buf[2])<<16 |
		uint32(buf[1])<<24

	return nil
}
