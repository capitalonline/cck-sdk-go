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
	DescribeTaskStatus("645c8c30-1640-11ee-843e-16ecbf5c1698")
}

func TestDeleteEbs(t *testing.T) {
	req := DeleteDiskArgs{
		DiskIds: []string{
			"disk-gpr4uegsooaa62jy",
		},
	}
	DeleteDisk(&req)
}

func TestAttachDisk(t *testing.T) {
	var req = AttachDiskArgs{
		DiskIds:             []string{"disk-2hjcw8fszoaap2oy"},
		InstanceId:          "ins-bmg6whbsvzhu8xg5",
		ReleaseWithInstance: 0,
	}
	AttachDisk(&req)
}

func TestDetachDisk(t *testing.T) {
	var req = DetachDiskArgs{
		DiskIds: []string{"disk-xo4c1husoooa025y"},
	}
	DetachDisk(&req)
}

func TestFindDiskByVolumeID(t *testing.T) {
	var req = &FindDiskByVolumeIDArgs{
		DiskId: "disk-xo4c1husoooa025y",
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
	DescribeInstance("ins-7z34duds5mm7dln3")
}
