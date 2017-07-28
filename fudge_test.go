package dice_test

import (
	"testing"

	. "github.com/justinian/dice"
	. "gopkg.in/check.v1"
)

/* =============================================================================
 * Fudge dice test suite
 */

func TestFudge(t *testing.T) { TestingT(t) }

type fudgeSuite struct{}

var _ = Suite(&fudgeSuite{})

/* =============================================================================
 * Std dice test cases
 */

func (s *fudgeSuite) TestBounds(c *C) {
	var roller FudgeRoller

	r, err := roller.Roll([]string{"100df", "100", "", ""})
	res := r.(FudgeResult)

	c.Assert(err, IsNil)

	for _, v := range res.Rolls {
		if v < -1 || v > +1 {
			c.Errorf("Rolled out of bounds on a dF: %d", v)
		}
	}
}

func (s *fudgeSuite) TestCount(c *C) {
	var roller FudgeRoller

	r, err := roller.Roll([]string{"100df", "100", "", ""})
	res := r.(FudgeResult)

	c.Assert(err, IsNil)
	c.Assert(res.Rolls, HasLen, 100)
}

func (s *fudgeSuite) TestAdd(c *C) {
	var roller FudgeRoller

	r, err := roller.Roll([]string{"1df+13", "1", "+13"})
	res := r.(FudgeResult)

	if res.Total < 12 || res.Total > 14 {
		c.Errorf("Added roll out of bounds on a dF: %d", res.Total)
	}

	c.Assert(err, IsNil)
	c.Assert(res.Rolls, HasLen, 1)
}

func (s *fudgeSuite) TestSubtract(c *C) {
	var roller FudgeRoller

	r, err := roller.Roll([]string{"1df+13", "1", "-13"})
	res := r.(FudgeResult)

	if res.Total < -14 || res.Total > -12 {
		c.Errorf("Added roll out of bounds on a dF: %d", res.Total)
	}

	c.Assert(err, IsNil)
	c.Assert(res.Rolls, HasLen, 1)
}
