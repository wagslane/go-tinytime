package tinytime

// GobDecode implements the gob.GobDecoder interface
func (tt *TinyTime) GobDecode(data []byte) error {
	return tt.UnmarshalBinary(data)
}

// GobEncode implements the gob.GobEncoder interface
func (tt TinyTime) GobEncode() ([]byte, error) {
	return tt.MarshalBinary()
}
