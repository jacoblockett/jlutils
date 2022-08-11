package jlutils

func ContainsString(a []string, p string) (matched bool) {
	for _, x := range a {
		if x == p {
			return true
		}
	}

	return false
}
