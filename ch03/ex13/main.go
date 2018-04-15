// Write const de clarations for KB, MB, up through YB as compactly as you can.

package main

const (
	_ = 1 << (10 * iota)
	// KiB 1024
	KiB
	// MiB 1048576
	MiB
	// GiB 1073741824
	GiB
	// TiB 1099511627776 (exceeds 1 << 32)
	TiB
	// PiB 1125899906842624
	PiB
	// EiB 1152921504606846976
	EiB
	// ZiB 1180591620717411303424 (exceeds 1 << 64)
	ZiB
	// YiB 1208925819614629174706176
	YiB
)

const (
	ex = 1000
	// KB 1,000
	KB = ex
	// MB 1,000,000
	MB = KB * ex
	// GB 1,000,000,000
	GB = MB * ex
	// TB 1,000,000,000,000
	TB = GB * ex
	// PB 1,000,000,000,000,000
	PB = TB * ex
	// EB 1,000,000,000,000,000,000
	EB = PB * ex
	// ZB 1,000,000,000,000,000,000,000
	ZB = EB * ex
	// YB 1,000,000,000,000,000,000,000,000
	YB = ZB * ex
)

func main() {
}
