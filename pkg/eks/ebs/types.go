package ebs

type Response struct {
	Code      interface{} `json:"Code"`
	Message   interface{} `json:"Msg"`
	Data      interface{} `json:"Data,omitempty"`
	RequestId interface{} `json:"RequestId"`
	OpenapiPage
}

type OpenapiPage struct {
	TotalCount int `json:"TotalCount,omitempty"`
	PageIndex  int `json:"PageIndex,omitempty"`
	PageSize   int `json:"PageSize,omitempty"`
}

type AttachDiskArgs struct {
	VolumeID string `json:"block_id"`
	NodeID   string `json:"node_id"`
}
type AttachDiskResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type DetachDiskArgs struct {
	VolumeID string `json:"block_id"`
}
type DetachDiskResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type DeleteDiskArgs struct {
	VolumeID string `json:"block_id"`
}
type DeleteDiskResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type FindDiskByVolumeIDArgs struct {
	VolumeID string `json:"block_id"`
}

type DiskInfo struct {
	NodeID   string `json:"node_id"`
	Status   string `json:"status"`
	Uuid     string `json:"disk_uuid"`
	IsFormat int    `json:"is_format"`
}
type FindDiskByVolumeIDResponse struct {
	Response
	Data struct {
		DiskSlice []DiskInfo `json:"block_info"`
	} `json:"data"`
}

type FindDeviceNameByVolumeIDArgs struct {
	VolumeID string `json:"disk_id"`
}
type FindDeviceNameByVolumeIDResponse struct {
	Response
	Data struct {
		DeviceName string `json:"device_name"`
	} `json:"Data"`
}

type CreateEbsReq struct {
	ProductSource     string `json:"ProductSource"`
	ProductServerId   string `json:"ProductServerId"`
	AvailableZoneCode string `json:"AvailableZoneCode"`
	DiskName          string `json:"DiskName"`
	DiskFeature       string `json:"DiskFeature"`
	Size              int    `json:"Size"`
	Number            int    `json:"Number"`
	BillingMethod     string `json:"BillingMethod"`
}

type CreateEbsResp struct {
	Response
	Data CreateEbsRespData
}

type CreateEbsRespData struct {
	EventId   string   `json:"EventId"`
	DiskIdSet []string `json:"DiskIdSet"`
}

type DescribeTaskStatusResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	} `json:"Data"`
}

type UpdateBlockFormatFlagArgs struct {
	BlockID  string `json:"block_id"`
	IsFormat int    `json:"is_format"`
}

type UpdateBlockFormatFlagResponse struct {
	Response
}
