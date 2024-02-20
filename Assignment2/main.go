package main

import (
	"errors"
	"fmt"
	"reflect"
)

func mergeSlices(a interface{}, b interface{}) (interface{}, error) {
	if a == nil && b == nil {
		return nil, errors.New("empty slices")
	}
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	var res []interface{}
	//for a

	if reflect.TypeOf(a).Kind() == reflect.Slice {
		slice := reflect.ValueOf(a)
		for i := 0; i < slice.Len(); i++ {
			if reflect.ValueOf(slice.Index(i)).Kind() == reflect.Slice {
				nestedSlice, err := mergeSlices(res, slice.Index(i).Interface())
				if err != nil {
					fmt.Println("error")
				}
				res = nestedSlice.([]interface{})
			} else {
				res = append(res, slice.Index(i).Interface())
			}
		}
	} else {
		res = append(res, a)
	}
	//for b

	if reflect.TypeOf(b).Kind() == reflect.Slice {
		slice := reflect.ValueOf(b)
		for i := 0; i < slice.Len(); i++ {
			if reflect.ValueOf(slice.Index(i).Interface()).Kind() == reflect.Slice {
				nestedSlice, err := mergeSlices(res, slice.Index(i).Interface())
				if err != nil {
					fmt.Println("error")
				}
				res = nestedSlice.([]interface{})
			} else {
				res = append(res, slice.Index(i).Interface())
			}
		}
	} else {
		res = append(res, b)
	}

	return res, nil

}

func main() {
	a := []interface{}{[]string{"BHAVYA"}, 9, 10, []int{1, 2}, []interface{}{"a", "xfy", 20.0}}
	b := []interface{}{1, 5, []int{3, 4}, []interface{}{true, "abh", 2}}

	merge, err := mergeSlices(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(merge)
}
