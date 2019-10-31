package sil

import "fmt"

func setArchive(filename string) error {
	fmt.Println("would be setting archive bit for ", filename)
	return nil
}

func unsetArchive(filename string) error {
	fmt.Println("would be unsetting archive bit for ", filename)
	return nil
}
