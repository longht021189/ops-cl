package action

import (
	"regexp"

	"github.com/longht021189/ops-cl/utils/file"
)

func GetBinaryDependencies(path string) ([]string, error) {
	var result []string

	lines := []string{
		"linux-vdso.so.1 (0x00007ffcfb0c5000)",
		"libFoundation.so => /usr/lib/swift/linux/libFoundation.so (0x00007f4c5a421000)",
		"libswiftGlibc.so => /usr/lib/swift/linux/libswiftGlibc.so (0x00007f4c5b111000)",
		"libm.so.6 => /lib/x86_64-linux-gnu/libm.so.6 (0x00007f4c5a083000)",
		"/lib64/ld-linux-x86-64.so.2 (0x00007f4c5aefe000)",
		"libswift_Concurrency.so => /usr/lib/swift/linux/libswift_Concurrency.so (0x00007f4c5b081000)",
	}

	re := regexp.MustCompile(`(.* ?=>)?(.*)\(.*\)`)
	for _, line := range lines {
		values := re.FindStringSubmatch(line)
		length := len(values)

		if length > 0 && file.Exists(values[length-1]) {
			result = append(result, values[length-1])
		}
	}

	return result, nil
}

// objdump -p $(which protoc-gen-grpc-swift) | grep NEEDED
//   NEEDED               libFoundation.so
//   NEEDED               libswiftGlibc.so
//   NEEDED               libm.so.6
//   NEEDED               libpthread.so.0
//   NEEDED               libutil.so.1
//   NEEDED               libdl.so.2
//   NEEDED               libswiftDispatch.so
//   NEEDED               libdispatch.so
//   NEEDED               libBlocksRuntime.so
//   NEEDED               libswift_Concurrency.so
//   NEEDED               libswiftCore.so
//   NEEDED               libc.so.6

//==================================================================

// ldd $(which protoc-gen-grpc-swift)
//
//
//
//
// 	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f4c59e64000)
// 	libutil.so.1 => /lib/x86_64-linux-gnu/libutil.so.1 (0x00007f4c59c61000)
// 	libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f4c59a5d000)
// 	libswiftDispatch.so => /usr/lib/swift/linux/libswiftDispatch.so (0x00007f4c5b0d8000)
// 	libdispatch.so => /usr/lib/swift/linux/libdispatch.so (0x00007f4c597fe000)
// 	libBlocksRuntime.so => /usr/lib/swift/linux/libBlocksRuntime.so (0x00007f4c595fb000)
//
// 	libswiftCore.so => /usr/lib/swift/linux/libswiftCore.so (0x00007f4c58fe4000)
// 	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f4c58bf3000)
// 	libicuucswift.so.65 => /usr/lib/swift/linux/libicuucswift.so.65 (0x00007f4c587f2000)
// 	libicui18nswift.so.65 => /usr/lib/swift/linux/libicui18nswift.so.65 (0x00007f4c582c7000)
// 	libgcc_s.so.1 => /lib/x86_64-linux-gnu/libgcc_s.so.1 (0x00007f4c580af000)
// 	libstdc++.so.6 => /usr/lib/x86_64-linux-gnu/libstdc++.so.6 (0x00007f4c57d26000)
//
// 	librt.so.1 => /lib/x86_64-linux-gnu/librt.so.1 (0x00007f4c57b1e000)
// 	libicudataswift.so.65 => /usr/lib/swift/linux/libicudataswift.so.65 (0x00007f4c55e6f000)
