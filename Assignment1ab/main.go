package main

import (
	"fmt"
	"reflect"
)

func SetKeyValue(key string, source map[string]interface{}, value interface{}) {
	_, exists := source[key]
	if exists == true {
		source[key] = value
	}

	for _, v := range source {

		// fmt.Println("/////////////////")
		// fmt.Println(reflect.TypeOf(v))
		// fmt.Println(reflect.TypeOf(v).Kind())

		//kind will tell type's kind
		// switch reflect.TypeOf(v).Kind() {

		// case reflect.Slice:
		// 	//converted slice to slice of interface
		// 	for _, i := range v.([]interface{}) {
		// 		//if slice to inside it all values consider, so loop inside the slice
		// 		SetKeyValue(key, i.(map[string]interface{}), value)
		// 	}
		// case reflect.Map:
		// 	//if found in map, update
		// 	SetKeyValue(key, v.(map[string]interface{}), value)

		// }

		check := reflect.ValueOf(v).Kind()
		if check == reflect.Slice {
			for _, i := range v.([]interface{}) {
				//if slice to inside it all values consider, so loop inside the slice
				SetKeyValue(key, i.(map[string]interface{}), value)
			}
		}
		//map h to again recursion call krdo
		if check == reflect.Map {
			SetKeyValue(key, v.(map[string]interface{}), value)
		}

	}
}

func RemoveKey(key string, source map[string]interface{}) {
	_, exists := source[key]
	if exists == true {
		delete(source, key)
	}

	for _, v := range source {

		check := reflect.ValueOf(v).Kind()
		if check == reflect.Slice {
			for _, i := range v.([]interface{}) {
				//if slice to inside it all values consider, so loop inside the slice
				RemoveKey(key, i.(map[string]interface{}))
			}
		}
		//map h to again recursion call krdo
		if check == reflect.Map {
			RemoveKey(key, v.(map[string]interface{}))
		}

	}

}

func main() {
	var m = map[string]interface{}{
		"Name": "Bhavya",
		"DOB":  15 - 07 - 2002,
		"city": "Delhi",
		"pin":  110034,
		// field named "Address" and assigns it a slice ([]interface{}). The square brackets [] indicate that it's a slice, and interface{} allows elements of any data type to be stored in the slice.
		"NewTest": map[string]interface{}{
			"street":  "Testtt",
			"plot_no": 96,
			"city":    "Example city",
			"pin":     633078,
		},
		"Address": []interface{}{
			//nside the slice, there are two elements, each represented by a map.

			map[string]interface{}{
				"street":  "abc",
				"plot_no": 101,
				"city":    "Dwarka",
				"pin":     110078,
			},
			map[string]interface{}{
				"street":  "Chowk",
				"plot_no": 26,
				"city":    "London",
				"pin":     923478,
			},
		},
		"Salary":      800000,
		"Designation": "Developer",
	}
	var val interface{}
	val = "New York"

	SetKeyValue("city", m, val)
	fmt.Println(m)

	RemoveKey("city", m)
	fmt.Println(m)

}
