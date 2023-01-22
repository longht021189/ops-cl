package docker

import (
	"testing"

	"github.com/longht021189/ops-cl/model/args"
)

func TestBuildScratch(t *testing.T) {
	err := runBuildScratch(&args.BuildScratchArgs{
		Binaries:    []string{"/tmp/binary"},
		Dockerfile:  "",
		BuilderName: "",
	})

	if err != nil {
		t.Fatalf("runBuildScratch:return_error: %v", err)
	}
}
