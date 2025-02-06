package utils

import "sort"

// SortKey 根据key排序
func SortKey(data map[string]string) (keys []string) {
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
