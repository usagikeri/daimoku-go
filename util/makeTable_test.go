package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeTable(t *testing.T) {

	table := MakeTable("../testdata/test.yaml")

	f, err := os.Open("../testdata/ans.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	ans := string(b)

	if table != ans {
		t.Fatal("Test Faild")
	}
}
