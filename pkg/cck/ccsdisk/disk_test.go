package ccsdisk

import (
	"testing"
)

func TestCreateDisk(t *testing.T) {
	// params
	regionID := "c2bfebe1-63e7-48b1-86df-e7a0b2c44ed0"
	iops := 3000
	diskType := "high_disk"
	sise := 500
	clusterName := "POD39-CLU03"

	// api request
	res, err := CreateDisk(&CreateDiskArgs{
		RegionID:    regionID,
		DiskType:    diskType,
		Iops:        iops,
		Size:        sise,
		ClusterName: clusterName,
	})

	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.Data.VolumeID == "" {
		t.Errorf("Failed, [TaskID/VolumeID] is empty, but expectation is not empty")
	}

	resStatus, err := GetDiskInfo(&DiskInfoArgs{
		VolumeID: res.Data.VolumeID,
	})
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if resStatus.Data.Status == "" {
		t.Errorf("Failed, Status is empty")
	}
}

func TestAttachDisk(t *testing.T) {
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

func TestGetDiskInfo(t *testing.T) {
	// params
	volumeID := "5a694727-57ef-4888-ada4-40316cc99564"

	// api request
	res, err := GetDiskInfo(&DiskInfoArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.Data.NodeID == "" || res.Data.Status == "" || res.Data.VolumeId == "" {
		t.Errorf("Failed, [NodeID/Status/Uuid/IsFormat] is empty, but expectation is not empty")
	}
}
