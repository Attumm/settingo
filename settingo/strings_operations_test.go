package settingo

import (
	"reflect"
	"sort"
	"testing"
)

func TestParseEnvStringToMap(t *testing.T) {
	testcases := []struct {
		input          string
		expected_key   string
		expected_value []string
	}{
		{"foo:bla", "foo", []string{"bla"}},
		{"foo:bar,bla", "foo", []string{"bar", "bla"}},
		{"foo:bar,bla;foo2:bar2,bla2", "foo", []string{"bar", "bla"}},
		{"foo:bar,bla;foo2:bar2,bla2", "foo2", []string{"bar2", "bla2"}},
	}

	for tcNumber, testcase := range testcases {
		mapResult := ParseLineToMap(testcase.input)
		result, found := mapResult[testcase.expected_key]
		if !found {
			t.Error("testcase", tcNumber, "expected", "found", "!=", "not found")
		}
		if !reflect.DeepEqual(result, testcase.expected_value) {
			t.Error("testcase", tcNumber, "expected", testcase.expected_value, "!=", result)
		}
	}
}

func TestParseEnvMapToString(t *testing.T) {
	testcases := []struct {
		input          string
		expected_key   string
		expected_value []string
	}{
		{"foo:bla", "foo", []string{"bla"}},
		{"foo:bar,bla", "foo", []string{"bar", "bla"}},
		{"foo:bar,bla;foo2:bar2,bla2", "foo", []string{"bar", "bla"}},
		{"foo:bar,bla;foo2:bar2,bla2", "foo2", []string{"bar2", "bla2"}},
	}

	for tcNumber, testcase := range testcases {
		mapResult := ParseLineToMap(testcase.input)
		lineResult := ParseMapToLine(mapResult)
		newMapResult := ParseLineToMap(lineResult)
		result, found := newMapResult[testcase.expected_key]

		if !found {
			t.Error("testcase", tcNumber, "expected", "found", "!=", "not found")
		}
		if !reflect.DeepEqual(result, testcase.expected_value) {
			t.Error("testcase", tcNumber, "expected", testcase.expected_value, "!=", result)
		}
	}
}

func TestParseEnvStringToFlatten(t *testing.T) {
	testcases := []struct {
		input    string
		expected []string
	}{
		{"foo:bla", []string{"bla"}},
		{"foo:bar,bla", []string{"bar", "bla"}},
		{"foo:bar,bla;foo2:bar2,bla2", []string{"bar", "bla", "bar2", "bla2"}},
		{"foo:bar,bla;foo2:bar2,bla2", []string{"bar", "bla", "bar2", "bla2"}},
	}

	for tcNumber, testcase := range testcases {
		mapResult := ParseLineToMap(testcase.input)
		result := FlattenMapStrSlice(mapResult)
		sort.Strings(result)
		sort.Strings(testcase.expected)
		if !reflect.DeepEqual(result, testcase.expected) {
			t.Error("testcase", tcNumber, "expected", testcase.expected, "!=", result)
		}
	}
}

func TestParseKeyValueErrors(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		wantKey  string
		wantVals []string
		wantErr  bool
	}{
		{
			name:     "missing separator",
			input:    "foobla",
			wantKey:  "",
			wantVals: nil,
			wantErr:  true,
		},
		{
			name:     "empty string",
			input:    "",
			wantKey:  "",
			wantVals: nil,
			wantErr:  true,
		},
		{
			name:     "multiple key separators",
			input:    "foo:bar:baz",
			wantKey:  "",
			wantVals: nil,
			wantErr:  true,
		},
		{
			name:     "only separator",
			input:    ":",
			wantKey:  "",
			wantVals: []string{""},
			wantErr:  false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			gotKey, gotVals, gotErr := parseKeyValue(tc.input)
			if gotErr != tc.wantErr {
				t.Errorf("parseKeyValue(%q) error = %v, want %v", tc.input, gotErr, tc.wantErr)
			}
			if gotKey != tc.wantKey {
				t.Errorf("parseKeyValue(%q) key = %q, want %q", tc.input, gotKey, tc.wantKey)
			}
			if !reflect.DeepEqual(gotVals, tc.wantVals) {
				t.Errorf("parseKeyValue(%q) vals = %v, want %v", tc.input, gotVals, tc.wantVals)
			}
		})
	}
}

func TestParseLineToMapErrors(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  map[string][]string
	}{
		{
			name:  "single invalid item",
			input: "invalid",
			want:  map[string][]string{},
		},
		{
			name:  "mix of valid and invalid items",
			input: "foo:bar;invalid;baz:qux",
			want: map[string][]string{
				"foo": {"bar"},
				"baz": {"qux"},
			},
		},
		{
			name:  "multiple invalid items",
			input: "invalid1;invalid2;invalid3",
			want:  map[string][]string{},
		},
		{
			name:  "empty input",
			input: "",
			want:  map[string][]string{},
		},
		{
			name:  "only delimiters",
			input: ";;;",
			want:  map[string][]string{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := ParseLineToMap(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ParseLineToMap(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
