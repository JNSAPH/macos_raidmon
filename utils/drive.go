package utils

import (
	"strconv"
	"strings"

	"github.com/JNSAPH/macos_raidmon/structs"
)

func IsExternalDisk(disk structs.Disk) bool {
	// Always exclude internal disks
	if disk.OSInternal {
		return false
	}

	// Check for non-system mount points
	if disk.MountPoint != "" && isExternalMountPoint(disk.MountPoint) {
		return true
	}

	// Check APFS volumes for external mount points
	for _, vol := range disk.APFSVolumes {
		if vol.MountPoint != "" && isExternalMountPoint(vol.MountPoint) {
			return true
		}
	}

	// Fallback: Disk number >= 4 and not explicitly internal
	return parseDiskNumber(disk.DeviceIdentifier) >= 4
}

func isExternalMountPoint(mountPoint string) bool {
	// Exclude system-related mount points
	systemMounts := []string{
		"/System/Volumes/",
		"/Volumes/Recovery",
		"/Volumes/Update",
		"/Volumes/Preboot",
		"/Volumes/VM",
	}

	for _, prefix := range systemMounts {
		if strings.HasPrefix(mountPoint, prefix) {
			return false
		}
	}

	// Only consider /Volumes/ mount points as external
	return strings.HasPrefix(mountPoint, "/Volumes/")
}

func parseDiskNumber(deviceID string) int {
	parts := strings.Split(deviceID, "disk")
	if len(parts) < 2 {
		return -1
	}

	base := strings.Split(parts[1], "s")[0]
	num, _ := strconv.Atoi(base)
	return num
}
