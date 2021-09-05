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
	"reflect"
	"strings"
	"unicode"

	"github.com/spf13/viper"
)

// LoadConfig populates the fields of interface i with the values from json
// configuration files in the paths provided and from environment variables that
// match the interface field names.
//
// paths is an array of directories to look for configuration files.
// fileName is the name of the configuration file. The extension must be a type
// supported by the viper package (JSON, TOML, YAML, HCL, envfile and Java
// properties config).
// i is an instance of an interface to be populated from the files and
// environment variables identified.
//
// The method converts camel case field names to environment variables inserting
// an underscore before upper case characters and outputing all characters as
// uppercase. For example; the camel case field name ServicePath would become
// SERVICE_PATH.
// The interface that is being used for the configuration needs to be consulted
// because viper passes an upper case version of the field name into the method.
func LoadConfig(paths []string, fileName string, i interface{}) error {
	v := viper.New()
	for _, path := range paths {
		v.AddConfigPath(path)
	}
	v.SetConfigFile(fileName)
	v.AutomaticEnv()
	v.AllowEmptyEnv(false)
	for _, n := range getFields(reflect.TypeOf(i).Elem()) {
		err := v.BindEnv(n, convert(n))
		if err != nil {
			return err
		}
	}
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	return v.Unmarshal(i)
}

func convert(s string) string {
	b := strings.Builder{}
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i != 0 {
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToUpper(c))
	}
	return b.String()
}

func getFields(t reflect.Type) []string {
	a := []string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch f.Type.Kind() {
		case reflect.Struct:
			a = append(a, getFields(f.Type)...)
		default:
			a = append(a, f.Name)
		}
	}
	return a
}
