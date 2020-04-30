package cck

import (
	"testing"
)

func TestDescribeNasInstances(t *testing.T) {
	//nasID, nasName, nasSiteID, clusterID string, usageFlag bool, pageNumber, pageSize int
	nasID := ""
	nasName := ""
	nasSiteID := ""
	clusterID := ""
	usageFlag := true
	pageNumber := 1
	pageSize := 2
	_, err := DescribeNasInstances(nasID, nasName, nasSiteID, clusterID, usageFlag, pageNumber, pageSize)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestCreateNas(t *testing.T){
	// nasSiteID, nasName, diskType, diskSize string
	nasSiteID := ""
	nasName := ""
	diskType := ""
	diskSize := ""
	_, err := CreateNas(nasSiteID, nasName, diskType, diskSize)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestResizeNas(t *testing.T){
	// nasSiteID, nasID, diskSize string
	nasSiteID := ""
	nasID := ""
	diskSize := ""
	_, err := ResizeNas(nasSiteID, nasID, diskSize)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestDeleteNas(t *testing.T) {
	// nasID string
	nasID := ""
	_, err := DeleteNas(nasID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestMountNas(t *testing.T) {
	// nasID, clusterID string
	nasID := ""
	clusterID := ""
	_, err := MountNas(nasID, clusterID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}

func TestUnMountNas(t *testing.T) {
	// nasID string
	nasID := ""
	_, err := UnMountNas(nasID)
	if err != nil {
		t.Errorf("Failed, err is: %s", err.Error())
	}
}