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

import (
	"fmt"
	"testing"
)

type Override struct {
	Common    `mapstructure:",squash"`
	LocalFile string `mapstructure:"localFile"`
}

func TestCamelConvertEmpty(t *testing.T) {
	i := ""
	e := ""
	if convert(i) != e {
		t.Fatal("Camel convert not failed")
	}
}

func TestCamelConvertNone(t *testing.T) {
	i := "None"
	e := "NONE"
	if convert(i) != e {
		t.Fatal("Camel convert none failed")
	}
}

func TestCamelConvertSingle(t *testing.T) {
	i := "LocalFile"
	e := "LOCAL_FILE"
	if convert(i) != e {
		t.Fatal("Camel convert single failed")
	}
}

func TestCamelConvertMultiple(t *testing.T) {
	i := "AzureStorageAccount"
	e := "AZURE_STORAGE_ACCOUNT"
	if convert(i) != e {
		t.Fatal("Camel convert multiple failed")
	}
}

func TestLocalConfigurationSettings(t *testing.T) {
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.local", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.LocalFile == "" {
		t.Error("Local file not set")
		return
	}
}

func TestLocalConfigurationEnvironment(t *testing.T) {
	e := "TEST ENV LOCAL FILE"
	t.Setenv("LOCAL_FILE", e)
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.local", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.LocalFile != e {
		t.Error("Local file not expected value")
		return
	}
}

func TestAwsConfigurationSettingsTrue(t *testing.T) {
	testAwsConfigurationSettings(t, true)
}

func TestAwsConfigurationSettingsFalse(t *testing.T) {
	testAwsConfigurationSettings(t, false)
}

func testAwsConfigurationSettings(t *testing.T, expected bool) {
	e := fmt.Sprintf("%t", expected)
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.aws."+e, &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.AwsEnabled != expected {
		t.Errorf("AWS Enabled not '%t'", expected)
		return
	}
}

func TestAwsConfigurationEnvironmentTrue(t *testing.T) {
	testAwsConfigurationEnvironment(t, true)
}

func TestAwsConfigurationEnvironmentFalse(t *testing.T) {
	testAwsConfigurationEnvironment(t, false)
}

func testAwsConfigurationEnvironment(t *testing.T, expected bool) {
	e := fmt.Sprintf("%t", expected)
	t.Setenv("AWS_ENABLED", e)
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.none", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.AwsEnabled != expected {
		t.Errorf("AWS Enabled not '%s'", e)
		return
	}
}

func TestGcpConfigurationSettings(t *testing.T) {
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.gcp", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.GcpProject == "" {
		t.Error("GCP Project not set")
		return
	}
}

func TestGcpConfigurationEnvironment(t *testing.T) {
	e := "PROJECT NAME"
	t.Setenv("GCP_PROJECT", e)
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.none", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.GcpProject != e {
		t.Error("GCP Project not expected value")
		return
	}
}

func TestAzureConfigurationSettings(t *testing.T) {
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.azure", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.AzureStorageAccount == "" || c.AzureStorageAccessKey == "" {
		t.Error("Azure not set")
		return
	}
}

func TestAzureConfigurationEnvironment(t *testing.T) {
	ea := "ACCOUNT"
	ek := "KEY"
	t.Setenv("AZURE_STORAGE_ACCOUNT", ea)
	t.Setenv("AZURE_STORAGE_ACCESS_KEY", ek)
	c := Override{}
	err := LoadConfig([]string{"."}, "appsettings.test.none", &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.AzureStorageAccount != ea || c.AzureStorageAccessKey != ek {
		t.Error("Azure not expected value")
		return
	}
}
