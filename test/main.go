package main

import (
	"fmt"

	. "github.com/egoholic/cfg"
)

var (
	err              error
	config           *Cfg
	boolArg          bool
	stringArg        string
	intArg           int
	defaultStringArg string
)

func init() {
	defaults := map[string]string{}
	defaults["default"] = "Default Value"
	config = Config(defaults)
	boolArg, err = config.BoolArg("Bool Arg", "Test boolean argument.", "bool_arg")
	if err != nil {
		panic(err)
	}
	stringArg, err = config.StringArg("String Arg", "Test string argument.", "string")
	if err != nil {
		panic(err)
	}
	intArg, err = config.IntArg("Integer Arg", "Test integer argument.", "int")
	if err != nil {
		panic(err)
	}
	defaultStringArg, err = config.StringArg("Default string Arg", "Test default string argument.", "default")
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Printf("BoolArg: %t\nStringArg: %s\nIntArg: %d\nDefaultStringArg: %s\n\n", boolArg, stringArg, intArg, defaultStringArg)
	config.AddHelpCommand()
}
