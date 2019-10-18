package sil

import "log"

func unsetArchiveBit(p string) error {
	log.Println("Not running on windows system, cannot unset archive bit")
	return nil
}
