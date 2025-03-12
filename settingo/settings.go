package settingo

import (
	"flag"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func truthiness(s string) bool {
	thruthy := map[string]bool{
		"y":    true,
		"true": true,
		"yes":  true,
	}
	return thruthy[s]
}

type Settings struct {
	msg              map[string]string
	VarString        map[string]string
	VarInt           map[string]int
	VarBool          map[string]bool
	VarMap           map[string]map[string][]string
	VarSlice         map[string][]string
	VarSliceSep      map[string]string
	Parsers          map[string]func(string) string
	ParsersInt       map[string]func(int) int
	ContextualCasing bool
}

func (s *Settings) Set(flagName, defaultVar, message string) {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarString[flagName] = defaultVar
}

func (s *Settings) SetString(flagName, defaultVar, message string) {
	s.Set(flagName, defaultVar, message)
}

func (s *Settings) SetInt(flagName string, defaultVar int, message string) {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarInt[flagName] = defaultVar
}

func (s *Settings) SetBool(flagName string, defaultVar bool, message string) {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarBool[flagName] = defaultVar
}

func (s *Settings) SetMap(flagName string, defaultVar map[string][]string, message string) {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarMap[flagName] = defaultVar
}

func (s *Settings) SetSlice(flagName string, defaultVar []string, message string, sep string) {
	if sep == "" {
		sep = ","
	}
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarSlice[flagName] = defaultVar
	s.VarSliceSep[flagName] = sep
}

func (s *Settings) SetParsed(flagName, defaultVar, message string, parserFunc func(string) string) {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarString[flagName] = defaultVar
	s.Parsers[flagName] = parserFunc
}

func (s *Settings) SetParsedInt(flagName, defaultVar, message string, parserFunc func(int) int) {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	s.msg[flagName] = message
	s.VarString[flagName] = defaultVar
	s.ParsersInt[flagName] = parserFunc
}

func (s Settings) Get(flagName string) string {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	return s.VarString[flagName]
}

func (s Settings) GetInt(flagName string) int {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	return s.VarInt[flagName]
}

func (s Settings) GetBool(flagName string) bool {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	return s.VarBool[flagName]
}

func (s Settings) GetMap(flagName string) map[string][]string {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	return s.VarMap[flagName]
}

func (s Settings) GetSlice(flagName string) []string {
	if s.ContextualCasing {
		flagName = strings.ToLower(flagName)
	}
	return s.VarSlice[flagName]
}

func (s *Settings) HandleCMDLineInput() {
	parsedString := make(map[string]*string)
	for key, val := range s.VarString {
		var newV = flag.String(key, val, s.msg[key])
		parsedString[key] = newV
	}
	parsedInt := make(map[string]*int)
	for key, val := range s.VarInt {
		var newV = flag.Int(key, val, s.msg[key])
		parsedInt[key] = newV
	}
	parsedBool := make(map[string]*string)
	for key, val := range s.VarBool {
		var newV = flag.String(key, strconv.FormatBool(val), s.msg[key])
		parsedBool[key] = newV
	}
	parsedMap := make(map[string]*string)
	for key, val := range s.VarMap {
		var newV = flag.String(key, ParseMapToLine(val), s.msg[key])
		parsedBool[key] = newV
	}
	parsedSlice := make(map[string]*string)
	for key, val := range s.VarSlice {
		var newV = flag.String(key, strings.Join(val, s.VarSliceSep[key]), s.msg[key])
		parsedSlice[key] = newV
	}
	flag.Parse()

	for key, val := range parsedString {
		if parseFunc, found := s.Parsers[key]; found {
			s.VarString[key] = parseFunc(*val)
		} else {
			s.VarString[key] = *val
		}
	}
	for key, val := range parsedInt {
		if parseFunc, found := s.ParsersInt[key]; found {
			s.VarInt[key] = parseFunc(*val)
		} else {
			s.VarInt[key] = *val
		}
	}
	for key, val := range parsedBool {
		s.VarBool[key] = truthiness(*val)
	}
	for key, val := range parsedMap {
		s.VarMap[key] = ParseLineToMap(*val)
	}

	for key, val := range parsedSlice {
		s.VarSlice[key] = strings.Split(*val, s.VarSliceSep[key])
	}
}

func (s *Settings) HandleOSInput() {
	for key := range s.VarString {
		lookupKey := key
		if s.ContextualCasing {
			lookupKey = strings.ToUpper(key)
		}
		varEnv, found := os.LookupEnv(lookupKey)
		if found {
			s.VarString[key] = varEnv
		}
	}
	for key := range s.VarInt {
		lookupKey := key
		if s.ContextualCasing {
			lookupKey = strings.ToUpper(key)
		}
		varEnv, found := os.LookupEnv(lookupKey)
		if found {
			if num, err := strconv.Atoi(varEnv); err == nil {
				s.VarInt[key] = num
			}
		}
	}
	for key := range s.VarBool {
		lookupKey := key
		if s.ContextualCasing {
			lookupKey = strings.ToUpper(key)
		}
		varEnv, found := os.LookupEnv(lookupKey)
		if found {
			s.VarBool[key] = truthiness(varEnv)
		}
	}
	for key := range s.VarMap {
		lookupKey := key
		if s.ContextualCasing {
			lookupKey = strings.ToUpper(key)
		}
		varEnv, found := os.LookupEnv(lookupKey)
		if found {
			s.VarMap[key] = ParseLineToMap(varEnv)
		}
	}
	for key := range s.VarSlice {
		lookupKey := key
		if s.ContextualCasing {
			lookupKey = strings.ToUpper(key)
		}
		varEnv, found := os.LookupEnv(lookupKey)
		if found {
			s.VarSlice[key] = strings.Split(varEnv, s.VarSliceSep[key])
		}
	}
}

func (s *Settings) Parse() {
	s.HandleOSInput()
	s.HandleCMDLineInput()
}

func (s *Settings) ParseTo(to interface{}) {
	s.LoadStruct(to)
	s.Parse()
	s.UpdateStruct(to)
}

// LoadStruct registers a struct's fields with SETTINGS
func (s *Settings) LoadStruct(cfg interface{}) {
	val := reflect.ValueOf(cfg)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		name := strings.ToUpper(field.Name)
		help := field.Tag.Get("settingo")

		switch value.Kind() {
		case reflect.String:
			s.SetString(name, value.String(), help)
		case reflect.Int:
			s.SetInt(name, int(value.Int()), help)
		case reflect.Bool:
			s.SetBool(name, value.Bool(), help)
		case reflect.Slice:
			if value.Type().Elem().Kind() == reflect.String {
				slice := make([]string, value.Len())
				for i := 0; i < value.Len(); i++ {
					slice[i] = value.Index(i).String()
				}
				s.SetSlice(name, slice, help, s.VarSliceSep[strings.ToLower(name)])
			}
		case reflect.Map:
			if value.Type().Key().Kind() == reflect.String &&
				value.Type().Elem().Kind() == reflect.Slice &&
				value.Type().Elem().Elem().Kind() == reflect.String {
				m := value.Interface().(map[string][]string)
				s.SetMap(name, m, help)
			}
		}
	}
}

// UpdateStruct updates a struct with values from SETTINGS after Parse()
func (s *Settings) UpdateStruct(cfg interface{}) {
	val := reflect.ValueOf(cfg)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		name := strings.ToUpper(field.Name)

		switch value.Kind() {
		case reflect.String:
			value.SetString(s.Get(name))
		case reflect.Int:
			value.SetInt(int64(s.GetInt(name)))
		case reflect.Bool:
			value.SetBool(s.GetBool(name))
		case reflect.Slice:
			if value.Type().Elem().Kind() == reflect.String {
				slice := s.GetSlice(name)
				newSlice := reflect.MakeSlice(value.Type(), len(slice), len(slice))
				for i, s := range slice {
					newSlice.Index(i).SetString(s)
				}
				value.Set(newSlice)
			}
		case reflect.Map:
			if value.Type().Key().Kind() == reflect.String &&
				value.Type().Elem().Kind() == reflect.Slice &&
				value.Type().Elem().Elem().Kind() == reflect.String {
				m := s.GetMap(name)
				newMap := reflect.MakeMap(value.Type())
				for k, v := range m {
					sliceValue := reflect.MakeSlice(value.Type().Elem(), len(v), len(v))
					for i, s := range v {
						sliceValue.Index(i).SetString(s)
					}
					newMap.SetMapIndex(reflect.ValueOf(k), sliceValue)
				}
				value.Set(newMap)
			}
		}
	}
}
