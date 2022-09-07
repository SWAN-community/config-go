/* ****************************************************************************
 * Copyright 2021 51 Degrees Mobile Experts Limited (51degrees.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 * ***************************************************************************/

package config

// Configuration details from appsettings.json for access to the AWS, GCP, Azure
// or local file storage.
type Base struct {
	AzureStorageAccount   string `mapstructure:"azureStorageAccount"`
	AzureStorageAccessKey string `mapstructure:"azureStorageAccessKey"`
	GcpProject            string `mapstructure:"gcpProject"`
	AwsEnabled            bool   `mapstructure:"awsEnabled"`
	Scheme                string `mapstructure:"scheme"` // Either https or http
	Debug                 bool   `mapstructure:"debug"`  // True for debug logging
	HttpPlatformPort      int    `mapstructure:"httpPlatformPort"`
	HttpsPlatformPort     int    `mapstructure:"httpsPlatformPort"`
	Port                  int    `mapstructure:"port"`
}

func (b *Base) GetPortHTTP() int {
	var port int
	if b.HttpPlatformPort != 0 {
		// Get the port environment variable from Azure App Services.
		port = b.HttpPlatformPort
	} else if b.Port != 0 {
		// Get the port environment variable from Amazon Web Services.
		port = b.Port
	}
	return port
}

func (b *Base) GetPortHTTPS() int {
	var port int
	if b.HttpsPlatformPort != 0 {
		// Get the port environment variable from Azure App Services.
		port = b.HttpsPlatformPort
	}
	return port
}
