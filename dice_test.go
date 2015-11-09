package dice_test

import (
	"testing"

	. "github.com/justinian/dice"
)

func TestRoll(t *testing.T) {
	roll := "3d8v3"
	res, _, _ := Roll(roll)
	if _, ok := res.(VsResult); !ok {
		t.Fatalf("%s is not a VsResult", roll)
	}

	roll = "3d8+2test"
	_, _, err := Roll(roll)
	if err != nil {
		t.Logf("err '%v' properly detected in %s", err, roll)
	} else {
		t.Fatalf("err not detected in %s", roll)
	}

	roll = "3b4bl"
	_, reason, err := Roll(roll)
	if reason == "4bl" {
		t.Fatalf("malformed dice format read as reason, %s", roll)
	}
	if err != nil {
		t.Logf("err '%v' properly detected in %s", err, roll)
	}

	roll = "9d9rv5"
	res, _, _ = Roll(roll)
	if _, ok := res.(VsResult); !ok {
		t.Fatalf("%s is not a VsResult", roll)
	}
}
