// Code generated by goctl. DO NOT EDIT.
package types

type ShrcbMonitorRes struct {
	Datetime        string `json:"datetime"`
	HostName        string `json:"hostName"`
	IpAddress       string `json:"ipAddress"`
	Msg             string `json:"msg"`
	Title           string `json:"title"`
	Severity        string `json:"severity"`
	SysNameEn       string `json:"sysNameEn"`
	SysNameCn       string `json:"sysNameCn"`
	BlindTimeMinute int    `json:"BlindTimeMinute"`
}

type ShrcbMonitorResp struct {
	Msg string `json:"msg"`
}

type ShrcbMonitorRespList struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []*ShrcbMonitorResp `json:"data"`
}

type RefreshResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type SystemList struct {
	SystemList []*System `json:"SystemList"`
}

type System struct {
	SystemNameEn        string                    `json:"SystemNameEn"`
	SystemNameCn        string                    `json:"SystemNameCn"`
	InnerIP             string                    `json:"InnerIP"`
	OuterIP             string                    `json:"OuterIP"`
	HostName            string                    `json:"HostName"`
	User                string                    `json:"User"`
	Passwd              string                    `json:"Passwd"`
	ClusterName         string                    `json:"ClusterName"`
	NeedReport          map[string]map[string]int `json:"NeedReport"`
	Severity            string                    `json:"Severity"`
	BlindTimeMinute     int                       `json:"BlindTimeMinute"`
	TellTheTales        bool                      `json:"TellTheTales"`
	BlindInsertDatabase bool                      `json:"BlindInsertDatabase"`
	ConfirmReportNumber int                       `json:"ConfirmReportNumber"`
}

type ServicesList struct {
	Items []*ApiService `json:"items"`
}

type ApiService struct {
	Name                        string            `json:"name"`
	ServiceType                 string            `json:"type"`
	ClusterRef                  *ApiClusterRef    `json:"clusterRef"`
	ServiceState                string            `json:"serviceState"`
	HealthSummary               string            `json:"healthSummary"`
	ConfigStale                 bool              `json:"configStale"`
	ConfigStalenessStatus       string            `json:"configStalenessStatus"`
	ClientConfigStalenessStatus string            `json:"clientConfigStalenessStatus"`
	HealthChecks                []*ApiHealthCheck `json:"healthChecks"`
	ServiceUrl                  string            `json:"serviceUrl"`
	RoleInstancesUrl            string            `json:"roleInstancesUrl"`
	MaintenanceMode             bool              `json:"maintenanceMode"`
	MaintenanceOwners           []string          `json:"maintenanceOwners"`
	DisplayName                 string            `json:"displayName"`
	ServiceVersion              string            `json:"serviceVersion"`
}

type ApiClusterRef struct {
	ClusterName string `json:"clusterName"`
	DisplayName string `json:"displayName"`
}

type ApiHealthCheck struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Explanation string `json:"explanation"`
	Suppressed  bool   `json:"suppressed"`
}

type EncryptRes struct {
	OriString []string `json:"oriString"`
	Encrypted []string `json:"encrypted"`
}

type EncryptRp struct {
	Encrypted map[string]string `json:"encrypted"`
	OriString map[string]string `json:"oriString"`
}

type EncryptResp struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data *EncryptRp `json:"data"`
}

type EmailInfo struct {
	Host      string           `json:"host"`
	Port      string           `json:"port"`
	Send2Who  []string         `json:"send2who"`
	EmailUser []*EmailUserInfo `json:"emailUser"`
}

type EmailUserInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

type PostEmailRes struct {
	Address []string `json:"address"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type PostEmailRp struct {
	Success bool `json:"success"`
}

type PostEmailResp struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data *PostEmailRp `json:"data"`
}

type EasyEmailRes struct {
	Send2Who  []string         `json:"send2who"`
	EmailUser []*EmailUserInfo `json:"emailUser"`
	Subject   string           `json:"subject"`
	Body      string           `json:"body"`
}

type EasyEmailRp struct {
	Success bool `json:"success"`
}

type EasyEmailResp struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data *PostEmailRp `json:"data"`
}

type ReportResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
