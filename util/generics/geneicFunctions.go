package generics

import (
	// "fmt"
	"reflect"
)

func isComparable(v interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			// recover from panic if v is not comparable
		}
	}()
	
	// Check if the type is comparable by attempting to compare
	// We do a dummy comparison with itself
	return reflect.TypeOf(v).Comparable()
}
