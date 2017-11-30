package myio

type KeyItem struct {
	Keyid     int
	Keyname   string
	KeyValue  string
	Keytype   string
	UsedValue string
}

type FeatureItem struct {
	FeatureId   int
	FeatureName string
	KeyItems    []*KeyItem
	DeadlineDay string
	PeriodDay   string //宽限期
}

type Licensefile struct {
	LicenSerialNo string
	Content       string
	Revoke        string
	ProductName   string
	VersionName   string
	Feature       []*FeatureItem
}

type WorkKeyItem struct {
	LicenSerialNo string
	Revoke        string
	ProductName   string
	VersionName   string
	Keyid         int
	Keyname       string
	KeyValue      string
	Keytype       string
	UsedValue     string
	DeadlineDay   string
	PeriodDay     string //宽限期
	FeatureName   string
}

type LicenseMgr struct {
	WorkKeyItem map[string]*WorkKeyItem //key工作Licens的key信息，吊销后，就不在这个map里了
	LicenseSet  map[string]*Licensefile //lsn所有license的文件信息
}

var LicenMgr *LicenseMgr

//把值cp出去
func GetActiveLicenseFileInfo() []Licensefile {

	return nil
}
