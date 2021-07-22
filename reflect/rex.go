package reflect

import (
	"fmt"
	"reflect"
)


func testReflect() {
	rv := []interface{}{"hi", 42, func() {}}
	for _, r := range rv {
		switch v := reflect.ValueOf(r); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
			fmt.Println(reflect.TypeOf(r))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}
}
