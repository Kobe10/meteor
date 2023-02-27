package utils

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var tenTo32 map[int]string = map[int]string{
	0:  "F",
	1:  "G",
	2:  "t",
	3:  "Q",
	4:  "a",
	5:  "n",
	6:  "z",
	7:  "h",
	8:  "c",
	9:  "U",
	10: "C",
	11: "e",
	12: "j",
	13: "8",
	14: "w",
	15: "V",
	16: "5",
	17: "S",
	18: "L",
	19: "y",
	20: "x",
	21: "Z",
	22: "s",
	23: "7",
	24: "K",
	25: "J",
	26: "6",
	27: "M",
	28: "9",
	29: "W",
	30: "r",
	31: "q",
	32: "H",
	33: "O",
	34: "p",
	35: "D",
	36: "N",
	37: "v",
	38: "Y",
	39: "3",
	40: "f",
	41: "u",
	42: "g",
	43: "l",
	44: "P",
	45: "A",
	46: "I",
	47: "d",
	48: "1",
	49: "0",
	50: "4",
	51: "B",
	52: "b",
	53: "E",
	54: "i",
	55: "m",
	56: "X",
	57: "T",
	58: "k",
	59: "R",
	60: "2",
	61: "o",
}

// 拼接短url
func SpliceShortUrl(num int) string {
	url := DecimalTo62(num, 62)
	return getUrlDomain() + url
}

// 10进制转62进制
func DecimalTo62(num, n int) string {
	newNumStr := ""
	var remainder int
	var remainderString string
	for num != 0 {
		remainder = num % n
		remainderString = tenTo32[remainder]
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}

// map根据value找key
func findKey(in string) int {
	result := -1
	for k, v := range tenTo32 {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func anyToDecimal(num string, n int) int {
	var newNum float64
	newNum = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(newNum)
}

// 初始化62个字符
var tenTo62InitStr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// 洗牌算法 -- 生成随机顺序的62个字符  -- 只是使用一次用来赋值给 【tenToAny】
func init62Str() map[int]string {
	tenToAny1 := map[int]string{}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tenTo62InitStr), func(i, j int) { tenTo62InitStr[i], tenTo62InitStr[j] = tenTo62InitStr[j], tenTo62InitStr[i] })
	println(tenTo62InitStr)
	for i := 0; i < len(tenTo62InitStr); i++ {
		tenToAny1[i] = tenTo62InitStr[i]
		print(i)
		print(":" + "\"")
		print(tenToAny1[i])
		println("\"" + ",")
	}
	return nil
}
