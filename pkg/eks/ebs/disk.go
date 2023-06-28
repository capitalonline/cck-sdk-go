package ebs

import (
	"cck-sdk/pkg/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateEbs(args *CreateEbsReq) (*CreateEbsResp, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionCreateEbs, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &CreateEbsResp{}
	err = json.Unmarshal(content, res)

	return res, err
}

func AttachDisk(args *AttachDiskArgs) (*AttachDiskResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionAttachDisk, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &AttachDiskResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func DetachDisk(args *DetachDiskArgs) (*DetachDiskResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionDetachDisk, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DetachDiskResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func DeleteDisk(args *DeleteDiskArgs) (*DeleteDiskResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionDeleteDisk, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DeleteDiskResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func FindDiskByVolumeID(args *FindDiskByVolumeIDArgs) (*FindDiskByVolumeIDResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionFindDiskByVolumeID, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &FindDiskByVolumeIDResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func FindDeviceNameByVolumeID(args *FindDeviceNameByVolumeIDArgs) (*FindDeviceNameByVolumeIDResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionFindDiskByVolumeID, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &FindDeviceNameByVolumeIDResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func DescribeTaskStatus(TaskID string) (*DescribeTaskStatusResponse, error) {
	payload := struct {
		EventId string `json:"EventId"`
	}{
		TaskID,
	}

	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionDiskTaskStatus, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeTaskStatusResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func UpdateBlockFormatFlag(args *UpdateBlockFormatFlagArgs) (*UpdateBlockFormatFlagResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionUpdateBlock, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &UpdateBlockFormatFlagResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}
