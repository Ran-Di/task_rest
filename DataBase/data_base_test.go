package DataBase

import (
	"testing"
)

func Test_DataBase(t *testing.T) {
	db, err := Init("")
	if err != nil {
		t.Errorf("Can't connected with Data Base")
	}
	str, strT := "TEST", "decrypt"

	if err = AddRec(db, strT, str, str); err != nil {
		t.Errorf("AddRec finished with %s", err)
	}
}
