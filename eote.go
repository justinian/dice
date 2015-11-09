package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type EoteResult struct {
	basicRollResult
	S int // success
	A int // advantage
	T int // triumph
	D int // despair
	F int // force
}

func (r *EoteResult) Add(o EoteResult) {
	r.S += o.S
	r.A += o.A
	r.T += o.T
	r.D += o.D
	r.F += o.F
}

func (r EoteResult) String() string {
	parts := make([]string, 0, 5)

	if r.S > 0 {
		parts = append(parts, fmt.Sprintf("%d success", r.S))
	} else if r.S < 0 {
		parts = append(parts, fmt.Sprintf("%d failure", -r.S))
	}

	if r.A > 0 {
		parts = append(parts, fmt.Sprintf("%d advantage", r.A))
	} else if r.A < 0 {
		parts = append(parts, fmt.Sprintf("%d disadvantage", -r.A))
	}

	if r.T > 0 {
		parts = append(parts, fmt.Sprintf("(%d triumph)", r.T))
	}

	if r.D > 0 {
		parts = append(parts, fmt.Sprintf("(%d despair)", r.D))
	}

	if r.F > 0 {
		parts = append(parts, fmt.Sprintf("%d light", r.F))
	} else if r.F < 0 {
		parts = append(parts, fmt.Sprintf("%d dark", -r.F))
	}

	if len(parts) < 1 {
		parts = append(parts, "no result")
	}

	return strings.Join(parts, " ")
}

var eoteDice = map[string][]EoteResult{
	"b":   {{}, {}, {A: 1}, {A: 2}, {S: 1}, {S: 1, A: 1}},
	"blk": {{}, {}, {A: -1}, {A: -1}, {S: -1}, {S: -1}},
	"g":   {{}, {A: 1}, {A: 1}, {A: 2}, {S: 1}, {S: 1}, {S: 1, A: 1}, {S: 2}},
	"p":   {{}, {A: -1}, {A: -1}, {A: -1}, {A: -2}, {S: -1}, {S: -1, A: -1}, {S: -2}},
	"y":   {{}, {A: 1}, {A: 2}, {A: 2}, {S: 1}, {S: 1}, {S: 1, A: 1}, {S: 1, A: 1}, {S: 1, A: 1}, {S: 2}, {S: 2}, {S: 1, T: 1}},
	"r":   {{}, {A: -1}, {A: -1}, {A: -2}, {A: -2}, {S: -1}, {S: -1}, {S: -1, A: -1}, {S: -1, A: -1}, {S: -2}, {S: -2}, {S: -1, D: 1}},
	"w":   {{F: -1}, {F: -1}, {F: -1}, {F: -1}, {F: -1}, {F: -1}, {F: -2}, {F: 1}, {F: 1}, {F: 2}, {F: 2}, {F: 2}},
}

type EoteRoller struct{}

var eotePattern = regexp.MustCompile(`([0-9]+(?:r|b|blk|p|g|y|w)\s*)+($|\s)`)
var diePattern = regexp.MustCompile(`([0-9]+)(r|b|blk|p|g|y|w)`)

func (EoteRoller) Pattern() *regexp.Regexp { return eotePattern }

func (EoteRoller) Roll(matches []string) (RollResult, error) {
	diePattern.Longest()

	res := EoteResult{basicRollResult: basicRollResult{matches[0]}}

	for _, die := range strings.Split(matches[0], " ") {
		parts := diePattern.FindStringSubmatch(strings.Trim(die, " \t\r\n"))
		if parts == nil {
			continue
		}

		num, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			continue
		}

		choices, ok := eoteDice[parts[2]]
		if !ok {
			continue
		}

		for i := int64(0); i < num; i++ {
			res.Add(choices[rand.Intn(len(choices))])
		}
	}

	return res, nil
}

func init() {
	addRollHandler(EoteRoller{})
}
