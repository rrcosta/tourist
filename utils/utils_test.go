package utils

import (
	"fmt"
	"testing"
)

const fileName = "files/input-test.csv"

func TestReadCsv(t *testing.T) {
	got, err := ReadCsv(fileName)
	// want := [[GRU BRC 10]]

	fmt.Println(got)
	fmt.Println(err)

	// if got != want {
	// 	t.Errorf("")
	// }
}
