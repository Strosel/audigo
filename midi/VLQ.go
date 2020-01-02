package midi

//VLQ is a midi variable length quantity
type VLQ uint32

//Bytes returns the VLQ as a byte array
func (v VLQ) Bytes() []byte {
	out := []byte{}

	//the first (read rightmost) byte should not have the bit 7 flag
	first := true
	for v > 0 {
		seven := byte(v & 127) //the first 7 bits
		if first {
			first = false
		} else {
			seven += 128
		}
		out = append(out, seven)
		v >>= 7 //remove the first 7 bits
	}

	//reverse the array for correct endianness
	for i, l := 0, len(out); i < l/2; i++ {
		out[i], out[l-i-1] = out[l-i-1], out[i]
	}

	return out
}
