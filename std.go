package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

type StdRoller struct{}

var stdPattern = regexp.MustCompile(`([0-9]+)d([0-9]+)([+-][0-9]+)?($|\s)`)

func (StdRoller) Pattern() *regexp.Regexp { return stdPattern }

type StdResult struct {
	basicRollResult
	Rolls []int
	Total int
}

func (r StdResult) String() string {
	return fmt.Sprintf("%d %v", r.Total, r.Rolls)
}

func (StdRoller) Roll(matches []string) (RollResult, error) {
	dice, err := strconv.ParseInt(matches[1], 10, 0)
	if err != nil {
		return nil, err
	}

	sides, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return nil, err
	}

	result := StdResult{
		basicRollResult: basicRollResult{matches[0]},
		Rolls:           make([]int, dice),
		Total:           0,
	}

	if matches[3] != "" {
		bonus, err := strconv.ParseInt(matches[3], 10, 0)
		if err != nil {
			return nil, err
		}
		result.Total += int(bonus)
	}

	for i := int64(0); i < dice; i++ {
		roll := rand.Intn(int(sides)) + 1
		result.Total += roll
		result.Rolls[i] = roll
	}
	return result, nil
}

func init() {
	addRollHandler(StdRoller{})
}
