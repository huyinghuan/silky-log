package schema

//Log 虚拟机配置
type Log struct {
	ID           int64  `xorm:"unique" json:"id"`
	IP           string `xorm:"ip" json:"ip"`
	SilkyVersion string `xorm:"silky_version" json:"silky_version"`
	OSPlatform   string `xorm:"os_platform" json:"os_platform"`
	OSArch       string `xorm:"os_arch" json:"os_arch"`
	OSRelease    string `xorm:"os_release" json:"os_release"`
	NodeVersion  string `xorm:"node_version" json:"node_version"`
	UpdatedAt    int64  `xorm:"updated"`
	Username     string `xorm:"username" json:"username"`
	ProjectName  string `xorm:"project_name" json:"project_name"`
	Type         string `xorm:"type" json:"type"`
	Status       string `xorm:"status" json:"status"`
}
