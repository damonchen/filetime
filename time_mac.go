//go:build darwin
// +build darwin

package utils

import (
	"os"
	"syscall"
	"time"
)

func StatTimes(name string) (ctime time.Time, mtime time.Time, atime time.Time, err error) {

	fileInfo, err := os.Stat(name)
	if err != nil {
		return
	}

	stat := fileInfo.Sys().(*syscall.Stat_t)
	atime = time.Unix(stat.Atimespec.Unix())
	mtime = time.Unix(stat.Mtimespec.Unix())
	ctime = time.Unix(stat.Ctimespec.Unix())
	return
}
