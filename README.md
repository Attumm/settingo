# Settingo
## _Settings should be simple, and with settingo it is._

Settings parses command line and environment variables on one line.
And makes it available throughout the code base. Making using settings in your project as boring and unimportant as it should be.
Settings vars is as simple as:
```go
 settingo.Set("FOO", "default value", "help text")
```
Getting vars out has the same level of complexity as setting the value.
```go
 settingo.Get("FOO")
```


## Features
- Simple to use
- supports command line and environment variables 
- Support for types: str, int, bool, map
- Singleton, makes it easy to use in program anywhere in the code-base
- Supports help text with --help on the binary
- Ease of use in any environment examples: linux, docker, k8


## Example
example of how to use. More can be found in the [example_project](https://github.com/Attumm/settingo_example_project/blob/main/main.go)
```go
package main

import (
        "fmt"
        "github.com/Attumm/settingo/settingo"
)

func main() {
        settingo.Set("FOO", "default value", "handy help text")
        
        settingo.Parse()
        fmt.Println("foobar =",  settingo.Get("FOOBAR"))
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

## Types
Settingo supports different types.
```go
// string
settingo.Set("FOO", "default", "help text")
settingo.Get("FOO")

// integer
settingo.SetInt("FOO", 42, "help text")
settingo.GetInt("FOO")

// boolean
settingo.SetBool("FOO", true, "help text")
settingo.GetBool("FOO")

// map
defaultMap := make(map[string][]string)
defaultMap["foo"] = []string{"bar"}
settingo.SetMap("FOO", defaultMap, "help text")
settingo.GetMap("FOO")
```

## Example project
Handy [example_project](https://github.com/Attumm/example_settingo) as starting point.

## License

MIT

