package settingo

import (
	"fmt"
	"strings"
)

// ITEM_DELIMITER is the delimiter used to separate key-value map items in a string representation.
//
// When parsing a string into a map using `ParseLineToMap`, this delimiter is used
// to split the string into individual key-value pairs.  For example, in the string
// "key1:value1,value2;key2:value3", the ITEM_DELIMITER is ";", separating "key1:value1,value2"
// and "key2:value3" as individual map items.
const ITEM_DELIMITER = ";"

// VAL_SEP is the separator used to delimit multiple values associated with a key in a string representation of a map.
//
// In `ParseLineToMap` and `ParseMapToLine`, this separator is used to split a string of values
// into a slice of strings. For example, in "key:value1,value2", VAL_SEP is ",", separating "value1" and "value2".
const VAL_SEP = ","

// KEY_SEP is the separator used to separate keys from their values in a string representation of a map.
//
// In `ParseLineToMap` and `ParseMapToLine`, this separator is used to split a key-value pair string
// into the key and the value string. For example, in "key:value1,value2", KEY_SEP is ":", separating "key" from "value1,value2".
const KEY_SEP = ":"

// parseKeyValue parses a single key-value string item into its key and values components.
//
// It expects the input string `s` to be in the format "key:value1,value2,...".
// The key and values are separated by KEY_SEP (":"). Multiple values are delimited by VAL_SEP (",").
//
// Args:
//
//	s: The string to parse, expected to be in "key:value1,value2,..." format.
//
// Returns:
//   - key:    The parsed key as a string. Returns an empty string if parsing fails.
//   - values: A slice of strings representing the parsed values. Returns nil if parsing fails.
//   - bool:   A boolean value indicating if an error occurred during parsing.
//     Returns true if the input string does not conform to the expected "key:value" format
//     (i.e., does not contain exactly one KEY_SEP), false otherwise.
//
// Example:
//
//	key, values, err := parseKeyValue("mykey:val1,val2")
//	// key will be "mykey"
//	// values will be []string{"val1", "val2"}
//	// err will be false
//
//	key, values, err = parseKeyValue("invalid-format")
//	// key will be ""
//	// values will be nil
//	// err will be true
func parseKeyValue(s string) (string, []string, bool) {
	items := strings.Split(s, KEY_SEP)
	if len(items) != 2 {
		return "", nil, true
	}
	values := strings.Split(items[1], VAL_SEP)
	return items[0], values, false
}

// ParseLineToMap parses a line string into a map[string][]string.
//
// It expects the input string `s` to be a series of key-value pairs separated by ITEM_DELIMITER (";").
// Each key-value pair is expected to be in the format "key:value1,value2,...", as parsed by `parseKeyValue`.
//
// If an item in the string cannot be parsed into a key-value pair (i.e., `parseKeyValue` returns an error),
// a message is printed to the console indicating the discarded item, and parsing continues with the next item.
//
// Args:
//
//	s: The line string to parse into a map.
//
// Returns:
//
//	A map[string][]string representing the parsed key-value pairs.
//	Keys are strings, and values are slices of strings.
//
// Example:
//
//	line := "key1:val1,val2;key2:val3;invalid-item"
//	parsedMap := ParseLineToMap(line)
//	// parsedMap will be:
//	// map[string][]string{
//	//    "key1": {"val1", "val2"},
//	//    "key2": {"val3"},
//	// }
//	// "Settingo: Unable to parse line, discarded: invalid-item" will be printed to console.
func ParseLineToMap(s string) map[string][]string {
	parsed := make(map[string][]string)
	items := strings.Split(s, ITEM_DELIMITER)
	for _, item := range items {
		key, values, err := parseKeyValue(item)
		if err {
			fmt.Println("Settingo: Unable to parse line, discarded:", item)
			continue
		}
		parsed[key] = values
	}
	return parsed
}

// ParseMapToLine converts a map[string][]string into a line string representation.
//
// It iterates through the input map `m` and formats each key-value pair into a string
// "key:value1,value2,...", where values are joined by VAL_SEP (",").
// These key-value strings are then joined together using ITEM_DELIMITER (";") to form the final line string.
//
// Args:
//
//	m: The map[string][]string to convert to a line string.
//
// Returns:
//
//	A string representation of the map, in the format "key1:value1,value2;key2:value3;...".
//	Returns an empty string if the input map is empty.
//
// Example:
//
//	inputMap := map[string][]string{
//		"key1": {"val1", "val2"},
//		"key2": {"val3"},
//	}
//	line := ParseMapToLine(inputMap)
//	// line will be "key1:val1,val2;key2:val3" (order of keys might vary)
func ParseMapToLine(m map[string][]string) string {
	items := []string{}
	for k, v := range m {
		items = append(items, fmt.Sprintf("%s%s%s", k, KEY_SEP, strings.Join(v, VAL_SEP)))
	}
	return strings.Join(items, ITEM_DELIMITER)
}

// FlattenMapStrSlice takes a map[string][]string and returns a flattened slice of unique string values.
//
// It iterates through all values in the input map `ss` (which are slices of strings),
// and adds each individual string value to a set (using a map[string]bool for efficient uniqueness checking).
// Finally, it converts the set of unique string values into a slice of strings.
//
// Args:
//
//	ss: The map[string][]string to flatten.
//
// Returns:
//
//	A slice of strings containing all unique values from all string slices in the input map.
//	The order of items in the returned slice is not guaranteed to be consistent.
//
// Example:
//
//	inputMap := map[string][]string{
//		"key1": {"val1", "val2", "val1"}, // "val1" is repeated
//		"key2": {"val3", "val4"},
//	}
//	flattenedSlice := FlattenMapStrSlice(inputMap)
//	// flattenedSlice might be []string{"val1", "val2", "val3", "val4"} (order not guaranteed)
func FlattenMapStrSlice(ss map[string][]string) []string {
	uniqItems := make(map[string]bool)
	for _, values := range ss {
		for _, val := range values {
			uniqItems[val] = true
		}
	}
	items := []string{}
	for item := range uniqItems {
		items = append(items, item)
	}
	return items
}
