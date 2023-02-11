//go:build linux
// +build linux

package utils

import (
	"os"
	"syscall"
	"time"
)

func StatTime(name string) (ctime time.Time, mtime time.Time, atime time.Time, err error) {
	fileInfo, err := os.Stat(name)
	if err != nil {
		return
	}

	stat := fileInfo.Sys().(*syscall.Stat_t)
	atime = time.Unix(stat.Atim.Unix())
	mtime = time.Unix(stat.Mtim.Unix())
	ctime = time.Unix(stat.Ctim.Unix())
	return
}
