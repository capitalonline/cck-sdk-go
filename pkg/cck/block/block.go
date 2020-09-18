package block

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cck-sdk-go/pkg/common"
	"io/ioutil"
	"net/http"
)

func CreateBlock(args *CreateBlockArgs) (*CreateBlockResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionCreateBlock, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &CreateBlockResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func AttachBlock(args *AttachBlockArgs) (*AttachBlockResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionAttachBlock, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &AttachBlockResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func DetachBlock(args *DetachBlockArgs) (*DetachBlockResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionDetachBlock, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DetachBlockResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func DeleteBlock(args *DeleteBlockArgs) (*DeleteBlockResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionDeleteBlock, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DeleteBlockResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func FindBlockByVolumeID(args *FindBlockByVolumeIDArgs) (*FindBlockByVolumeIDResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionFindBlockByVolumeID, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &FindBlockByVolumeIDResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}

func DescribeTaskStatus(TaskID string) (*DescribeTaskStatusResponse, error) {
	payload := struct {
		TaskID string `json:"task_id"`
	}{
		TaskID,
	}

	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionBlockTaskStatus, http.MethodPost, nil, body)
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

	req, err := common.NewCCKRequest(common.ActionUpdateBmBlock, http.MethodPost, nil, body)

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
