package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/egoholic/cfg"
)

var (
	err              error
	config           *Cfg
	boolArg          bool
	stringArg        string
	stringArrayArg   []string
	intArg           int
	intArrayArg      []int
	defaultStringArg string
)

func init() {
	defaults := map[string]interface{}{}
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
	intArrayArg, err = config.IntArrayArg("Integer Array Arg", "Test integer array argument.", "ints")
	if err != nil {
		panic(err)
	}
	stringArrayArg, err = config.StringArrayArg("String Array Arg", "Test string aarray rgument.", "strings")
	if err != nil {
		panic(err)
	}
	defaultStringArg, err = config.StringArg("Default string Arg", "Test default string argument.", "default")
	if err != nil {
		panic(err)
	}
}

func main() {
	var ints []string
	for _, i := range intArrayArg {
		ints = append(ints, strconv.Itoa(i))
	}
	fmt.Printf("BoolArg: %t\nStringArg: %s\nIntArg: %d\nDefaultStringArg: %s\nStringArrayArg: %s\nIntArrayArg: %s\n\n", boolArg, stringArg, intArg, defaultStringArg, strings.Join(stringArrayArg, ","), ints)
	config.AddHelpCommand()
}
