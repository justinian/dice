package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

type stdRoller struct{}

var stdPattern = regexp.MustCompile(`([0-9]+)d([0-9]+)([+-][0-9]+)?$`)

func (stdRoller) Pattern() *regexp.Regexp { return stdPattern }

type stdResult struct {
	rolls []int
	total int
}

func (r stdResult) String() string {
	return fmt.Sprintf("%d", r.result)
}

func (stdRoller) Roll(matches []string) (fmt.Stringer, error) {
	dice, err := strconv.ParseInt(matches[1], 10, 0)
	if err != nil {
		return nil, err
	}

	sides, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return nil, err
	}

	result := stdResult{
		rolls: make([]int, dice),
		total: 0,
	}

	if matches[3] != "" {
		bonus, err := strconv.ParseInt(matches[3], 10, 0)
		if err != nil {
			return nil, err
		}
		result.total += int(bonus)
	}

	for i := int64(0); i < dice; i++ {
		roll := rand.Intn(sides)
		result.total += roll
		result.rolls[i] = roll
	}
	return result, nil
}

func init() {
	addRollHandler(stdRoller{})
}
