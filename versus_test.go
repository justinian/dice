package dice_test

import (
	. "github.com/justinian/dice"
	. "gopkg.in/check.v1"
)

/* =============================================================================
 * vs dice test suite
 */

type vsSuite struct {
}

/*
func (s *DiceSuite) SetUpSuite(c *C) {
}

func (s *DiceSuite) SetUpTest(c *C) {
}
*/

var _ = Suite(&vsSuite{})

/* =============================================================================
 * vs dice test cases
 */

func (s *vsSuite) TestBounds(c *C) {
	var roller VsRoller

	r, err := roller.Roll([]string{"100d2v2", "100", "2", "", "2"})
	res := r.(VsResult)

	c.Assert(err, IsNil)

	for _, v := range res.Rolls {
		if v <= 0 || v > 2 {
			c.Errorf("Rolled out of bounds on a d2: %d", v)
		}
	}
}
