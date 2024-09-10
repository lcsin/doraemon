package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5 md5
func MD5(text string) string {
	md := md5.New()
	md.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(md.Sum(nil)))
}

// Paginate 逻辑分页，返回分页后的数组
func Paginate[T any](list []T, pageNo, pageSize int64) []T {
	if pageNo <= 0 || pageSize <= 0 {
		return []T{}
	}

	total := int64(len(list))
	totalPage := total / pageSize
	if totalPage%pageSize != 0 {
		totalPage++
	}

	start := (pageNo - 1) * pageSize
	if start >= total {
		return []T{}
	}

	end := start + pageSize
	if end > total {
		return list[start:]
	}
	return list[start:end]
}

// DelSliceElement 删除切片中某个下标的元素
func DelSliceElement[T any](list []T, idx int) []T {
	return append(list[:idx], list[idx+1:]...)
}
