package common

import (
	"os"
)

const (
	defaultApiHost         = "http://cdsapi.capitalonline.net"
	apiHostLiteral         = "API_HOST"
	accessKeyIdLiteral     = "ACCESS_KEY_ID"
	accessKeySecretLiteral = "ACCESS_KEY_SECRET"
	cckProductType         = "CCK"
	version                = "2020-05-30"
	signatureVersion       = "1.0"
	signatureMethod        = "HMAC-SHA1"
	timeStampFormat        = "2006-01-02T15:04:05Z"
)

const (
	ActionMountNas = "MountNas"
)

var (
	APIHost         string
	AccessKeyID     string
	AccessKeySecret string
)

func init() {
	if APIHost == "" {
		APIHost = os.Getenv(apiHostLiteral)
	}
	if AccessKeyID == "" {
		AccessKeyID = os.Getenv(accessKeyIdLiteral)
	}
	if AccessKeySecret == "" {
		AccessKeySecret = os.Getenv(accessKeySecretLiteral)
	}

	if APIHost == "" {
		APIHost = defaultApiHost
	}
}

func IsAccessKeySet() bool {
	return AccessKeyID != "" && AccessKeySecret != ""
}
