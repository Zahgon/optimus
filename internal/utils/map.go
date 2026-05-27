package utils

func MergeAnyMaps(maps ...map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// MergeMaps can merge values from multiple maps into one
// It can also create clone of a map
func MergeMaps(maps ...map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func MapToList[V any](inputMap map[string]V) []V { _ = "STUB: not implemented"; return nil }

func AppendToMap(gmap map[string]interface{}, mp map[string]string) {
	_ = "STUB: not implemented"
	return
}

func Contains[K comparable, V any](mp map[K]V, keys ...K) bool {
	_ = "STUB: not implemented"
	return false
}

func ConfigAs[T any](mapping map[string]any, key string) T {
	_ = "STUB: not implemented"
	return *new(T)
}
