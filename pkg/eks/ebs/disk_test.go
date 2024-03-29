package ebs

import (
	"fmt"
	_ "github.com/capitalonline/cck-sdk-go/pkg/common"
	"testing"
)

func TestCreateEbs(t *testing.T) {
	var req = CreateEbsReq{
		//AvailableZoneCode: "CN_DaBieShan_A",
		AvailableZoneCode: "SR_SaoPaulo_A",
		DiskName:          "dyl-test",
		DiskFeature:       "SSD",
		Size:              64,
		Number:            nil,
		BillingMethod:     "",
	}
	resp, err := CreateEbs(&req)
	fmt.Println(resp, err)
}

func TestDescribeTaskStatus(t *testing.T) {
	DescribeTaskStatus("60dbf132-17d4-11ee-843e-16ecbf5c1698")
}

func TestDeleteEbs(t *testing.T) {
	req := DeleteDiskArgs{
		DiskIds: []string{
			"disk-tw669d2s9ofab26y",
		},
	}
	DeleteDisk(&req)
}

func TestAttachDisk(t *testing.T) {
	var req = AttachDiskArgs{
		DiskIds:             []string{"disk-2hjcw8fszoaap2oy"},
		InstanceId:          "disk-tw669d2s9ofab26y",
		ReleaseWithInstance: 0,
	}
	AttachDisk(&req)
}

func TestDetachDisk(t *testing.T) {
	var req = DetachDiskArgs{
		DiskIds: []string{"disk-tw669d2s9ofab26y"},
	}
	DetachDisk(&req)
}

func TestFindDiskByVolumeID(t *testing.T) {
	var req = &FindDiskByVolumeIDArgs{
		DiskId: "disk-tw669d2s9ofab26y",
	}
	FindDiskByVolumeID(req)
}

func TestExtendDisk(t *testing.T) {
	var req = &ExtendDiskArgs{
		DiskId:       "disk-xo4c1husoooa025y",
		ExtendedSize: 72,
	}
	ExtendDisk(req)
}

func TestDescribeDiskQuota(t *testing.T) {
	DescribeDiskQuota("SR_SaoPaulo_A")
}

func TestDescribeInstance(t *testing.T) {
	res, er := DescribeInstance("ins-r4h21z8s00mvji7g")
	fmt.Println(res.Data.Disk)
	fmt.Println(res, er)
}
