package libbb

import (
	"os"
	"syscall"
)

func Open(name string, flag int, perm os.FileMode) (int, error) {
	return syscall.Open(name, flag, uint32(perm.Perm()))
}
