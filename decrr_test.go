package decrr_test

import (
	"errors"
	"os"
	"runtime"
	"testing"

	"github.com/aldy505/decrr"
)

func TestWrap(t *testing.T) {
	x := errors.New("why does a chicken cross the road?")
	w := decrr.Wrap(x)
	wd, _ := os.Getwd()

	gt := runtime.GOROOT()

	var expected string
	expected += "why does a chicken cross the road?\n\n"
	expected += "github.com/aldy505/decrr_test.TestWrap " + wd + "/decrr_test.go:14\n"
	expected += "testing.tRunner " + gt + "/src/testing/testing.go:1259\n"
	expected += "runtime.goexit " + gt + "/src/runtime/asm_amd64.s:1581"

	if w.Error() != expected {
		t.Error("different than expected:", w.Error(), "\nexpected:", expected)
	}
}
