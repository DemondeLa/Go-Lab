package slice

func Intersect[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(b))
	for _, v := range b {
		m[v] = struct{}{}
	}

	result := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v)
		}
	}

	return result
}

func Union[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(a)+len(b))
	result := make([]T, 0, len(a)+len(b))

	add := func(v T) {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = struct{}{}
		}
	}
	for _, v := range a {
		add(v)
	}
	for _, v := range b {
		add(v)
	}

	return result
}

func Diff[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(b))
	for _, v := range b {
		m[v] = struct{}{}
	}
	result := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = struct{}{}
		}
	}
	return result
}
