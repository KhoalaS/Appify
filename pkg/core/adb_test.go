package core

import "testing"

func TestAdbListDevices(t *testing.T) {
	adb := &ADB{}

	err := adb.Connect()
	if err != nil {
		t.Error(err)
	}
}
