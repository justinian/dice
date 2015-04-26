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
	open - regexp.MustCompile(`([0-9]+)d([0-9]+)(e)?o$`)
	sil - regexp.MustCompile(`([0-9]+)d([0-9]+)(e)?s$`)
*/

func Roll(desc string) (fmt.Stringer, error) {
	for _, rollHandler := range rollHandlers {
		if r := rollHandler.Pattern().FindStringSubmatch(desc); r != nil {
			return rollHandler.Roll(r)
		}
	}

	return nil, errors.New("Bad roll format: " + desc)
}
