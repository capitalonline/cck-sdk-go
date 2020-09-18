package block

type Response struct {
	Code     string `json:"Code"`
	Message  string `json:"Message"`
	CodeDesc string `json:"codeDesc,omitempty"`
}

type AttachBlockArgs struct {
	VolumeID string `json:"bm_block_id"`
	NodeID   string `json:"node_id"`
}
type AttachBlockResponse struct {
	Response
	TaskID string `json:"TaskId"`
	Data   struct {
		Npn  string `json:"npn"`
		Vip  string `json:"vip"`
		Port string `json:"port"`
	} `json:"data"`
}

type DetachBlockArgs struct {
	VolumeID string `json:"bm_block_id"`
}
type DetachBlockResponse struct {
	Response
	TaskID string `json:"TaskId"`
	Data   struct {
		Npn string `json:"npn"` // 裸金属磁盘npn设备信息
	} `json:"data"`
}

type DeleteBlockArgs struct {
	VolumeID string `json:"bm_block_id"`
}
type DeleteBlockResponse struct {
	Response
	TaskID string `json:"TaskId"`
}

type FindBlockByVolumeIDArgs struct {
	VolumeID string `json:"bm_block_id"`
}

type BlockInfo struct {
	NodeID   string `json:"node_id"`
	Status   string `json:"status"`
	Uuid     string `json:"bm_disk_id"`
	IsFormat int    `json:"is_format"`
}
type FindBlockByVolumeIDResponse struct {
	Response
	Data struct {
		BlockSlice []BlockInfo `json:"bm_block_info"`
	} `json:"data"`
}

type CreateBlockArgs struct {
	Name      string `json:"name"`
	RegionID  string `json:"site_id"`
	DiskType  string `json:"bm_disk_type"`
	Size      int    `json:"size"`
	Iops      int    `json:"iops"`
	Bandwidth int    `json:"bandwidth"`
}
type CreateBlockResponse struct {
	Response
	Data struct {
		VolumeID string `json:"bm_block_id"`
	} `json:"Data"`
	TaskID string `json:"TaskId"`
}

type DescribeTaskStatusResponse struct {
	Response
	Data struct {
		Status string `json:"status"`
	} `json:"Data"`
}

type UpdateBlockFormatFlagArgs struct {
	BlockID  string `json:"bm_block_id"`
	IsFormat int    `json:"is_format"`
	Size     int    `json:"size"`
}

type UpdateBlockFormatFlagResponse struct {
	Response
	TaskID string `json:"TaskId"`
}
