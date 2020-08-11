package disk

type Response struct {
	Code     string `json:"Code"`
	Message  string `json:"Message"`
	CodeDesc string `json:"codeDesc,omitempty"`
}

type AttachDiskArgs struct {
	VolumeID string `json:"volume_id"`
	NodeID   string `json:"node_id"`
}
type AttachDiskResponse struct {
	Response
	TaskID string `json:"taskID"`
}

type DetachDiskArgs struct {
	VolumeID string `json:"volume_id"`
	NodeID   string `json:"node_id"`
}
type DetachDiskResponse struct {
	Response
	TaskID string `json:"taskID"`
}

type DeleteDiskArgs struct {
	VolumeID string `json:"disk_id"`
}
type DeleteDiskResponse struct {
	Response
	TaskID string `json:"taskID"`
}

type FindDiskByVolumeIDArgs struct {
	VolumeID string `json:"volume_id"`
}
type FindDiskByVolumeIDResponse struct {
	Response
	Data struct {
		NodeID string `json:"instance_id"`
		Status     string `json:"status"`
	} `json:"Data"`
}

type FindDeviceNameByVolumeIDArgs struct {
	VolumeID string `json:"volume_id"`
}
type FindDeviceNameByVolumeIDResponse struct {
	Response
	Data struct {
		DeviceName string `json:"device_name"`
	} `json:"Data"`
}

type CreateDiskArgs struct {
	Name        string `json:"name"`
	ClusterID   string `json:"cluster_id"`
	RegionID    string `json:"region_id"`
	Fstype      string `json:"fstype"`
	StorageType string `json:"storage_type"`
	RequestGB   int    `json:"requestGB"`
	ZoneID      string `json:"zone_id"`
}
type CreateDiskResponse struct {
	Response
	Data struct {
		VolumeID string `json:"volume_id"`
	} `json:"Data"`
	TaskID string `json:"taskID"`
}

type DescribeTaskStatusResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	} `json:"Data"`
}
