package tools

func GetKeys[T comparable, T2 any](m map[T]T2) []T {
	keys := []T{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}