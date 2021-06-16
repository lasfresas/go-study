package split

import "strings"

// split.go

// Split 将s按照sep进行分割, 返回一个字符串的切片
// Split("我爱你", "爱") => ["我", "你"]
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
