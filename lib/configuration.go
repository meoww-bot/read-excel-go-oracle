package lib

type SitefolderEntry struct {
	Category     string `excel:"应用大类" json:"category"`       // 应用大类
	ClusterName  string `excel:"所属集群" json:"cluster_name"`   // 所属集群
	ClusterGroup string `excel:"集群属组" json:"cluster_group"`  // 集群属组
	Biz          string `excel:"承载业务" json:"biz"`            // 承载业务
	Purpose      string `excel:"用途" json:"purpose"`          // 用途
	Region       string `excel:"局址" json:"region"`           // 局址
	RackLocation string `excel:"机房位置" json:"rack_location"`  // 机房位置
	Brand        string `excel:"品牌" json:"brand"`            // 品牌
	Model        string `excel:"型号" json:"model"`            // 型号
	Ram          int    `excel:"内存" json:"ram"`              // 内存
	Disk         string `excel:"硬盘" json:"disk"`             // 硬盘
	CpuModel     string `excel:"CPU型号" json:"cpu_model"`     // CPU 型号
	RaidPlan     string `excel:"RAID规划" json:"raid_plan"`    // RAID 规划
	Bandwidth    string `excel:"带宽" json:"bandwidth"`        // 带宽
	NicInfo      string `excel:"网口信息" json:"nic_info"`       // 网口信息
	Hostname     string `excel:"主机名称" json:"hostname"`       // 主机名称
	Domain       string `excel:"域" json:"domain"`            // 域
	DeviceSN     string `excel:"设备序列号" json:"device_sn"`     // 设备序列号
	ServiceIp    string `excel:"业务地址" json:"service_ip"`     // 业务地址IP
	IsCtrl       string `excel:"是否接入4A" json:"is_ctrl"`      // 是否接入4A
	BmcMgmtIp    string `excel:"BMC管理地址" json:"bmc_mgmt_ip"` // BMC 管理地址
	BmcMgmtUser  string `excel:"BMC用户" json:"bmc_mgmt_user"` // BMC 用户 密码信息
	MonitorIp    string `excel:"监控IP" json:"monitor_ip"`     // 监控 IP
	Status       string `excel:"状态" json:"status"`           // 状态
	DevFrom      string `excel:"机器来源" json:"dev_from"`       // 机器来源
	CreateTime   string `excel:"接收时间" json:"create_time"`    // 接收时间
	QuitTime     string `excel:"退出时间" json:"quit_time"`      // 退出时间
	OsVersion    string `excel:"OS版本" json:"os_version"`     // OS 版本
	Remarks1     string `excel:"备注01" json:"remarks1"`       // 备注01
	Remarks2     string `excel:"备注02" json:"remarks2"`       // 备注02
	BrandModel   string `excel:"品牌及型号" json:"brand_model"`   // 品牌及型号
}

// https://books.studygolang.com/xorm/chapter-02/4.columns.html
type Inventory struct {
	AppType       string `json:"app_type" xorm:"APP_TYPE"`                  // 应用大类
	ClusterName   string `json:"cluster_name" xorm:"CLUSTER_NAME"`          // 所属集群
	ClusterGroup  string `json:"cluster_group" xorm:"CLUSTER_GROUP"`        // 集群属组
	Biz           string `json:"biz" xorm:"BIZ"`                            // 承载业务
	Purpose       string `json:"purpose" xorm:"PURPOSE"`                    // 用途 没法做 tag 了
	Region        string `json:"region" xorm:"REGION"`                      // 局址
	RackLocation  string `json:"rack_location" xorm:"RACK_LOCATION"`        // 机架位置
	Brand         string `json:"brand" xorm:"BRAND"`                        // 品牌
	Model         string `json:"model" xorm:"MODEL"`                        // 型号
	RamInfo       int    `json:"ram" xorm:"RAM_INFO"`                       // 内存
	DiskInfo      string `json:"disk" xorm:"DISK_INFO"`                     // 硬盘
	CpuInfo       string `json:"cpu_info" xorm:"CPU_INFO"`                  // CPU 型号
	RaidInfo      string `json:"raid_info" xorm:"RAID_INFO"`                // RAID 规划
	Bandwidth     string `json:"bandwidth" xorm:"BANDWIDTH"`                // 带宽 1000M 10GE
	NicInfo       string `json:"nic_info" xorm:"NIC_INFO"`                  // 网口信息
	NicNum        int    `json:"nic_num" xorm:"NIC_NUM"`                    // 网口个数
	NicBondMode   string `json:"nic_bond_mode" xorm:"NIC_BOND_MODE"`        // 网口bond类型
	Hostname      string `json:"hostname" xorm:"HOSTNAME"`                  // 主机名称
	ShortHostname string `json:"short_hostname" xorm:"SHORT_HOSTNAME"`      // 简短主机名称
	Domain        string `json:"domain" xorm:"DOMAIN"`                      // 域
	SerialNumber  string `json:"serial_number" xorm:"SERIAL_NUMBER"`        // 设备序列号
	ServiceIp     string `json:"service_ip" xorm:"SERVICE_IP"`              // 业务地址IP
	IsCtrl        int    `json:"is_ctrl" xorm:"IS_CTRL"`                    // 是否接入4A
	BmcMgmtIp     string `json:"bmc_mgmt_ip" xorm:"BMC_MGMT_IP"`            // BMC 管理地址
	BmcMgmtUser   string `json:"bmc_mgmt_user" xorm:"BMC_MGMT_USER"`        // BMC 用户 密码信息
	MonitorIp     string `json:"monitor_ip" xorm:"MONITOR_IP"`              // 监控 IP
	Status        int    `json:"status" xorm:"STATUS"`                      // 状态
	DeviceFrom    string `json:"device_from" xorm:"DEVICE_FROM"`            // 机器来源
	CreateTime    string `json:"create_time" xorm:"CREATE_TIME"`            // 接收时间
	QuitTime      string `json:"quit_time" xorm:"QUIT_TIME"`                // 退出时间
	OsVersion     string `json:"os_version" xorm:"'OS_VERSION'"`            // OS 版本
	Remarks1      string `json:"remarks1" xorm:"VARCHAR2(1000) 'REMARKS1'"` // 备注01
	Remarks2      string `json:"remarks2" xorm:"VARCHAR2(1000) 'REMARKS2'"` // 备注02
	UUID          string `json:"uuid" xorm:"VARCHAR2(255) 'UUID'"`          // UUID
}

func (Inventory) TableName() string {

	return "CMDB_INVENTORY"
}

const (
	StatusQuit             = -1 + iota // 退网
	StatusOK                           // 正常
	StatusUnderContruction             // 建设中
	StatusWaitQuiting                  // 待退网
	StatusMoved                        // 移交
	StatusUnknown                      // 未知
	StatusDiskIssue                    // 硬盘故障
	StatusNetworkIssue                 // 网络故障
	StatusMemoryIssue                  // 内存故障
	StatusCpuIssue                     // CPU故障
	StatusMotherBoardIssue             // 主板故障
)

var StatusMap = map[int]string{
	StatusQuit:             "退网",
	StatusOK:               "正常",
	StatusUnderContruction: "建设中",
	StatusWaitQuiting:      "待退网",
	StatusMoved:            "移交",
	StatusUnknown:          "未知",
	StatusDiskIssue:        "硬盘故障",
	StatusNetworkIssue:     "网络故障",
	StatusMemoryIssue:      "内存故障",
	StatusCpuIssue:         "CPU故障",
	StatusMotherBoardIssue: "主板故障",
}
