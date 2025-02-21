# Settingo
[![GitHub release](https://img.shields.io/github/v/release/Attumm/settingo?sort=semver)](https://github.com/Attumm/settingo/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Attumm/settingo)](https://goreportcard.com/report/github.com/Attumm/settingo)
[![CI](https://github.com/Attumm/settingo/actions/workflows/ci.yml/badge.svg)](https://github.com/Attumm/settingo/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/Attumm/settingo/branch/main/graph/badge.svg)](https://codecov.io/gh/Attumm/settingo)
[![Go Reference](https://pkg.go.dev/badge/github.com/Attumm/settingo/settingo.svg)](https://pkg.go.dev/github.com/Attumm/settingo/settingo)

## _Settings should be simple, and with settingo it is._

Settingo parses command line and environment variables, all within one tool.
Setting up settings is as simple as creating a struct with helpful messages for your project and --help on the CLI.
```go
type Config struct {
    APIKey    string `settingo:"API key for authentication"`
}

func main()  {
    config := Config{
        APIKey:    "foo-bar",
    }
    settingo.ParseTo(config)
}
```

Now the struct will hold always hold value, either default, environment var, cli. based on the context making settings simple.
Letting you focus on your project.
```go
config.APIKey
```

## Features
- Simplicity: Set up settings within a single line of code.
- Flexibility: Utilize command-line flags, environment variables, or defaults.
- Typesafety: Seamlessly work with strings, integers, slices, booleans, and maps.
- Convenience: Global access with a singleton pattern.
- User-friendly: Automatic --help flag generation for your applications.
- Versatility: Works flawlessly in Linux, Docker, Kubernetes, and other environments.

## Example
example of how to use. More can be found in the [example_project](https://github.com/Attumm/settingo_example_project/blob/main/main.go)
```go
package main

import (
	"fmt"
	"github.com/Attumm/settingo/settingo"
)

// Define your configuration with various types and help messages
type Config struct {
	APIKey         string              `settingo:"API key for authentication"`
	Port           int                 `settingo:"Port to run the server on"`
	Verbose        bool                `settingo:"Enable verbose output"`
	Hosts          []string            `settingo:"List of allowed hosts (comma-separated)"`
	Items          []string            `settingo:"List of items (pipe-separated, sep=|)" settingo:"sep=|"`
	Headers        map[string][]string `settingo:"HTTP headers to include (key:value1,value2;key2:value3 format)"`
}

func main() {
	// Initialize config with default values
	config := &Config{
		APIKey:         "foo-bar",
		Port:           8080,
		Verbose:        true,
		Hosts:          []string{"localhost", "127.0.0.1"},
		Items:          []string{"alpha", "beta", "gamma"},
		Headers:        map[string][]string{"Accept": {"application/json"}},
	}

	// Parse command-line flags and environment variables into config
	settingo.ParseTo(config)

	// Print out the configuration values
	fmt.Println("Configuration:")
	fmt.Println("APIKey      =", config.APIKey)
	fmt.Println("Port        =", config.Port)
	fmt.Println("Verbose     =", config.Verbose)
	fmt.Println("Headers     =", config.Headers)
	fmt.Println("Hosts       =", config.Hosts)
	fmt.Println("Items       =", config.Items)
}
```
When you build your application (e.g., go build -o myapp) and run ./myapp --help, settingo automatically generates help text based on struct tags and default values:
```bash
Usage of ./myapp:
  -APIKEY string
        API key for authentication (default "foo-bar")
  -HEADERS string
        HTTP headers to include (key:value1,value2;key2:value3 format) (default "Accept:application/json")
  -HOSTS string
        List of allowed hosts (comma-separated) (default "localhost,127.0.0.1")
  -ITEMS string
        List of items (pipe-separated, sep=|) (default "alpha,beta,gamma")
  -PORT int
        Port to run the server on (default 8080)
  -VERBOSE string
        Enable verbose output (default "true")
```

```go
package main

import (
        "fmt"
        "github.com/Attumm/settingo/settingo"
)

func main() {
        settingo.Set("FOO", "default value", "handy help text")
        
        settingo.Parse()
        fmt.Println("foo =",  settingo.Get("FOO"))
}
```
The above go will produce binary that can be used as follows.
Get handy help text set in the above example on the same line.
This can get very handy when the project grows and is used in different environments
```sh
$ ./example --help
Usage of ./example:
  -FOO string
      handy help text (default "default value")
```

When no value is given, default value is used
```sh
$ ./example
foo = default value
```

Running the binary with command line input
```sh
$ ./example -FOO bar
foo = bar
```
Running the binary with environment variable
```sh
$ FOO=ok;./example
foo = ok
```

## Order of preference
variables are set with preference
variables on the command line will have highest preference.
This because while testing you might want to override environment
The priority order is as follows
1. Command line input
2. Environment variables 
3. Default values

## Example: Custom Parsing for "Messy" Input with `SetParsed`

Sometimes, environment variables or command-line arguments might not be perfectly formatted.  You might receive an empty string, mixed-case input, or data that needs transformation.  `settingo`'s `SetParsed` is ideal for cleaning up and standardizing such "messy" input.

This example demonstrates handling a `RAW_USERNAME` environment variable, ensuring the `Username` setting is always a lowercase, non-empty string, defaulting to "anonymous" if the input is blank:

```go
package main

import (
    "fmt"
	"github.com/Attumm/settingo/settingo"
    "strings"
)

// Define your configuration with Parsed setting
type Config struct {
    Username string `settingo:"USERNAME for application access"`
}

func main() {
    config := &Config{
        Username: "default",
    }

    // Use SetParsed to handle potentially messy Username input
    settingo.SetParsed("USERNAME", "default", "Username for application access", func(input string) string {
        if input == "" {
            return "anonymous" // Default to "anonymous" if empty input
        }
        return strings.ToLower(input) // Convert username to lowercase
    })
	
    settingo.ParseTo(config)
    fmt.Println("Configured Username:", config.Username)
}
```

```bash
$./example
Configured Username: default
$ ./example --USERNAME ''
Configured Username: anonymous
$ ./example --USERNAME FOOBAR
Configured Username: foobar
```

## installation
```bash
go get "github.com/Attumm/settingo/settingo"
```

## Example project
Handy [example_project](https://github.com/Attumm/settingo_example_project) as starting point.

## License
MIT
