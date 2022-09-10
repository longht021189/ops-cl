package exe

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	data := "asdasdasdsd"
	out, err := Run("echo", data)

	if err != nil {
		t.FailNow()
	}

	outStr := strings.TrimSpace(string(out))
	if data != outStr {
		t.FailNow()
	}
}
