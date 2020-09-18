package block

import (
	"testing"
)

func TestCreateDisk(t *testing.T) {
	// params
	name := "testing-go-test-001"
	regionID := "c2bfebe1-63e7-48b1-86df-e7a0b2c44ed0"
	iops := 1000
	diskType := "ssd"
	size := 500
	bandwidth := 500

	// api request
	res, err := CreateBlock(&CreateBlockArgs{
		Name:      name,
		RegionID:  regionID,
		DiskType:  diskType,
		Iops:      iops,
		Size:      size,
		Bandwidth: bandwidth,
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

func TestAttachBlock(t *testing.T) {
	// params
	volumeID := ""
	nodeID := ""

	// api request
	res, err := AttachBlock(&AttachBlockArgs{
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

func TestDetachBlock(t *testing.T) {
	// params
	volumeID := ""

	// api request
	res, err := DetachBlock(&DetachBlockArgs{
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

func TestDeleteBlock(t *testing.T) {
	// params
	volumeID := ""

	// api request
	res, err := DeleteBlock(&DeleteBlockArgs{
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
	res, err := FindBlockByVolumeID(&FindBlockByVolumeIDArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.Data.BlockSlice[0].NodeID == "" || res.Data.BlockSlice[0].Status == "" || res.Data.BlockSlice[0].Uuid == "" {
		t.Errorf("Failed, [NodeID/Status/Uuid/IsFormat] is empty, but expectation is not empty")

	}

}
