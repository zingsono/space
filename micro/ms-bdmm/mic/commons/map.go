package commons

// 合并Map集合
func MapMerge(args ...map[string]*graphql.Field) (m map[string]*graphql.Field) {
	for _, arg := range args {
		for k, v := range arg {
			m[k] = v
		}
	}
	return m
}
