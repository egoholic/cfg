package cfg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/egoholic/cfg/doc"
	"github.com/egoholic/cfg/multikey"
)

var TrueValues = []string{"true", "1", "enabled", "on", "yes", "y", "Y"}
var FalseValues = []string{"false", "0", "disabled", "off", "no", "n", "N"}

type Cfg struct {
	defaults map[string]interface{} // string or []string
	docs     map[string]string
}

func Config(defaults map[string]interface{}) (cfg *Cfg) {
	return &Cfg{
		defaults: defaults,
		docs:     map[string]string{},
	}
}

func (cfg *Cfg) AddHelpCommand() {
	helpCmd, _ := cfg.CommandArg("Help", "Presents help information.", "help")
	if helpCmd {
		fmt.Print(cfg.Documentation())
		os.Exit(0)
	}
}

func (cfg *Cfg) Documentation() string {
	var sb strings.Builder
	for _, fragment := range cfg.docs {
		sb.WriteString(fragment)
	}
	return sb.String()
}

func (cfg *Cfg) add(n, doc string) error {
	_, found := cfg.docs[n]
	if found {
		return fmt.Errorf("argument `%s` already defined", n)
	}
	cfg.docs[n] = doc
	return nil
}

func (cfg *Cfg) find(mk multikey.MK) (val string, err error) {
	val = os.Getenv(mk.ENVVar)
	if len(val) > 0 {
		return
	}
	for i, arg := range os.Args {
		if arg == mk.Flag {
			val = os.Args[i+1]
			return
		}
	}
	raw, found := cfg.defaults[mk.Key]
	if !found {
		err = fmt.Errorf("argument '%s' not provided", mk.Key)
		return
	}
	val, ok := raw.(string)
	if !ok {
		err = fmt.Errorf("wrong type of default value, should be: 'string', got: %#v ", raw)
		return
	}
	return
}

func (cfg *Cfg) findMany(mk multikey.MK) (val []string, err error) {
	valStr := os.Getenv(mk.ENVVar)
	if len(valStr) > 0 {
		val = strings.Split(valStr, ":")
		return
	}
	foundIdx := 0 // not found
	for i, arg := range os.Args {
		if arg == mk.Flag {
			foundIdx = i
		}
	}
	for _, arg := range os.Args[foundIdx+1:] {
		if arg[0] == '-' {
			return
		}
		val = append(val, arg)
	}
	if len(val) > 0 {
		return
	}
	raw, found := cfg.defaults[mk.Key]
	if !found {
		err = fmt.Errorf("argument '%s' not provided", mk.Key)
		return
	}
	val, ok := raw.([]string)
	if !ok {
		err = fmt.Errorf("wrong type of default value, should be: 'string', got: %#v ", raw)
		return
	}
	return
}

func (cfg *Cfg) IntArg(name, desc, key string) (intValue int, err error) {
	mk := multikey.New(key)
	d, err := doc.New(name, doc.Integer, desc, mk)
	if err != nil {
		return
	}
	cfg.add(d.Name, d.Documentation)
	raw, err := cfg.find(mk)
	if err != nil {
		return
	}
	return strconv.Atoi(raw)
}

func (cfg *Cfg) IntArrayArg(name, desc, key string) (intVals []int, err error) {
	mk := multikey.New(key)
	d, err := doc.New(name, doc.IntegersArray, desc, mk)
	if err != nil {
		return
	}
	cfg.add(d.Name, d.Documentation)
	raw, err := cfg.findMany(mk)
	if err != nil {
		return
	}
	var v int
	for _, rawElem := range raw {
		v, err = strconv.Atoi(rawElem)
		if err != nil {
			return
		}
		intVals = append(intVals, v)
	}
	return
}

func (cfg *Cfg) StringArg(name, desc, key string) (strVal string, err error) {
	mk := multikey.New(key)
	d, err := doc.New(name, doc.String, desc, mk)
	if err != nil {
		return
	}
	cfg.add(d.Name, d.Documentation)
	return cfg.find(mk)
}

func (cfg *Cfg) StringArrayArg(name, desc, key string) (strVals []string, err error) {
	mk := multikey.New(key)
	d, err := doc.New(name, doc.StringsArray, desc, mk)
	if err != nil {
		return
	}
	cfg.add(d.Name, d.Documentation)
	return cfg.findMany(mk)
}

func (cfg *Cfg) BoolArg(name, desc, key string) (boolValue bool, err error) {
	mk := multikey.New(key)
	d, err := doc.New(name, doc.Bool, desc, mk)
	if err != nil {
		return
	}
	cfg.add(d.Name, d.Documentation)
	raw, err := cfg.find(mk)
	if err != nil {
		return
	}
	for _, tv := range TrueValues {
		if raw == tv {
			return true, nil
		}
	}
	for _, fv := range FalseValues {
		if raw == fv {
			return false, nil
		}
	}
	return false, fmt.Errorf("bool argument '%s' not provided, got: '%s'", name, raw)
}

func (cfg *Cfg) CommandArg(name, desc, key string) (bool, error) {
	mk := multikey.New(key)
	d, err := doc.New(name, doc.Command, desc, mk)
	if err != nil {
		return false, err
	}
	err = cfg.add(d.Name, d.Documentation)
	if err != nil {
		return false, err
	}
	if len(os.Args) > 1 && os.Args[1] == mk.Key {
		return true, nil
	}
	return false, nil
}
