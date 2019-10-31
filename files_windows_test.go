package sil

import (
	"syscall"
	"testing"
)

// Test that the archive bit can be unset on a file and then reset
func Test_archiveBit(t *testing.T) {
	filename := ".gitignore"
	err := unsetArchive(filename)
	if err != nil {
		t.Fatal(err)

	}
	pointer, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		t.Fatal(err)

	}
	att, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		t.Fatal(err)

	}
	if att&syscall.FILE_ATTRIBUTE_ARCHIVE != 0 {
		t.Fatal("archive bit still set")
	}

	err = setArchive(filename)
	if err != nil {
		t.Fatal(err)

	}
	att, err = syscall.GetFileAttributes(pointer)
	if err != nil {
		t.Fatal(err)

	}
	if att&syscall.FILE_ATTRIBUTE_ARCHIVE == 0 {
		t.Fatal("archive bit not set, manually fix your .gitignore")
	}

}
