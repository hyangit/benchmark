package distinct

// ByMap is
func ByMap(i []string) []string {
	om := make(map[string]bool)
	for _, t := range i {
		om[t] = true
	}
	o := make([]string, 0, len(om))
	for t := range om {
		o = append(o, t)
	}
	return o
}

// BySlice is
func BySlice(i []string) []string {
	o := make([]string, 0, len(i))
	for _, t := range i {
		exist := false
		for _, m := range o {
			if m == t {
				exist = true
				break
			}
		}
		if !exist {
			o = append(o, t)
		}
	}
	return o
}
