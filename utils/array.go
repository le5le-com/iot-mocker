package utils

import "strings"

func ContainString(arr []string, v string) bool {
	for _, item := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ContainInt(arr []int, v int) bool {
	for _, item := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func UniqueJoin(words []string) string {
	uniqueMap := make(map[string]bool) // 创建map以存储唯一的单词
	var uniqueWords []string           // 创建一个切片，用于存储去重后的单词

	for _, word := range words {
		if _, exists := uniqueMap[word]; !exists {
			uniqueMap[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}

	return strings.Join(uniqueWords, " ")
}
