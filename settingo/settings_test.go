package settingo

import (
	"net/url"
	"os"
	"reflect"
	"testing"
)

type TestConfig struct {
	Foobar      string              `settingo:"help text for foobar"`
	FoobarInt   int                 `settingo:"help text for FoobarInt"`
	FoobarBool  bool                `settingo:"help text for FoobarBool"`
	FoobarMap   map[string][]string `settingo:"help text FoobarMap"`
	FoobarSlice []string            `settingo:"help text for FoobarSlice"`
	FooParse    string              `settingo:"help text for FooParse"`
}

func Test_struct_types_default(t *testing.T) {

	expected := "default_value_for_foobar"
	expectedInt := 42
	expectedBool := true
	expectedMap := make(map[string][]string)

	expectedMap["foo"] = []string{"bar"}
	expectedMap["foo1"] = []string{"bar1", "bar2"}

	os.Setenv("FOOBARSLICE", "item1,item2,item3")
	expectedSlice := []string{"item1", "item2", "item3"}

	os.Setenv("FOOPARSE", "postgres://user:pass@database.example.com:5432/mydb")
	expectedFooParse := "database.example.com"

	// Parser that extracts hostname from database URL
	SETTINGS.SetParsed("FOOPARSE", "", "database hostname", func(input string) string {
		u, err := url.Parse(input)
		if err != nil {
			return input
		}
		return u.Hostname()
	})

	config := &TestConfig{
		Foobar:      expected,
		FoobarInt:   expectedInt,
		FoobarBool:  expectedBool,
		FoobarMap:   expectedMap,
		FoobarSlice: expectedSlice,
		FooParse:    expectedFooParse,
	}

	SETTINGS.LoadStruct(config)

	if config.Foobar != expected {
		t.Error(config.Foobar, " != ", expected)
	}

	if config.FoobarInt != expectedInt {
		t.Error(config.FoobarInt, " != ", expectedInt)
	}

	if config.FoobarBool != expectedBool {
		t.Error(config.FoobarBool, " != ", expectedBool)
	}

	if !reflect.DeepEqual(config.FoobarMap, expectedMap) {
		t.Error(config.FoobarMap, " != ", expectedMap)
	}

	if !reflect.DeepEqual(config.FoobarSlice, expectedSlice) {
		t.Error(config.FoobarSlice, " != ", expectedSlice)
	}

	if config.FooParse != expectedFooParse {
		t.Error(config.FooParse, " != ", expectedFooParse)
	}
}
