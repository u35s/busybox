package libbb

import (
	"strconv"
	"strings"
)

func Atoi_sfx(a string, sfx []suffix_mult) int64 {
	for i := range sfx {
		if strings.HasSuffix(a, sfx[i].suffix) {
			num, err := strconv.Atoi(strings.TrimSuffix(a, sfx[i].suffix))
			if err == nil {
				return int64(num) * int64(sfx[i].mult)
			}
			return 0
		}
	}
	return 0
}

type suffix_mult struct {
	suffix string
	mult   int
}

var Suffixes_bkm = []suffix_mult{
	{"b", 512},
	{"k", 1024},
	{"m", 1024 * 1024},
	{"", 0},
}

var Suffixes_cwbkMG = []suffix_mult{
	{"c", 1},
	{"w", 2},
	{"b", 512},
	{"kB", 1000},
	{"kD", 1000},
	{"k", 1024},
	{"KB", 1000}, /* compat with coreutils dd */
	{"KD", 1000}, /* compat with coreutils dd */
	{"K", 1024},  /* compat with coreutils dd */
	{"MB", 1000000},
	{"MD", 1000000},
	{"M", 1024 * 1024},
	{"GB", 1000000000},
	{"GD", 1000000000},
	{"G", 1024 * 1024 * 1024},
	/* "D" suffix for decimal is not in coreutils manpage, looks like it's deprecated */
	/* coreutils also understands TPEZY suffixes for tera- and so on, with B suffix for decimal */
	{"", 0},
}

var Suffixes_kmg_i = []suffix_mult{
	{"KiB", 1024},
	{"kiB", 1024},
	{"K", 1024},
	{"k", 1024},
	{"MiB", 1048576},
	{"miB", 1048576},
	{"M", 1048576},
	{"m", 1048576},
	{"GiB", 1073741824},
	{"giB", 1073741824},
	{"G", 1073741824},
	{"g", 1073741824},
	{"KB", 1000},
	{"MB", 1000000},
	{"GB", 1000000000},
	{"", 0},
}
