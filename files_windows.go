package sil

import (
	"syscall"
)

func setArchive(filename string) error {
	pointer, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}
	return syscall.SetFileAttributes(pointer, syscall.FILE_ATTRIBUTE_ARCHIVE)
}

func unsetArchive(filename string) error {
	pointer, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}
	return syscall.SetFileAttributes(pointer, syscall.FILE_ATTRIBUTE_NORMAL)
}
