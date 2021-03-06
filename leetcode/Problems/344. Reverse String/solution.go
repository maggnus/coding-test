func reverseString(s []byte) {
	i := 0
	size := len(s)
	for i <= size-i-1 {
		tmp := s[i]
		s[i] = s[size-i-1]
		s[size-i-1] = tmp
		i++
	}
}