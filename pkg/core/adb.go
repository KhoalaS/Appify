package core

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type AppVariant string

const (
	DebugAppVariant   AppVariant = "debug"
	ReleaseAppVariant AppVariant = "release"
)

func (e *AppVariant) String() string {
	return string(*e)
}

func (e *AppVariant) Set(v string) error {
	switch v {
	case "debug", "release":
		*e = AppVariant(v)
		return nil
	default:
		return errors.New(`must be one of "debug" or "release"`)
	}
}

func (e *AppVariant) Type() string {
	return "AppVariant"
}

type ADB struct {
	connected      bool
	executablePath string
}

func (adb *ADB) Connect() error {
	adbPath, err := exec.LookPath("adb")
	if err != nil {
		return err
	}

	adb.executablePath = adbPath

	adbCommand := exec.Command(adbPath, "devices")
	stdout, err := adbCommand.Output()
	if err != nil {
		return err
	}

	if string(stdout) == "List of devices attached\n\n" {
		return errors.New("no devies attached")
	}

	adb.connected = true
	return nil
}

func (adb *ADB) InstallApp(projectDirectory string, variant AppVariant) error {
	if !adb.connected {
		return errors.New("[adb]: No Devices connected to adb")
	}

	baseBuildDir := filepath.Join(projectDirectory, "app", "build", "outputs", "apk")

	var apkPath string

	switch variant {
	case ReleaseAppVariant:
		apkPath = filepath.Join(baseBuildDir, "release", "app-release-unsigned.apk")
	default:
		apkPath = filepath.Join(baseBuildDir, "debug", "app-debug.apk")
	}

	adbCommand := exec.Command(adb.executablePath, "install", apkPath)
	_, err := adbCommand.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (adb *ADB) BuildApp(projectDirectory string, variant AppVariant) error {
	err := os.Chdir(projectDirectory)
	if err != nil {
		return err
	}

	var gradleWrapperPath string

	if runtime.GOOS == "windows" {
		gradleWrapperPath = "./gradlew.bat"
	} else {
		gradleWrapperPath = "./gradlew"
	}

	var gradleTask string

	switch variant {
	case ReleaseAppVariant:
		gradleTask = "app:assembleRelease"
	default:
		gradleTask = "app:assembleDebug"
	}

	gradleCommand := exec.Command(gradleWrapperPath, gradleTask)
	err = gradleCommand.Run()
	if err != nil {
		return err
	}

	err = os.Chdir("../")
	if err != nil {
		return err
	}

	return nil
}

func (adb *ADB) StartApp(packageName string) error {
	if !adb.connected {
		return errors.New("[adb]: No Devices connected to adb")
	}

	activityName := fmt.Sprintf("%s/%s.MainActivity", packageName, packageName)

	adbCommand := exec.Command(adb.executablePath, "shell", "am", "start", "-n", activityName)
	_, err := adbCommand.Output()
	if err != nil {
		return err
	}

	return nil
}
