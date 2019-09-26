package multikey

import "strings"

type MK struct {
	Key    string
	ENVVar string
	Flag   string
}

func New(base string) MK {
	base = strings.ToLower(base)
	return MK{
		Key:    base,
		ENVVar: strings.ToUpper(base),
		Flag:   "-" + base,
	}
}
