package cck

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

type MountNasResponse struct {
	BaseResponse
	Data struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}
