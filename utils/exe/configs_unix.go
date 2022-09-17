//go:build !windows

package exe

var (
	prefixCmd = []string{}
)

func RunMeElevated() error {
	return nil
}

func IsAdmin() bool {
	return false
}
