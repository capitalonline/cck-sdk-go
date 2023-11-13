package cck

import "testing"

func TestDescribeNasInstances(t *testing.T) {
	// nasID, nasName, nasSiteID, clusterID string, usageFlag bool, pageNumber, pageSize int
	nasID := ""
	nasName := ""
	nasSiteID := ""
	clusterID := ""
	usageFlag := 1
	pageNumber := 1
	pageSize := 1
	_, err1 := DescribeNasInstances(nasID, nasName, nasSiteID, clusterID, usageFlag, pageNumber, pageSize)
	if err1 != nil {
		t.Errorf("Failed, err is: %s", err1.Error())
	}

	nasID = "340d569c-7899-11ea-a06c-82b0d54620aa"
	nasName = "zbh-rancher"
	// nasSiteID = ""
	nasSiteID = "c4089dcd-15c2-4caf-9d1e-874770a31880"
	clusterID = "3b02449e-8843-11ea-988c-0242ac11034c"
	usageFlag = 0
	pageNumber = 2
	pageSize = 2
	_, err2 := DescribeNasInstances(nasID, nasName, nasSiteID, clusterID, usageFlag, pageNumber, pageSize)
	if err2 != nil {
		t.Errorf("Failed, err is: %s", err2.Error())
	}
}

func TestCreateNas(t *testing.T) {
	// nasSiteID, nasName, diskType, diskSize string
	nasSiteID := "c4089dcd-15c2-4caf-9d1e-874770a31880"
	nasName := "test-go-sdk"
	diskType := "high_disk"
	diskSize := 500
	_, err := CreateNas(nasSiteID, nasName, diskType, diskSize, 1)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestDeleteNas(t *testing.T) {
	// nasID string
	nasID := "b32f0172-8f7e-11ea-86a1-0242ac1103c0"
	_, err := DeleteNas(nasID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestMountNas(t *testing.T) {
	// nasID, clusterID string
	nasID := "340d569c-7899-11ea-a06c-82b0d54620aa"
	clusterID := "3b02449e-8843-11ea-988c-0242ac11034c"
	_, err := MountNas(nasID, clusterID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestUnMountNas(t *testing.T) {
	// nasID string
	nasID := "340d569c-7899-11ea-a06c-82b0d54620aa"
	_, err := UnMountNas(nasID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestResizeNas(t *testing.T) {
	// nasSiteID, nasID, diskSize string
	nasSiteID := "c4089dcd-15c2-4caf-9d1e-874770a31880"
	nasID := "340d569c-7899-11ea-a06c-82b0d54620aa"
	diskSize := 600
	_, err := ResizeNas(nasSiteID, nasID, diskSize)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestDescribeNasUsage(t *testing.T) {
	// cluserID, nasIP
	clusterID := ""
	nasIP := ""

	res, err := DescribeNasUsage(clusterID, nasIP)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
	if res.Data.NasInfo[0].UsageRate == "" {
		t.Errorf("Failed, UsageRate excepts not empty, but actual is empty ")
	}

}
