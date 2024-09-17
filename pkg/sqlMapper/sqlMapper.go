// SQL Mapper LITE Package.
//
// Aims to provide native lite solutions to the sql struct mapping problem
package sqlMapper

import (
	"reflect"
)

// GetPointerArray returns an array with pointers to values of input struct of Type T
//
// Sample Struct [github.com/pczavre/go-sql-mapper-lite/examples/sqlMapperLite]
// using reflect
func GetPointerArray[T any](t *T) []interface{} {
	var ptrArray []interface{}

	ref := reflect.ValueOf(t)
	fieldCount := reflect.Indirect(ref).NumField()

	for i := range fieldCount {
		curr := reflect.Indirect(ref).Field(i)
		currAddr := curr.Addr()
		ptrArray = append(ptrArray, currAddr.Interface())
	}

	return ptrArray
}
