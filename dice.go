package dice

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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

func Roll(desc string) (fmt.Stringer, string, error) {
	for _, rollHandler := range rollHandlers {
		rollHandler.Pattern().Longest()

		if r := rollHandler.Pattern().FindStringSubmatch(desc); r != nil {
			result, err := rollHandler.Roll(r)
			if err != nil {
				return nil, "", err
			}

			indexes := rollHandler.Pattern().FindStringSubmatchIndex(desc)
			reason := strings.Trim(desc[indexes[0]+len(r[0]):], " \t\r\n")
			return result, reason, nil
		}
	}

	return nil, "", errors.New("Bad roll format: " + desc)
}
