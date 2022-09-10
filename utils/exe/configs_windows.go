//go:build windows

package exe

var (
	prefixCmd = []string{"powershell", "-nologo", "-noprofile"}
)
