package cck

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cck-sdk-go/pkg/common"
	"io/ioutil"
	"net/http"
)

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
