package kpp

import (
	"fmt"
	"os"
	"testing"
)

func TestPackage(t *testing.T) {
	pkg, err := Open(os.DirFS("testdata/hello"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", pkg.Metadata())
	stmts, err := pkg.Statements()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", stmts)
	tests, err := pkg.TestData()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", tests)
}
