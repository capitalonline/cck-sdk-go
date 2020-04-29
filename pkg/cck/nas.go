package cck

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cck-sdk-go/pkg/common"
	"io/ioutil"
	"net/http"
)

func DescribeNasInstances(nasID, nasName, nasSiteID, clusterID string, usageFlag bool, pageNumber, pageSize int) (*DescribeNasInstancesResponse, error) {
	payload := struct {
		NasID		string
		NasName		string
		NasSiteID	string
		ClusterID 	string
		UsageFlag	bool
		PageNumber	int
		PageSize	int
	}{
		nasID,
		nasName,
		nasSiteID,
		clusterID,
		usageFlag,
		pageNumber,
		pageSize,
	}
	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDescribeNasInstances, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DescribeNasInstancesResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func CreateNas(nasSiteID, nasName, diskType, diskSize string) (*CreateNasResponse, error) {
	payload := struct {
		NasSiteID	string
		NasName		string
		DiskType 	string
		DiskSize 	string
	}{
		nasSiteID,
		nasName,
		diskType,
		diskSize,
	}
	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionCreateNas, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &CreateNasResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func ResizeNas(nasSiteID, nasID, diskSize string) (*ResizeNasResponse, error) {
	payload := struct {
		NasSiteID	string
		NasID 		string
		DiskSize	string
	}{
		nasSiteID,
		nasID,
		diskSize,
	}
	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionResizeNas, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &ResizeNasResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func DeleteNas(nasID string) (*DeleteNasResponse, error) {
	payload := struct {
		NasID 		string
	}{
		nasID,
	}
	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionDeleteNas, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &DeleteNasResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func MountNas(nasID, clusterID string) (*MountNasResponse, error) {
	payload := struct {
		NasID     string
		ClusterID string
	}{
		nasID,
		clusterID,
	}
	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionMountNas, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &MountNasResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}

func UnMountNas(nasID string) (*UnMountNasResponse, error) {
	payload := struct {
		NasID     string
	}{
		nasID,
	}
	body, err := common.MarshalJsonToIOReader(payload)
	if err != nil {
		return nil, err
	}
	req, err := common.NewCCKRequest(common.ActionUnMountNas, http.MethodPost, nil, body)
	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &UnMountNasResponse{}
	err = json.Unmarshal(content, res)
	return res, err
}




