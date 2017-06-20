package getmapkeys

// Get way
const (
	GetWayWithInit    = 1
	GetWayWithoutInit = 2
)

// GetMapKeys get keys from map
func GetMapKeys(i map[string]string, way int, o []string) {
	switch way {
	case GetWayWithoutInit:
		for a := range i {
			o = append(o, a)
		}
	case GetWayWithInit:
		o = make([]string, 0, len(i))
		for a := range i {
			o = append(o, a)
		}
	}
}
