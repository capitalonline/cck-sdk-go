package ebs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/capitalonline/cck-sdk-go/pkg/common"
	"io"
	"net/http"
)

func CreateEbs(args *CreateEbsReq) (*CreateEbsResp, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionCreateEbs, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &CreateEbsResp{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}

	return res, err
}

func AttachDisk(args *AttachDiskArgs) (*AttachDiskResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionAttachEbs, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &AttachDiskResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
	return res, err
}

func DetachDisk(args *DetachDiskArgs) (*DetachDiskResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionDetachEbs, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
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

	req, err := common.NewEbsRequest(common.ActionDeleteEbs, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DeleteDiskResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
	return res, err
}

func FindDiskByVolumeID(args *FindDiskByVolumeIDArgs) (*FindDiskByVolumeIDResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionDescribeEbs, http.MethodGet, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &FindDiskByVolumeIDResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
	return res, err
}

func FindDeviceNameByVolumeID(args *FindDeviceNameByDiskIdArgs) (*FindDeviceNameByDiskIdResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionFindDiskByVolumeID, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &FindDeviceNameByDiskIdResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
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

	req, err := common.NewEbsRequest(common.ActionDescribeEvent, http.MethodGet, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeTaskStatusResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
	return res, err
}

func UpdateBlockFormatFlag(args *UpdateBlockFormatFlagArgs) (*UpdateBlockFormatFlagResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionUpdateBlock, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &UpdateBlockFormatFlagResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
	return res, err
}

func ExtendDisk(args *ExtendDiskArgs) (*ExtendDiskResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewEbsRequest(common.ActionExtendEbs, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &ExtendDiskResponse{}
	err = json.Unmarshal(content, res)
	if res.Code != common.EbsSuccessCode {
		return nil, errors.New(fmt.Sprintf("%s request failed,msg:%s", common.ActionCreateEbs, res.Message))
	}
	return res, err
}
