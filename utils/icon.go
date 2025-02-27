package utils

import (
	"os"
)

func GetDriveIco() []byte {
	arr, err := os.ReadFile("assets/drive.ico")
	if err != nil {
		panic(err)
	}
	return arr
}

func GetDegradedDriveIco() []byte {
	arr, err := os.ReadFile("assets/degraded_drive.ico")
	if err != nil {
		panic(err)
	}
	return arr
}
