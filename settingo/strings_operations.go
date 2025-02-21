package settingo

import (
	"fmt"
	//	"strconv"
	"strings"
)

// add me
const ITEM_DELIMITER = ";"
const VAL_SEP = ","
const KEY_SEP = ":"

// should refactor to use error.
// Sanic gotta go fast
func parseKeyValue(s string) (string, []string, bool) {
	items := strings.Split(s, KEY_SEP)
	if len(items) != 2 {
		return "", nil, true
	}
	values := strings.Split(items[1], VAL_SEP)
	return items[0], values, false
}

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

func ParseMapToLine(m map[string][]string) string {
	items := []string{}
	for k, v := range m {
		items = append(items, fmt.Sprintf("%s%s%s", k, KEY_SEP, strings.Join(v, VAL_SEP)))
	}
	return strings.Join(items, ITEM_DELIMITER)
}

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
