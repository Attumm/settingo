package settingo

import (
	"net/url"
	"os"
	"reflect"
	"testing"
)

type ExampleConfig struct {
	Foobar      string              `settingo:"help text for foobar"`
	FoobarInt   int                 `settingo:"help text for FoobarInt"`
	FoobarBool  bool                `settingo:"help text for FoobarBool"`
	FoobarMap   map[string][]string `settingo:"help text FoobarMap"`
	FoobarSlice []string            `settingo:"help text for FoobarSlice"`
	FooParse    string              `settingo:"help text for FooParse"`
	FooParseInt int                 `settingo:"help text for FooParseInt"`
}

func Test_struct_types_os_env(t *testing.T) {
	expected := "other value"
	os.Setenv("FOOBAR", expected)
	defaultStr := "default value"

	expectedInt := 44
	os.Setenv("FOOBARINT", "44")
	defaultInt := 42

	os.Setenv("FOOBARBOOL", "y")
	expectedBool := true
	defaultBool := false

	os.Setenv("FOOBARMAP", "foo:bar;foo1:bar1,bar2")
	expectedMap := make(map[string][]string)
	defaultMap := make(map[string][]string)

	expectedMap["foo"] = []string{"bar"}
	expectedMap["foo1"] = []string{"bar1", "bar2"}

	os.Setenv("FOOBARSLICE", "item1,item2,item3")
	expectedSlice := []string{"item1", "item2", "item3"}
	defaultSlice := []string{}

	os.Setenv("FOOPARSE", "postgres://user:pass@database.example.com:5432/mydb")
	expectedFooParse := "database.example.com"
	defaultFooParse := "foobar"

	SetParsed("FOOPARSE", defaultFooParse, "database hostname", func(input string) string {
		u, err := url.Parse(input)
		if err != nil {
			return input
		}
		return u.Hostname()
	})

	Set("FOOBAR", defaultStr, "help text for foobar")
	SetString("FOOBAR", defaultStr, "help text for foobar")
	SetInt("FOOBARINT", defaultInt, "help text for FoobarInt")
	SetBool("FOOBARBOOL", defaultBool, "help text for FoobarBool")
	SetMap("FOOBARMAP", defaultMap, "help text FoobarMap")
	SetSlice("FOOBARSLICE", defaultSlice, "help text for FoobarSlice", ",")

	config := &ExampleConfig{
		Foobar:      defaultStr,
		FoobarInt:   defaultInt,
		FoobarBool:  defaultBool,
		FoobarMap:   defaultMap,
		FoobarSlice: defaultSlice,
		FooParse:    defaultFooParse,
	}

	ParseTo(config)

	if Get("FOOBAR") != expected {
		t.Error(Get("FOOBAR"), " != ", expected)
	}

	if GetInt("FOOBARINT") != expectedInt {
		t.Error(GetInt("FOOBARINT"), " != ", expectedInt)
	}

	if GetBool("FOOBARBOOL") != expectedBool {
		t.Error(GetBool("FOOBARBOOL"), " != ", expectedBool)
	}

	if !reflect.DeepEqual(GetMap("FOOBARMAP"), expectedMap) {
		t.Error(GetMap("FOOBARMAP"), " != ", expectedMap)
	}

	if !reflect.DeepEqual(GetSlice("FOOBARSLICE"), expectedSlice) {
		t.Error(GetSlice("FOOBARSLICE"), " != ", expectedSlice)
	}

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

	// Cleanup
	os.Unsetenv("FOOBAR")
	os.Unsetenv("FOOBARINT")
	os.Unsetenv("FOOBARBOOL")
	os.Unsetenv("FOOBARMAP")
	os.Unsetenv("FOOBARSLICE")
	os.Unsetenv("FOOPARSE")
	os.Unsetenv("FOOPARSEINT")
}
