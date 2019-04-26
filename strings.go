package helferlein

//StringInSlice checks if a string is in a slice of strings
func StringInSlice(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}
