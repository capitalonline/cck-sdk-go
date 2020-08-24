package disk

import (
	"testing"
)

func TestCreateDisk(t *testing.T) {
	// params
	name := "testing-go-test-001"
	regionID := "c2bfebe1-63e7-48b1-86df-e7a0b2c44ed0"
	iops := 3000
	diskType := "high_disk"
	sise := 500
	zoneID := "POD39-CLU03"

	// api request
	res, err := CreateDisk(&CreateDiskArgs{
		Name:     name,
		RegionID: regionID,
		DiskType: diskType,
		Iops:     iops,
		Size:     sise,
		ZoneID:   zoneID,
	})

	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	// result
	if res.TaskID == "" || res.Data.VolumeID == "" {
		t.Errorf("Failed, [TaskID/VolumeID] is empty, but expectation is not empty")
	}

	// check DescribeTaskStatus
	resStatus, err := DescribeTaskStatus(res.TaskID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	// check status
	if resStatus.Data.Status == "" {
		t.Errorf("Failed, Status is empty")
	}

}

func TestAttachDisk(t *testing.T) {
	// params
	volumeID := ""
	nodeID := ""

	// api request
	res, err := AttachDisk(&AttachDiskArgs{
		VolumeID: volumeID,
		NodeID:   nodeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.TaskID == "" {
		t.Errorf("Failed, [TaskID] is empty, but expectation is not empty")
	}

}

func TestDetachDisk(t *testing.T) {
	// params
	volumeID := ""

	// api request
	res, err := DetachDisk(&DetachDiskArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.TaskID == "" {
		t.Errorf("Failed, [TaskID] is empty, but expectation is not empty")
	}

}

func TestDeleteDisk(t *testing.T) {
	// params
	volumeID := ""

	// api request
	res, err := DeleteDisk(&DeleteDiskArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.TaskID == "" {
		t.Errorf("Failed, [TaskID] is empty, but expectation is not empty")
	}
}

func TestFindDiskByVolumeID(t *testing.T) {
	// params
	volumeID := "5a694727-57ef-4888-ada4-40316cc99564"

	// api request
	res, err := FindDiskByVolumeID(&FindDiskByVolumeIDArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.Data[0].NodeID == "" || res.Data[0].Status == "" || res.Data[0].Uuid == "" {
		t.Errorf("Failed, [NodeID/Status/Uuid] is empty, but expectation is not empty")
	}

}

func TestDescribeTaskStatus(t *testing.T) {
	// params
	taskID := "c971c21c-e5f0-11ea-a2ed-7abe85ed524b"

	// api request
	res, err := DescribeTaskStatus(taskID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	// check status
	if res.Data.Status == "" {
		t.Errorf("Failed, Status is empty")



	}

