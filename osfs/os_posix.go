/*
Sniperkit-Bot
- Date: 2018-08-12 12:10:57.421562389 +0200 CEST m=+0.041271547
- Status: analyzed
*/

// +build !windows

package osfs

import (
	"syscall"
)

func (f *file) Lock() error {
	f.m.Lock()
	defer f.m.Unlock()

	return syscall.Flock(int(f.File.Fd()), syscall.LOCK_EX)
}

func (f *file) Unlock() error {
	f.m.Lock()
	defer f.m.Unlock()

	return syscall.Flock(int(f.File.Fd()), syscall.LOCK_UN)
}
