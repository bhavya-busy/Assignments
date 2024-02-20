package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

type Address struct {
	City  string
	State string
}

func PopulateStruct(m map[string]interface{}, p interface{}) {
	resValue := reflect.ValueOf(p).Elem()
	//elem because pointer sent, so to work with actual struct and not pointer

	for key, value := range m {
		// gives field name that matches the key
		field := resValue.FieldByName(key)
		if field.IsValid() {
			//if field struct, to uske andr traverse krna pdhega
			if field.Kind() == reflect.Struct {
				if nestedMap, ok := value.(map[string]interface{}); ok {
					nestedStruct := reflect.New(field.Type()).Interface()
					PopulateStruct(nestedMap, nestedStruct)
					field.Set(reflect.ValueOf(nestedStruct).Elem())
				}
			} else {

				field.Set(reflect.ValueOf(value))
			}
		}
	}

}

func main() {
	fmt.Println("populate")
	var m = map[string]interface{}{
		"Name":    "Bhavya",
		"Age":     21,
		"Email":   "abc@gmail.com",
		"Pincode": "110034",
		"Address": map[string]interface{}{
			"City":  "Delhi",
			"State": "Delhi",
		},
	}
	var p Person
	//p := Person{}
	PopulateStruct(m, &p) //made pointer
	fmt.Println(p)
}
