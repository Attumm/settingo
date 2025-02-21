package settingo

var SETTINGS = Settings{
	msg:         make(map[string]string),
	VarString:   make(map[string]string),
	VarInt:      make(map[string]int),
	VarMap:      make(map[string]map[string][]string),
	VarSlice:    make(map[string][]string),
	VarSliceSep: make(map[string]string),
	Parsers:     make(map[string]func(string) string),
	ParsersInt:  make(map[string]func(int) int),
	VarBool:     make(map[string]bool),
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

func SetSlice(flagName string, defaultVar []string, message string, sep string) {
	SETTINGS.SetSlice(flagName, defaultVar, message, sep)
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

func GetSlice(flagName string) []string {
	return SETTINGS.GetSlice(flagName)
}

func Parse() {
	SETTINGS.Parse()
}
func ParseTo(to interface{}) {
	SETTINGS.ParseTo(to)
}
