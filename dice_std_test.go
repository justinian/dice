package dice_test

import (
	"testing"

	. "github.com/justinian/dice"
	. "gopkg.in/check.v1"
)

/* =============================================================================
 * Std dice test suite
 */

func TestStd(t *testing.T) { TestingT(t) }

type stdSuite struct {
}

/*
func (s *DiceSuite) SetUpSuite(c *C) {
}

func (s *DiceSuite) SetUpTest(c *C) {
}
*/

var _ = Suite(&stdSuite{})

/* =============================================================================
 * Std dice test cases
 */

func (s *stdSuite) TestBounds(c *C) {
	var roller StdRoller

	r, err := roller.Roll([]string{"100d2", "100", "2", ""})
	res := r.(StdResult)

	c.Assert(err, IsNil)

	for _, v := range res.Rolls {
		if v <= 0 || v > 2 {
			c.Errorf("Rolled out of bounds on a d2: %d", v)
		}
	}
}

func (s *stdSuite) TestCount(c *C) {
	var roller StdRoller

	r, err := roller.Roll([]string{"100d2", "100", "2", ""})
	res := r.(StdResult)

	c.Assert(err, IsNil)
	c.Assert(res.Rolls, HasLen, 100)
}
