//go:build windows
// +build windows

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

	atime = time.Unix(0, fileInfo.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
	mtime = time.Unix(0, fileInfo.Sys().(*syscall.Win32FileAttributeData).LastWriteTime.Nanoseconds())
	ctime = time.Unix(0, fileInfo.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
	return
}
