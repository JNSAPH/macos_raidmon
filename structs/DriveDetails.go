package structs

type DriveDetails struct {
	AllDisks              []string `plist:"AllDisks"`
	AllDisksAndPartitions []Disk   `plist:"AllDisksAndPartitions"`
	VolumesFromDisks      []string `plist:"VolumesFromDisks"`
	WholeDisks            []string `plist:"WholeDisks"`
}

type Disk struct {
	Content            string       `plist:"Content,omitempty"`
	DeviceIdentifier   string       `plist:"DeviceIdentifier"`
	OSInternal         bool         `plist:"OSInternal,omitempty"`
	Partitions         []Partition  `plist:"Partitions,omitempty"`
	Size               uint64       `plist:"Size"`
	APFSPhysicalStores []APFSStore  `plist:"APFSPhysicalStores,omitempty"`
	APFSVolumes        []APFSVolume `plist:"APFSVolumes,omitempty"`
	MountPoint         string       `plist:"MountPoint,omitempty"`
	VolumeName         string       `plist:"VolumeName,omitempty"`
	VolumeUUID         string       `plist:"VolumeUUID,omitempty"`
}

type Partition struct {
	Content          string `plist:"Content"`
	DeviceIdentifier string `plist:"DeviceIdentifier"`
	DiskUUID         string `plist:"DiskUUID,omitempty"`
	Size             uint64 `plist:"Size"`
	VolumeName       string `plist:"VolumeName,omitempty"`
	VolumeUUID       string `plist:"VolumeUUID,omitempty"`
}

type APFSStore struct {
	DeviceIdentifier string `plist:"DeviceIdentifier"`
}

type APFSVolume struct {
	CapacityInUse    uint64     `plist:"CapacityInUse"`
	DeviceIdentifier string     `plist:"DeviceIdentifier"`
	DiskUUID         string     `plist:"DiskUUID,omitempty"`
	MountPoint       string     `plist:"MountPoint,omitempty"`
	OSInternal       bool       `plist:"OSInternal"`
	Size             uint64     `plist:"Size"`
	VolumeName       string     `plist:"VolumeName"`
	VolumeUUID       string     `plist:"VolumeUUID"`
	MountedSnapshots []Snapshot `plist:"MountedSnapshots,omitempty"`
}
type Snapshot struct {
	Sealed             string `plist:"Sealed"`
	SnapshotBSD        string `plist:"SnapshotBSD"`
	SnapshotMountPoint string `plist:"SnapshotMountPoint"`
	SnapshotName       string `plist:"SnapshotName"`
	SnapshotUUID       string `plist:"SnapshotUUID"`
}
