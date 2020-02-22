package reflect

import (
	"fmt"
	"reflect"
)

// 反射编程测试
func CheckType(v interface{})  {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("浮点数")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("带符号整数")
	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		fmt.Println("无符号整数")
	default:
		fmt.Println("Unknown",t)
	}
}
