package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
)

type FudgeRoller struct{}

var fudgePattern = regexp.MustCompile(`([0-9]+)df?([+-][0-9]+)?($|\s)`)

func (FudgeRoller) Pattern() *regexp.Regexp { return fudgePattern }

type FudgeResult struct {
	basicRollResult
	Rolls []int
	Total int
}

func (r FudgeResult) String() string {
	return fmt.Sprintf("%d %v", r.Total, r.Rolls)
}

func (r FudgeResult) Int() int {
	return r.Total
}

func (FudgeRoller) Roll(matches []string) (RollResult, error) {
	dice, err := strconv.ParseInt(matches[1], 10, 0)
	if err != nil {
		return nil, err
	}

	result := FudgeResult{
		basicRollResult: basicRollResult{matches[0]},
		Rolls:           make([]int, dice),
		Total:           0,
	}

	if matches[2] != "" {
		bonus, err := strconv.ParseInt(matches[2], 10, 0)
		if err != nil {
			return nil, err
		}
		result.Total += int(bonus)
	}

	for i := 0; i < len(result.Rolls); i++ {
		result.Rolls[i] = rand.Intn(3) - 1
		result.Total += result.Rolls[i]
	}

	sort.Ints(result.Rolls)

	return result, nil
}

func init() {
	addRollHandler(FudgeRoller{})
}
