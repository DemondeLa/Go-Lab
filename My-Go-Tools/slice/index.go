package slice

func Contains[T comparable](s []T, target T) bool {
	// 用底层原语组合出上层 API
	return Index(s, target) >= 0
}

func ContainsFunc[T any](s []T, match func(T) bool) bool {
	return IndexFunc(s, match) >= 0
}

func Index[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

func IndexFunc[T any](s []T, match func(T) bool) int {
	for i, v := range s {
		if match(v) {
			return i
		}
	}
	return -1
}
