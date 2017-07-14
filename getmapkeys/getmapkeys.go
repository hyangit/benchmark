package getmapkeys

// WithInit is
func WithInit(i map[string]string) []string {
	o := make([]string, 0, len(i))
	for a := range i {
		o = append(o, a)
	}
	return o
}

// WithoutInit is
func WithoutInit(i map[string]string) []string {
	var o []string
	for a := range i {
		o = append(o, a)
	}
	return o
}
