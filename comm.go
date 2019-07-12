/**
 * Copyright 2019 Dasn And Peihuwang
 *公共类库，常用的函数或者操作的方法
 */
package php2go

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

const (
	//定义需要格式化出的时间戳格式
	SHORT_TIMESTR_LAYLOT = "2006-01-02"
	LOGIN_TIMESTR_LAYOUT = "2006-01-02 15:04:05"
	SHORT_DAY_INT        = "20060102"
	HOUR_MINUTE          = "1504"
	//PhoneRegular = `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	PhoneRegular = "1{1}\\d{10}"
)

/**
 *检测是否是手机号
 *@param  string  手机号
 *@return {[type]}     bool
 */
func IsPhone(phone string) bool {
	rgx := regexp.MustCompile(PhoneRegular)
	return rgx.MatchString(phone)
}

/**
 *md5加密字符串
 */
func Mdwstr(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	md5 := fmt.Sprintf("%x", h.Sum(nil))
	return md5
	// byteData := []byte(str)
	// return fmt.Sprintf("%x", md5.Sum(byteData))
}

/**
 * 中文字符串长度
 */
func StrLen(str string) int {
	return utf8.RuneCountInString(str)
}

/**
 *字符串转byte
 *https://segmentfault.com/a/1190000005006351
 */
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/**
 *byte转字符串
 */
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

/**
 *unicode编码解析成中文
 */
func UnicodeToStr(str string) string {
	str2, _ := strconv.Unquote(`"` + str + `"`)
	return string(str2)
}

/**
 *json字符串bytes了下转map数据类型
 */
func JsonToMap(jstr string) map[string]interface{} {
	//fmt.Println("接收到的数据:", jstr, "-end")

	jsonMap := map[string]interface{}{}

	json.Unmarshal(Str2bytes(jstr), &jsonMap)
	//fmt.Println("Str2bytes:", Str2bytes(jstr), "-end")
	//fmt.Println("获取到的json=", jsonMap)
	return jsonMap
}

/**
 *map类型转字符串
 */
func MapToJson(mapData map[string]interface{}) string {

	data, _ := json.Marshal(mapData)

	return Bytes2str(data)
}

/**
 *json中的单个数组或者字符串转map
 *主要目的是再为止类型的情况下强制map[string]string 会panic一个语法错误终止程序
 */
func ArrOrMapToMap(val interface{}) map[string]string {
	uidList := map[string]string{}
	switch vv := val.(type) {
	case []interface{}: //是数组
		for key, kval := range vv {
			uidList[GetString(kval)] = GetString(key)
		}
	case map[string]interface{}: //是object
		for key, kval := range vv {
			uidList[key] = GetString(kval)
		}
	case int:
		uidList[GetString(val)] = "1"
	case string:
		uidList[GetString(val)] = "1"
	case float64:
		uidList[GetString(val)] = "1"
	default:
		fmt.Println("无法转换的数据类型")
	}
	return uidList
}

/**
 *强制转换map，arr到map[string]int格式
 */
func ArrOrMapToIntValMap(val interface{}) map[string]int {
	uidList := map[string]int{}
	switch vv := val.(type) {
	case []interface{}: //是数组
		for key, kval := range vv {
			uidList[GetString(kval)] = GetInt(key)
		}
	case map[string]interface{}: //是object
		for key, kval := range vv {
			uidList[key] = GetInt(kval)
		}
	case int:
		uidList[GetString(val)] = 1
	case string:
		uidList[GetString(val)] = 1
	case float64:
		uidList[GetString(val)] = 1
	default:
		fmt.Println("无法转换的数据类型")
		return uidList
	}
	fmt.Println("最终转换后的值", uidList)
	return uidList
}

/**
 * =============时间相关===================
 *返回时间戳
 *unix返回的是int64所以要转下类型
 */
func Now() int {
	return int(time.Now().Unix())
}

/**
 *8位日期，字符串的形式显示
 */
func DayTimeStr(timeNum int) string {
	if timeNum <= 0 {
		timeNum = Now()
	}
	return time.Unix(int64(timeNum), 0).Format(SHORT_DAY_INT)
}

/**
 * 获取当天8位整数
 */
func TodayInt() int {
	return GetInt(time.Unix(int64(Now()), 0).Format(SHORT_DAY_INT))
}

/**
 * 返回小时和分钟
 * @param {[type]} timeNum [description]
 */
func DayHourMinuteStr(timeNum int) string {
	if timeNum <= 0 {
		timeNum = Now()
	}
	return time.Unix(int64(timeNum), 0).Format(HOUR_MINUTE)
}

/**
 *8位日期整数
 */
func DayInt(timeNum int) int {
	if s, err := strconv.Atoi(DayTimeStr(0)); err == nil {
		return s
	}
	return 0
}

/**
 *格式化时间戳
 *Unix接收的是int64所以需要转换下
 */
func TimeStr(timeNum int) string {
	return time.Unix(int64(timeNum), 0).Format(LOGIN_TIMESTR_LAYOUT)
}

/**
 *整形数据转字符串
 */
func IntToString(val int) string {
	return strconv.Itoa(val)
}

/**
 *字符转小写
 */
func StrToLower(str string) string {
	return strings.ToLower(str)
}

// convert interface to string.
func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	default:
		if v != nil {
			return fmt.Sprintf("%v", result)
		}
	}
	return ""
}

// convert interface to int.
func GetInt(v interface{}) int {
	switch result := v.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	case float64:
		return int(result)
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.Atoi(d)
			return value
		}
	}
	return 0
}

// convert interface to int64.
func GetInt64(v interface{}) int64 {
	switch result := v.(type) {
	case int:
		return int64(result)
	case int32:
		return int64(result)
	case int64:
		return result
	default:

		if d := GetString(v); d != "" {
			value, _ := strconv.ParseInt(d, 10, 64)
			return value
		}
	}
	return 0
}

// convert interface to float64.
func GetFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case float64:
		return result
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.ParseFloat(d, 64)
			return value
		}
	}
	return 0
}

// convert interface to bool.
func GetBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.ParseBool(d)
			return value
		}
	}
	return false
}

// convert interface to byte slice.
func getByteArray(v interface{}) []byte {
	switch result := v.(type) {
	case []byte:
		return result
	case string:
		return []byte(result)
	default:
		return nil
	}
}
