package recursion

func DecimalToBin(i int) []byte {
	if i == 0 {
		return []byte{}
	}

	d := i / 2
	r := i % 2

	bin := []byte{byte(r)}
	bin = append(DecimalToBin(d), bin...)

	return bin
}
