package git

import "testing"

func TestListSubmodule(t *testing.T) {
	_, err := ListSubmodule()
	if err != nil {
		t.FailNow()
	}
}
