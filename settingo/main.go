package settingo

import (
	"flag"
	"os"
	"strconv"
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
	msg        map[string]string
	VarString  map[string]string
	VarInt     map[string]int
	VarBool    map[string]bool
	VarMap     map[string]map[string][]string
	Parsers    map[string]func(string) string
	ParsersInt map[string]func(int) int
}

func (s *Settings) Set(flagName, defaultVar, message string) {
	s.msg[flagName] = message
	s.VarString[flagName] = defaultVar
}

func (s *Settings) SetString(flagName, defaultVar, message string) {
	s.Set(flagName, defaultVar, message)
}

func (s *Settings) SetInt(flagName string, defaultVar int, message string) {
	s.msg[flagName] = message
	s.VarInt[flagName] = defaultVar
}

func (s *Settings) SetBool(flagName string, defaultVar bool, message string) {
	s.msg[flagName] = message
	s.VarBool[flagName] = defaultVar
}

func (s *Settings) SetMap(flagName string, defaultVar map[string][]string, message string) {
	s.msg[flagName] = message
	s.VarMap[flagName] = defaultVar
}

func (s *Settings) SetParsed(flagName, defaultVar, message string, parserFunc func(string) string) {
	s.msg[flagName] = message
	s.VarString[flagName] = defaultVar
	s.Parsers[flagName] = parserFunc
}

func (s *Settings) SetParsedInt(flagName, defaultVar, message string, parserFunc func(int) int) {
	s.msg[flagName] = message
	s.VarString[flagName] = defaultVar
	s.ParsersInt[flagName] = parserFunc
}

func (s Settings) Get(flagName string) string {
	return s.VarString[flagName]
}

func (s Settings) GetInt(flagName string) int {
	return s.VarInt[flagName]
}

func (s Settings) GetBool(flagName string) bool {
	return s.VarBool[flagName]
}

func (s Settings) GetMap(flagName string) map[string][]string {
	return s.VarMap[flagName]
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
}

func (s *Settings) HandleOSInput() {
	for key := range s.VarString {
		varEnv, found := os.LookupEnv(key)
		if found {
			s.VarString[key] = varEnv
		}
	}
	for key := range s.VarInt {
		varEnv, found := os.LookupEnv(key)
		if found {
			if num, err := strconv.Atoi(varEnv); err == nil {
				s.VarInt[key] = num
			}
		}
	}
	for key := range s.VarBool {
		varEnv, found := os.LookupEnv(key)
		if found {
			s.VarBool[key] = truthiness(varEnv)
		}
	}
	for key := range s.VarMap {
		varEnv, found := os.LookupEnv(key)
		if found {
			s.VarMap[key] = ParseLineToMap(varEnv)
		}
	}
}

func (s *Settings) Parse() {
	s.HandleOSInput()
	s.HandleCMDLineInput()
}

var SETTINGS = Settings{
	msg:        make(map[string]string),
	VarString:  make(map[string]string),
	VarInt:     make(map[string]int),
	VarMap:     make(map[string]map[string][]string),
	Parsers:    make(map[string]func(string) string),
	ParsersInt: make(map[string]func(int) int),
	VarBool:    make(map[string]bool),
}

func Get(x string) string {
	return SETTINGS.Get(x)
}
func Set(flagName, defaultVar, message string) {
	SETTINGS.Set(flagName, defaultVar, message)
}

func SetString(flagName, defaultVar, message string) {
	SETTINGS.Set(flagName, defaultVar, message)
}

func SetInt(flagName string, defaultVar int, message string) {
	SETTINGS.SetInt(flagName, defaultVar, message)
}

func SetBool(flagName string, defaultVar bool, message string) {
	SETTINGS.SetBool(flagName, defaultVar, message)
}

func SetMap(flagName string, defaultVar map[string][]string, message string) {
	SETTINGS.SetMap(flagName, defaultVar, message)
}

func SetParsed(flagName, defaultVar, message string, parserFunc func(string) string) {
	SETTINGS.SetParsed(flagName, defaultVar, message, parserFunc)
}

func SetParsedInt(flagName, defaultVar, message string, parserFunc func(int) int) {
	SETTINGS.SetParsedInt(flagName, defaultVar, message, parserFunc)
}

func GetInt(flagName string) int {
	return SETTINGS.GetInt(flagName)
}

func GetBool(flagName string) bool {
	return SETTINGS.GetBool(flagName)
}

func GetMap(flagName string) map[string][]string {
	return SETTINGS.GetMap(flagName)
}

func Parse() {
	SETTINGS.Parse()
}
