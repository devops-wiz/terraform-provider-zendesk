package utils

func SliceFilter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func GroupSlice[T any, K comparable](slice []T, keyFn func(T) K) map[K][]T {
	grouped := make(map[K][]T)
	for _, item := range slice {
		key := keyFn(item)
		grouped[key] = append(grouped[key], item)
	}
	return grouped
}
