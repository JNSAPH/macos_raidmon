package structs

type Plist struct {
	AppleRAIDSets []AppleRAIDSet `plist:"AppleRAIDSets"`
}

type AppleRAIDSet struct {
	AppleRAIDSetUUID string        `plist:"AppleRAIDSetUUID"`
	BSDName          string        `plist:"BSD Name"`
	ChunkCount       int           `plist:"ChunkCount"`
	ChunkSize        int           `plist:"ChunkSize"`
	Content          string        `plist:"Content"`
	Level            string        `plist:"Level"`
	Members          []RAIDMember  `plist:"Members"`
	Name             string        `plist:"Name"`
	Rebuild          string        `plist:"Rebuild"`
	Size             int64         `plist:"Size"`
	Spares           []interface{} `plist:"Spares"`
	Status           string        `plist:"Status"`
}

type RAIDMember struct {
	AppleRAIDMemberUUID string `plist:"AppleRAIDMemberUUID"`
	BSDName             string `plist:"BSD Name"`
	MemberStatus        string `plist:"MemberStatus"`
}
