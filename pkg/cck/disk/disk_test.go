package disk

import (
	"testing"
)

func TestCreateDisk(t *testing.T) {
	// params
	clusterID := ""
	regionID := ""
	fstype := ""
	diskType := ""
	requestGB := 500
	readOnly := false

	// api request
	res, err := CreateDisk(&CreateDiskArgs{
		ClusterID: clusterID,
		RegionID: regionID,
		Fstype: fstype,
		Type: diskType,
		RequestGB: requestGB,
		ReadOnly: readOnly,
	})

	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	// result
	if res.TaskID == "" || res.Data.VolumeID == "" {
		t.Errorf("Failed, [TaskID/VolumeID] is empty, but expectation is not empty")
	}

	// check DescribeTaskStatus
	_, err = DescribeTaskStatus(res.TaskID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

}

func TestAttachDisk(t *testing.T) {
	// params
	volumeID := ""
	nodeID := ""

	// api request
	res, err := AttachDisk(&AttachDiskArgs{
		VolumeID: volumeID,
		NodeID: nodeID,
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
	nodeID := ""

	// api request
	res, err := DetachDisk(&DetachDiskArgs{
		VolumeID: volumeID,
		NodeID: nodeID,
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
	volumeID := ""

	// api request
	res, err := FindDiskByVolumeID(&FindDiskByVolumeIDArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.Data.InstanceID == "" || res.Data.Status == "" {
		t.Errorf("Failed, [InstanceID/Status] is empty, but expectation is not empty")
	}

}

func TestFindDeviceNameByVolumeID (t *testing.T) {
	// params
	volumeID := ""

	// api request
	res, err := FindDeviceNameByVolumeID(&FindDeviceNameByVolumeIDArgs{
		VolumeID: volumeID,
	})

	// result
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}

	if res.Data.DeviceName == "" {
		t.Errorf("Failed, [DeviceName] is empty, but expectation is not empty")
	}

}