package settingo

import (
	"os"
	"testing"
)

func Test_types_default(t *testing.T) {
	expected := "default_value_for_foobar"
	expectedInt := 42
	expectedBool := true

	SETTINGS.Set("FOOBAR", expected, "help text")
	SETTINGS.SetInt("FOOBAR_INT", expectedInt, "help text")
	SETTINGS.SetBool("FOOBAR_BOOL", expectedBool, "help text")

	//SETTINGS.Parse()

	foobar := SETTINGS.Get("FOOBAR")
	if foobar != expected {
		t.Error(foobar, " != ", expected)
	}

	foobarInt := SETTINGS.GetInt("FOOBAR_INT")
	if foobarInt != expectedInt {
		t.Error(foobarInt, " != ", expectedInt)
	}

	foobarBool := SETTINGS.GetBool("FOOBAR_BOOL")
	if foobarBool != expectedBool {
		t.Error(foobarBool, " != ", expectedBool)
	}
}

func Test_types_os_env(t *testing.T) {

	expected := "other value"
	os.Setenv("FOOBAR", expected)
	defaultStr := "default value"

	expectedInt := 44
	os.Setenv("FOOBAR_INT", "44")
	defaultInt := 42

	os.Setenv("FOOBAR_BOOL", "y")
	expectedBool := true
	defaultBool := false

	SETTINGS.Set("FOOBAR", defaultStr, "help text")
	SETTINGS.SetInt("FOOBAR_INT", defaultInt, "help text")
	SETTINGS.SetBool("FOOBAR_BOOL", defaultBool, "help text")

	SETTINGS.Parse()

	foobar := SETTINGS.Get("FOOBAR")
	if foobar != expected {
		t.Error(foobar, " != ", expected)
	}

	foobarInt := SETTINGS.GetInt("FOOBAR_INT")
	if foobarInt != expectedInt {
		t.Error(foobarInt, " != ", expectedInt)
	}

	foobarBool := SETTINGS.GetBool("FOOBAR_BOOL")
	if foobarBool != expectedBool {
		t.Error(foobarBool, " != ", expectedBool)
	}
}
