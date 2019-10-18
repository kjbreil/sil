package sil

import "syscall"

func unsetArchiveBit(p string) error {

	nameptr, err := windows.UTF16PtrFromString(p)
	if err != nil {
		return err
	}

	// Do whatever windows calls are needed to change
	// the file into a hidden file; something like
	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_NORMAL)
	if err != nil {
		return err
	}

	return nil
}
