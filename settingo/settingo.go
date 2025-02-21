package settingo

// SETTINGS is the global instance of the Settings struct for the settingo package.
//
// It provides a package-level access point to manage application settings.
// Use the package functions (e.g., settingo.Set, settingo.Get, settingo.Parse)
// to interact with this global settings registry.
//
// Initialization:
//
//	SETTINGS is initialized with empty maps for all setting types when the package is loaded.
//	Settings are registered and accessed through the package-level functions.
//
// Example:
//
//	package main
//
//	import "path/to/settingo"
//
//	func main() {
//		settingo.SetString("outputDir", "/tmp", "Directory to write output files")
//		settingo.Parse()
//		outputDir := settingo.Get("outputDir")
//		println("Output directory:", outputDir)
//	}
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

// Get retrieves the current string value of a registered string setting from the global SETTINGS instance.
//
// It's a package-level function that delegates to the Get method of the global SETTINGS variable.
//
// Args:
//
//	x: The name of the setting flag to retrieve.
//
// Returns:
//
//	The current string value of the setting from the global SETTINGS instance.
func Get(x string) string {
	return SETTINGS.Get(x)
}

// Set is a package-level function to register a string setting within the global SETTINGS instance.
//
// It delegates to the Set method of the global SETTINGS variable.
// It associates a setting flag name with a default string value and a help message.
//
// Args:
//
//	flagName:   The name of the setting flag (e.g., "port", "output-dir").
//	defaultVar: The default string value for the setting.
//	message:    The help message describing the setting's purpose.
func Set(flagName, defaultVar, message string) {
	SETTINGS.Set(flagName, defaultVar, message)
}

// SetString is a package-level function to register a string setting within the global SETTINGS instance.
//
// It delegates to the SetString method of the global SETTINGS variable.
// Registers a string setting that can be configured via environment variables or command-line flags.
//
// Args:
//
//	flagName:   The name of the setting flag (e.g., "output-format").
//	            This name will be used for both environment variable lookup
//	            (case-sensitive, typically uppercase) and command-line flag parsing
//	            (lowercase with hyphens).
//	defaultVar: The default string value if no environment variable or
//	            command-line flag is provided.
//	message:    The help message describing the setting's purpose, shown in
//	            command-line help output.
//
// Example:
//
//		settingo.SetString("outputFormat", "json", "Format for output (json|xml)")
//
//	 // Can be set via:
//	 // - Environment variable: OUTPUTFORMAT=xml
//	 // - Command-line flag: --output-format=xml
func SetString(flagName, defaultVar, message string) {
	SETTINGS.Set(flagName, defaultVar, message)
}

// SetInt is a package-level function to register an integer setting within the global SETTINGS instance.
//
// It delegates to the SetInt method of the global SETTINGS variable.
// Registers an integer setting that can be configured via environment variables or command-line flags.
//
// Args:
//
//	flagName:   The name of the setting flag (e.g., "port").
//	            Used for environment variable lookup and command-line flag parsing.
//	defaultVar: The default integer value.
//	message:    The help message.
//
// Example:
//
//		settingo.SetInt("port", 8080, "Port to listen on")
//
//	 // Can be set via:
//	 // - Environment variable: PORT=8081
//	 // - Command-line flag: --port=8081
func SetInt(flagName string, defaultVar int, message string) {
	SETTINGS.SetInt(flagName, defaultVar, message)
}

// SetBool is a package-level function to register a boolean setting within the global SETTINGS instance.
//
// It delegates to the SetBool method of the global SETTINGS variable.
// Registers a boolean setting that can be configured via environment variables or command-line flags.
//
// When set via environment variables or command-line flags, the string values
// are interpreted as boolean using the truthiness function (see truthiness()).
//
// Args:
//
//	flagName:   The name of the setting flag (e.g., "verbose").
//	            Used for environment variable lookup and command-line flag parsing.
//	defaultVar: The default boolean value.
//	message:    The help message.
//
// Example:
//
//		settingo.SetBool("verbose", false, "Enable verbose output")
//
//	 // Can be set via:
//	 // - Environment variable: VERBOSE=true or VERBOSE=y or VERBOSE=yes
//	 // - Command-line flag: --verbose=true or --verbose=y or --verbose=yes
func SetBool(flagName string, defaultVar bool, message string) {
	SETTINGS.SetBool(flagName, defaultVar, message)
}

// SetMap is a package-level function to register a map setting within the global SETTINGS instance.
//
// It delegates to the SetMap method of the global SETTINGS variable.
// Registers a map setting with string keys and string slice values
// that can be configured via environment variables or command-line flags.
//
// The map is parsed from a string representation in the format "key1:value1,value2;key2:value3".
// Keys and values are separated by colons, key-value pairs by semicolons, and values within a key by commas.
//
// Args:
//
//	flagName:   The name of the setting flag (e.g., "headers").
//	            Used for environment variable lookup and command-line flag parsing.
//	defaultVar: The default map value.
//	message:    The help message.
//
// Example:
//
//		defaultHeaders := map[string][]string{"Content-Type": {"application/json"}, "Accept": {"application/json"}}
//		settingo.SetMap("headers", defaultHeaders, "HTTP headers to include")
//
//	 // Can be set via:
//	 // - Environment variable: HEADERS="Content-Type:application/json,text/plain;Accept:application/json"
//	 // - Command-line flag: --headers="Content-Type:application/json,text/plain;Accept:application/json"
func SetMap(flagName string, defaultVar map[string][]string, message string) {
	SETTINGS.SetMap(flagName, defaultVar, message)
}

// SetSlice is a package-level function to register a string slice setting within the global SETTINGS instance.
//
// It delegates to the SetSlice method of the global SETTINGS variable.
// Registers a string slice setting that can be configured via environment variables or command-line flags.
//
// The slice is parsed from a string by splitting it using the provided separator.
// If no separator is provided, it defaults to a comma ",".
//
// Args:
//
//	flagName:   The name of the setting flag (e.g., "hosts").
//	            Used for environment variable lookup and command-line flag parsing.
//	defaultVar: The default slice value.
//	message:    The help message.
//	sep:        The separator string used to split the environment variable or
//	            command-line flag value into a slice of strings. If empty, defaults to ",".
//
// Example:
//
//		defaultHosts := []string{"localhost", "127.0.0.1"}
//		settingo.SetSlice("hosts", defaultHosts, "List of hosts", ",")
//
//	 // Can be set via:
//	 // - Environment variable: HOSTS="host1,host2,host3"
//	 // - Command-line flag: --hosts="host1,host2,host3"
//	 // or using a different separator:
//	 // settingo.SetSlice("ports", []string{"80", "81"}, "Ports to listen on", ";")
//	 // - Environment variable: PORTS="80;81;82"
//	 // - Command-line flag: --ports="80;81;82"
func SetSlice(flagName string, defaultVar []string, message string, sep string) {
	SETTINGS.SetSlice(flagName, defaultVar, message, sep)
}

// SetParsed is a package-level function to register a string setting with a custom parsing function within the global SETTINGS instance.
//
// It delegates to the SetParsed method of the global SETTINGS variable.
// Registers a string setting with a custom parsing function.
//
// This is useful when you need to perform custom validation or transformation
// on the string value obtained from environment variables or command-line flags
// before using it in your application.
//
// Args:
//
//	flagName:   The name of the setting flag.
//	defaultVar: The default string value.
//	message:    The help message.
//	parserFunc: A function that takes the raw string value (from env or flag)
//	            and returns a parsed string value. This function will be called
//	            after retrieving the value from the environment or command line.
//
// Example:
//
//	settingo.SetParsed("username", "default", "Username", func(s string) string {
//		if s == "" {
//			return "anonymous"
//		}
//		return strings.ToLower(s)
//	})
func SetParsed(flagName, defaultVar, message string, parserFunc func(string) string) {
	SETTINGS.SetParsed(flagName, defaultVar, message, parserFunc)
}

// SetParsedInt is a package-level function to register an integer setting with a custom parsing function within the global SETTINGS instance.
//
// It delegates to the SetParsedInt method of the global SETTINGS variable.
// Registers an integer setting with a custom parsing function.
//
// This is useful for custom validation or transformation of integer settings.
// Note that while the setting is treated as an integer internally, the defaultVar
// is still a string for consistency with command-line flag handling, and the
// raw value from environment or flag parsing is initially a string.
//
// Args:
//
//	flagName:   The name of the setting flag.
//	defaultVar: The default string value (will be converted to int if possible).
//	message:    The help message.
//	parserFunc: A function that takes the raw integer value (after initial parsing from string)
//	            and returns a parsed integer value.  Note that the input to this function
//	            is an int, but the original input from environment/flag was a string.
//
// Example:
//
//	settingo.SetParsedInt("retries", "3", "Number of retries", func(i int) int {
//		if i < 0 {
//			return 0 // Ensure retries is not negative
//		}
//		return i
//	})
func SetParsedInt(flagName, defaultVar, message string, parserFunc func(int) int) {
	SETTINGS.SetParsedInt(flagName, defaultVar, message, parserFunc)
}

// GetInt retrieves the current integer value of a registered integer setting from the global SETTINGS instance.
//
// It's a package-level function that delegates to the GetInt method of the global SETTINGS variable.
//
// Args:
//
//	flagName: The name of the setting flag to retrieve.
//
// Returns:
//
//	The current integer value of the setting from the global SETTINGS instance.
func GetInt(flagName string) int {
	return SETTINGS.GetInt(flagName)
}

// GetBool retrieves the current boolean value of a registered boolean setting from the global SETTINGS instance.
//
// It's a package-level function that delegates to the GetBool method of the global SETTINGS variable.
//
// Args:
//
//	flagName: The name of the setting flag to retrieve.
//
// Returns:
//
//	The current boolean value of the setting from the global SETTINGS instance.
func GetBool(flagName string) bool {
	return SETTINGS.GetBool(flagName)
}

// GetMap retrieves the current map value of a registered map setting from the global SETTINGS instance.
//
// It's a package-level function that delegates to the GetMap method of the global SETTINGS variable.
//
// Args:
//
//	flagName: The name of the setting flag to retrieve.
//
// Returns:
//
//	The current map value of the setting from the global SETTINGS instance.
func GetMap(flagName string) map[string][]string {
	return SETTINGS.GetMap(flagName)
}

// GetSlice retrieves the current slice value of a registered slice setting from the global SETTINGS instance.
//
// It's a package-level function that delegates to the GetSlice method of the global SETTINGS variable.
//
// Args:
//
//	flagName: The name of the setting flag to retrieve.
//
// Returns:
//
//	The current slice value of the setting from the global SETTINGS instance.
func GetSlice(flagName string) []string {
	return SETTINGS.GetSlice(flagName)
}

// Parse parses settings from both OS environment variables and command-line flags using the global SETTINGS instance.
//
// It's a package-level function that delegates to the Parse method of the global SETTINGS variable.
//
// It first calls HandleOSInput() to parse settings from environment variables,
// and then calls HandleCMDLineInput() to parse command-line flags.
//
// This order ensures that command-line flags take precedence over environment
// variables if a setting is defined in both sources.
//
// Call this function after registering all settings using the Set... functions
// to populate the global SETTINGS instance with values from the environment and command line.
func Parse() {
	SETTINGS.Parse()
}

// ParseTo parses settings from environment variables and command-line flags for the global SETTINGS instance
// and then updates the fields of the provided struct with the parsed values.
//
// It's a package-level function that delegates to the ParseTo method of the global SETTINGS variable.
//
// It performs the following steps:
//  1. LoadStruct(to): Registers settings based on the struct fields and "settingo" tags
//     of the input struct 'to'. This effectively declares the settings to be parsed
//     and associates them with struct fields.
//  2. Parse(): Parses settings from OS environment variables and command-line flags,
//     populating the global SETTINGS instance's internal storage.
//  3. UpdateStruct(to): Updates the fields of the input struct 'to' with the parsed
//     values from the global SETTINGS instance.
//
// This function simplifies the process of configuring an application by directly
// mapping settings to struct fields.
//
// Args:
//
//	to: A pointer to a struct whose fields will be updated with parsed settings.
//	    The struct's fields must be exported and can be tagged with "settingo"
//	    to provide help messages (see LoadStruct for tag usage in Settings type).
func ParseTo(to interface{}) {
	SETTINGS.ParseTo(to)
}
