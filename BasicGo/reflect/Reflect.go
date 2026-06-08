package reflect

import (
	"fmt"
	"reflect"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("signed integer")
	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		fmt.Println("unsigned integer")
	default:
		fmt.Println("Unknown", t)
	}
}
