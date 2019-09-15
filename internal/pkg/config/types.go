/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config interface {
	Reader
	Writer
	GetEnvironmentConfig() Reader
}

type Reader interface {
	GetString(key string) string
	GetStringOrDefault(key string, defaultValue string) string
	GetStringForKeyEntry(key KeyEntry) string
}

type KeyEntry struct {
	Key          string
	DefaultValue string
}

type ViperConfig struct {
	userConfigManager *ViperManager
	envConfigManager *ViperManager
}

func (cliConfig *ViperConfig) GetEnvironmentConfig() Reader {
	return cliConfig.envConfigManager
}

func (cliConfig *ViperConfig) GetString(key string) string {
	return cliConfig.userConfigManager.GetString(key)
}

func (cliConfig *ViperConfig) GetStringOrDefault(key string, defaultValue string) string {
	return cliConfig.userConfigManager.GetStringOrDefault(key, defaultValue)
}

func (cliConfig *ViperConfig) GetStringForKeyEntry(keyEntry KeyEntry) string {
	return cliConfig.userConfigManager.GetStringForKeyEntry(keyEntry)
}

type ViperManager struct {
	viperInstance *viper.Viper
}

func (cliConfig *ViperManager) GetString(key string) string {
	value := cliConfig.viperInstance.GetString(key)
	return value
}

func (cliConfig *ViperManager) GetStringOrDefault(key string, defaultValue string) string {
	if value := cliConfig.GetString(key); value != "" {
		return value
	} else {
		return defaultValue
	}
}

func (cliConfig *ViperManager) GetStringForKeyEntry(keyEntry KeyEntry) string {
	return cliConfig.GetStringOrDefault(keyEntry.Key, keyEntry.DefaultValue)
}

type Writer interface {
	SetString(key string, value string)
	SetStringForKeyEntry(keyEntry KeyEntry, value string)
}

func (cliConfig *ViperManager) SetString(key string, value string) {
	cliConfig.viperInstance.Set(key, value)
	if err := cliConfig.viperInstance.WriteConfig(); err != nil {
		fmt.Printf("Could not write to config. Key: %s\n", key)
	}
}

func (cliConfig *ViperManager) SetStringForKeyEntry(keyEntry KeyEntry, value string) {
	cliConfig.SetString(keyEntry.Key, value)
}

func (cliConfig *ViperConfig) SetString(key string, value string) {
	cliConfig.userConfigManager.SetString(key, value)
}

func (cliConfig *ViperConfig) SetStringForKeyEntry(keyEntry KeyEntry, value string) {
	cliConfig.userConfigManager.SetStringForKeyEntry(keyEntry, value)
}