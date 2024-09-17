package sqlMapper

import (
	"database/sql"
	"fmt"
	"testing"
)

// Sample User Row
type UserRow struct {
	Id    sql.NullInt16
	Name  sql.NullString
	Email sql.NullString
	DOB   sql.NullString
}

func TestGeneric(t *testing.T) {

	user := &UserRow{}
	actualPtrArray := []interface{}{&user.Id, &user.Name, &user.Email, &user.DOB}
	genericPtrArray := GetPointerArray(user)

	if len(actualPtrArray) != len(genericPtrArray) {
		t.Fatal("Pointer Array length mismatch ", "actual Count : ", len(actualPtrArray), " generic Count : ", len(genericPtrArray))
	}

	for i, actualPtr := range actualPtrArray {

		genericPtr := genericPtrArray[i]

		if genericPtr != actualPtr {
			t.Fatal("Pointers do not match ", "Actual Pointer: ", actualPtr, " Generic Pointer: ", genericPtr)
		}

	}
	fmt.Println(actualPtrArray)
	fmt.Println(genericPtrArray)
}
