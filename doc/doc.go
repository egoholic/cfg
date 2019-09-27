package doc

import (
	"fmt"
	"strings"

	"github.com/egoholic/cfg/multikey"
)

const (
	Integer       = "Integer"
	IntegersArray = "Array of Integers"
	String        = "String"
	StringsArray  = "Array of Strings"
	Bool          = "Bool"
	Command       = "Command"
)

var ArgTypes = []string{Integer, IntegersArray, String, StringsArray, Bool, Command}

type Doc struct {
	Name          string
	Documentation string
}

func New(name, typ, desc string, mk multikey.MK) (doc Doc, err error) {
	if len(name) < 4 {
		err = fmt.Errorf("`name` should be at least 4 chars long, got: `%s` (len: %d)", name, len(name))
		return
	}
	err = checkTyp(typ)
	if err != nil {
		return
	}
	if len(desc) < 16 {
		err = fmt.Errorf("`description` should be at least 16 chars long, got: `%s` (len: %d)", desc, len(desc))
		return
	}
	var sb strings.Builder
	sb.WriteString("\n\tName: ")
	sb.WriteString(name)
	sb.WriteString(" (of type: '")
	sb.WriteString(typ)
	sb.WriteString("')\n\n")
	sb.WriteString("\t\tDescription: ")
	sb.WriteString(desc)
	sb.WriteString("\n\n")
	if typ != Command {
		sb.WriteString("\t\tENV Var:    ")
		sb.WriteString(mk.ENVVar)
		sb.WriteRune('\n')
		sb.WriteString("\t\tFlag:       ")
		sb.WriteString(mk.Flag)
		sb.WriteRune('\n')
	}
	sb.WriteString("\t\tKey:        ")
	sb.WriteString(mk.Key)
	sb.WriteString("\n\n\n")
	doc = Doc{
		Name:          name,
		Documentation: sb.String(),
	}
	return
}

func checkTyp(typ string) error {
	for _, argT := range ArgTypes {
		if argT == typ {
			return nil
		}
	}
	return fmt.Errorf("wrong command type, should be one of: [%s], got: %s", strings.Join(ArgTypes, ", "), typ)
}
