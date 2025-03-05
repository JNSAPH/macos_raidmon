package utils

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/JNSAPH/macos_raidmon/structs"
	"howett.net/plist"
)

func GetRaidDetails() structs.RaidDetails {
	// Run the diskutil appleRAID list command
	cmd := exec.Command("diskutil", "appleRAID", "list", "-plist")

	// Get the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command: %v\n", err)
	}

	var data structs.RaidDetails
	f := bytes.NewReader(output)
	decoder := plist.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	// Return the plist
	return data
}

func GetDriveDetails() structs.DriveDetails {

	return structs.DriveDetails{}

}
