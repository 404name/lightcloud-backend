package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func IntToPoint(i int) *int {
	return &i
}

func BoolToPoint(b bool) *bool {
	return &b
}

func UintToInt(u uint) *int {
	i := int(u)
	return &i
}

func StringToSlice(str string) []uint {
	str = strings.Trim(str, "[]")
	numbers := strings.Split(str, " ")
	var slice []uint
	for _, num := range numbers {
		i, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return slice
		}
		slice = append(slice, uint(i))
	}
	return slice
}
