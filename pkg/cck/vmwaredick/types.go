package vmwaredisk

type Response struct {
	Code     string `json:"Code"`
	Message  string `json:"Message"`
	CodeDesc string `json:"codeDesc,omitempty"`
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
	NodeID string `json:"node_id"`
	Status string `json:"status"`
	Uuid   string `json:"disk_uuid"`
}

type FindDiskByVolumeIDResponse struct {
	Response
	Data struct {
		DiskSlice []DiskInfo `json:"block_info"`
	} `json:"data"`
}

type CreateDiskArgs struct {
	RegionID    string `json:"site_id"`
	DiskType    string `json:"disk_type"`
	Size        int    `json:"size"`
	Iops        int    `json:"iops"`
	ClusterName string `json:"cluster_name"`
}
type CreateDiskResponse struct {
	Response
	Data struct {
		VolumeID string `json:"block_id"`
	} `json:"Data"`
	TaskID string `json:"TaskId"`
}

type DescribeTaskStatusResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	} `json:"Data"`
}
