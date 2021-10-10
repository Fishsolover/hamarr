package typechange

import "strconv"

// Int64ToInt int64转int
func Int64ToInt(i int64) int {
	return StringToInt(Int64ToString(i))
}

// Int64ToString int64转字符串
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// StringToInt64 字符串转int64
func StringToInt64(s string) int64 {
	return int64(StringToInt(s))
}

// StringToInt 字符串转int
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		//TODO 打印错误
		return 0
	}
	return i
}
