package settingo

import (
	"os"
	"reflect"
	"testing"
)

type TestConfig struct {
	Foobar     string              `settingo:"help text"`
	FoobarInt  int                 `settingo:"help text"`
	FoobarBool bool                `settingo:"help text"`
	FoobarMap  map[string][]string `settingo:"help text"`
}

func Test_struct_types_default(t *testing.T) {

	expected := "default_value_for_foobar"
	expectedInt := 42
	expectedBool := true
	expectedMap := make(map[string][]string)

	expectedMap["foo"] = []string{"bar"}
	expectedMap["foo1"] = []string{"bar1", "bar2"}

	config := &TestConfig{
		Foobar:     expected,
		FoobarInt:  expectedInt,
		FoobarBool: expectedBool,
		FoobarMap:  expectedMap,
	}

	SETTINGS.ParseTo(config)

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

}

type ExampleConfig struct {
	Foobar     string              `settingo:"help text"`
	FoobarInt  int                 `settingo:"help text"`
	FoobarBool bool                `settingo:"help text"`
	FoobarMap  map[string][]string `settingo:"help text"`
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

	config := &ExampleConfig{
		Foobar:     defaultStr,
		FoobarInt:  defaultInt,
		FoobarBool: defaultBool,
		FoobarMap:  defaultMap,
	}

	SETTINGS.ParseTo(config)

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

	// Cleanup
	os.Unsetenv("FOOBAR")
	os.Unsetenv("FOOBAR_INT")
	os.Unsetenv("FOOBAR_BOOL")
	os.Unsetenv("FOOBAR_MAP")
}
