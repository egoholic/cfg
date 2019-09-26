package doc

import (
	"fmt"
	"strings"

	"github.com/egoholic/cfg/multikey"
)

const (
	Integer = "Integer"
	String  = "String"
	Bool    = "Bool"
	Command = "Command"
)

type Doc struct {
	Name          string
	Documentation string
}

func New(name, typ, desc string, mk multikey.MK) (doc Doc, err error) {
	if len(name) < 4 {
		err = fmt.Errorf("`name` should be at least 4 chars long, got: `%s` (len: %d)", name, len(name))
		return
	}
	if typ != Integer && typ != String && typ != Bool && typ != Command {
		err = fmt.Errorf("`typ` should be one of: `%s`,`%s`,`%s`,`%s`, got: %s", Integer, String, Bool, Command, typ)
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
