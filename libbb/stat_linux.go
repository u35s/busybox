package libbb

import "syscall"

type Stat_t struct {
	*syscall.Stat_t
}

func (s *Stat_t) ChangeTime() syscall.Timespec {
	return s.Ctim
}
func (s *Stat_t) AccessTime() syscall.Timespec {
	return s.Atim
}
func (s *Stat_t) ModifyTime() syscall.Timespec {
	return s.Mtim
}
