package alarm

type SendAlarmArgs struct {
	ClusterId  string `json:"cluster_id"`
	Site       string `json:"site"`
	Msg        string `json:"msg"`
	Hostname   string `json:"hostname,omitempty"`
	NtfName    string `json:"ntf_name,omitempty"`
	AlarmType  string `json:"alarm_type,omitempty"`
	AlarmGroup string `json:"alarm_group,omitempty"`
	Ip         string `json:"ip"`
	Tag1       string `json:"tag1"`
}

type SendAlarmResponse struct {
	Code     string `json:"Code"`
	Message  string `json:"Message"`
	CodeDesc string `json:"codeDesc,omitempty"`
}
