package utils

import "log"

func LogIfError(prefix string, err error) {
	if err != nil {
		log.Println(prefix, err)
	}
}
