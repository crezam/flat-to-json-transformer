package main

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestProcessFile(t *testing.T) {
	t.Log("Given the need to test the function transforms the flat file into the expected JSON")
	{
		f, err := os.Open("testdata.txt")
		check(err)
		defer f.Close()
		t.Log("Processing testdata.txt")
		actual := processFile(f)
		b, _ := json.Marshal(actual)
		var out bytes.Buffer
		json.Indent(&out, b, "", "\t")
		t.Log("Processed output:", out.String())
		expf, err := os.Open("testdataout.json")
		check(err)
		jsonParser := json.NewDecoder(expf)
		expected := TransactionsWrap{}
		jsonParser.Decode(&expected)
		t.Log("Expected transactions struct: ", expected)
		t.Log("Comparing expected to actual")
		if !reflect.DeepEqual(expected, actual) {
			t.FailNow()
		}
	}
}
