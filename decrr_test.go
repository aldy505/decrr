package decrr_test

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/aldy505/decrr"
)

func TestWrap(t *testing.T) {
	x := errors.New("why does a chicken cross the road?")
	w := decrr.Wrap(x)
	wd, _ := os.Getwd()

	cmd, err := exec.Command("bash", "-c", "go env | grep GOROOT").CombinedOutput()
	if err != nil {
		t.Error(err)
	}

	gt := strings.Replace(strings.Replace(strings.Replace(string(cmd), "GOROOT=", "", 1), "\"", "", 2), "\n", "", -1)

	var expected string
	expected += "why does a chicken cross the road?\n\n"
	expected += "github.com/aldy505/decrr_test.TestWrap " + wd + "/decrr_test.go:15\n"
	expected += "testing.tRunner " + gt + "/src/testing/testing.go:1259\n"
	expected += "runtime.goexit " + gt + "/src/runtime/asm_amd64.s:1581"

	if w.Error() != expected {
		t.Error("different than expected:", w.Error(), "\nexpected:", expected)
	}
}
