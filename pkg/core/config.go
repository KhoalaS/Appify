package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type ProjectConfiguration struct {
	AppName          string   `json:"appName"`
	Website          string   `json:"website"`
	BlockedHosts     []string `json:"blockedHosts"`
	PackageName      string   `json:"packageName"`
	UserAgentString  string   `json:"userAgentString"`
	SslBypass        []string `json:"sslBypass"`
	Globals          string   `json:"globals"`
	OnloadScripts    string   `json:"onloadScripts"`
	ProjectDirectory string   `json:"projectDirectory"`
}

type TemplateProjectConfiguration struct {
	AppName          string `json:"appName"`
	Website          string `json:"website"`
	BlockedHosts     string `json:"blockedHosts"`
	PackageName      string `json:"packageName"`
	UserAgentString  string `json:"userAgentString"`
	SslBypass        string `json:"sslBypass"`
	Globals          string `json:"globals"`
	OnloadScripts    string `json:"onloadScripts"`
	ProjectDirectory string `json:"projectDirectory"`
}

func ReadConfigFromFile(inputFile string) (*ProjectConfiguration, error) {
	configFile, err := os.Open(inputFile)

	if err != nil {
		return nil, err
	}

	fileContent, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var config ProjectConfiguration
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *ProjectConfiguration) ToTemplateConfig() (*TemplateProjectConfiguration, error) {
	scriptFiles, err := os.ReadDir(config.OnloadScripts)
	scriptNames := []string{}

	for _, file := range scriptFiles {
		scriptNames = append(scriptNames, file.Name())
	}

	if err != nil {
		return nil, err
	}

	return &TemplateProjectConfiguration{
		AppName:          config.AppName,
		Website:          StringWithDoubleQuotes(config.Website),
		BlockedHosts:     SliceToKotlinListString(config.BlockedHosts),
		PackageName:      config.PackageName,
		UserAgentString:  StringWithDoubleQuotes(config.UserAgentString),
		SslBypass:        SliceToKotlinListString(config.SslBypass),
		Globals:          config.Globals,
		OnloadScripts:    SliceToKotlinListString(scriptNames),
		ProjectDirectory: config.ProjectDirectory,
	}, nil
}

func SliceToKotlinListString(values []string) string {
	valuesWithQuotes := []string{}
	for _, value := range values {
		valuesWithQuotes = append(valuesWithQuotes, StringWithDoubleQuotes(value))
	}

	return strings.Join(valuesWithQuotes, ", ")
}

func StringWithDoubleQuotes(inputValue string) string {
	return fmt.Sprintf("\"%s\"", inputValue)
}
