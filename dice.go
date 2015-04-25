package dice

import (
	"errors"
	"fmt"
	"regexp"
)

type roller interface {
	Pattern() *regexp.Regexp
	Roll([]string) (fmt.Stringer, error)
}

var rollHandlers []roller

func addRollHandler(handler roller) {
	rollHandlers = append(rollHandlers, handler)
}

/*
	{regexp.MustCompile(`([0-9]+)d([0-9]+)(e)?v([0-9]+)$`), versusRoll},
	{regexp.MustCompile(`([0-9]+)d([0-9]+)(e)?o$`), openRoll},
	{regexp.MustCompile(`([0-9]+)d([0-9]+)(e)?s$`), silRoll},
	{regexp.MustCompile(`([0-9]+(?:r|b|blk|p|g|y|w)\s*)+`), eoteRoll},
}
*/

func Roll(desc string) (fmt.Stringer, error) {
	for _, rollHandler := range rollHandlers {
		if r := rollHandler.Pattern().FindStringSubmatch(desc); r != nil {
			return rollHandler.Roll(r)
		}
	}

	return nil, errors.New("Bad roll format: " + desc)
}
