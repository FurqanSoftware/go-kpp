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
	fmt.Printf("Metadata: %v\n", pkg.Metadata())
	stmts, err := pkg.Statements()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Statements: %v\n", stmts)
	tests, err := pkg.TestData()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Tests:\n")
	for _, e := range tests {
		fmt.Printf("  %v\n", e)
		fmt.Printf("  Name: %s\n", e.Name())
		fmt.Printf("  Group: %s\n", e.Group())
		fmt.Printf("  Hint: %s\n", mustString(t)(e.Hint()))
		fmt.Printf("  Description: %s\n", mustString(t)(e.Description()))
		fmt.Println()
	}
	subms, err := pkg.Submissions()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Submissions: %v\n", subms)
}

func mustString(t *testing.T) func(s string, err error) string {
	return func(s string, err error) string {
		if err != nil {
			t.Fatal(err)
			return ""
		}
		return s
	}
}
