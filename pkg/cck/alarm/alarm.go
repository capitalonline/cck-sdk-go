package alarm

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cck-sdk-go/pkg/common"
	"io/ioutil"
	"net/http"
)

func SendAlarm(args *SendAlarmArgs) (*SendAlarmResponse, error) {
	body, err := common.MarshalJsonToIOReader(args)
	if err != nil {
		return nil, err
	}

	req, err := common.NewCCKRequest(common.ActionSendAlarm, http.MethodPost, nil, body)

	response, err := common.DoRequest(req)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("http error:%s, %s", response.Status, string(content))
	}

	res := &SendAlarmResponse{}
	err = json.Unmarshal(content, res)

	return res, err
}
