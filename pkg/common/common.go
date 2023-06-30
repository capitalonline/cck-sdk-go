package common

import (
	"os"
)

const (
	defaultApiHost         = "http://cdsapi.capitalonline.net"
	defaultApiHostOversea  = "http://cdsapi.capitalonline.net"
	apiHostLiteral         = "API_HOST"
	accessKeyIdLiteral     = "ACCESS_KEY_ID"
	accessKeySecretLiteral = "ACCESS_KEY_SECRET"
	overseaFlag            = "CDS_OVERSEA"
	cckProductType         = "cck"
	ebsProductType         = "ebs/v1"
	ecsProductType         = "ecs/v1"
	version                = "2019-08-08"
	signatureVersion       = "1.0"
	signatureMethod        = "HMAC-SHA1"
	timeStampFormat        = "2006-01-02T15:04:05Z"
)

const (
	// Nas
	ActionDescribeNasInstances = "DescribeNasInstances"
	ActionMountNas             = "MountNas"
	ActionUmountNas            = "UmountNas"
	ActionCreateNas            = "CreateNas"
	ActionResizeNas            = "ResizeNas"
	ActionDeleteNas            = "DeleteNas"
	ActionTaskStatus           = "CheckNasTaskStatus"

	// Disk
	ActionCreateDisk         = "CreateBlock"
	ActionAttachDisk         = "AttachBlock"
	ActionDetachDisk         = "DetachBlock"
	ActionDeleteDisk         = "DeleteBlock"
	ActionFindDiskByVolumeID = "DescribeBlock"
	ActionDiskTaskStatus     = "CheckBlockTaskStatus"
	ActionUpdateBlock        = "UpdateBlock"

	// Ebs
	ActionCreateEbs         = "CreateDisk"
	ActionDescribeEvent     = "DescribeEvent"
	ActionDeleteEbs         = "DeleteDisk"
	ActionAttachEbs         = "AttachDisk"
	ActionDetachEbs         = "DetachDisk"
	ActionDescribeEbs       = "DescribeDisk"
	ActionExtendEbs         = "ExtendDisk"
	ActionDescribeDiskQuota = "DescribeDiskQuota"

	EbsSuccessCode = "Success"
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

	// True is oversea cluster; False is domestic cluster
	if APIHost == "" {
		APIHost = defaultApiHost
		if os.Getenv(overseaFlag) == "True" {
			APIHost = defaultApiHostOversea
		}
	}
}

func IsAccessKeySet() bool {
	return AccessKeyID != "" && AccessKeySecret != ""
}
