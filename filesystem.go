//Copyright 2016 lyobzik
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package utils

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"regexp"
)

func EnsureDir(path string) error {
	return os.MkdirAll(path, os.ModeDir|0777)
}

func EnsureDirs(paths ...string) error {
	for _, path := range paths {
		if err := EnsureDir(path); err != nil {
			return errors.Wrapf(err, "cannot create directory '%s'", path)
		}
	}
	return nil
}

func GetFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read file list in '%s'", path)
	}

	fileNames := make([]string, 0)
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func GetFilteredFiles(path string, filter string) ([]string, error) {
	files, err := GetFiles(path)
	if err != nil {
		return nil, err
	}
	filterRegexp, err := regexp.Compile(filter)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot compile filter regexp '%s'", filter)
	}
	filteredFiles := make([]string, 0, len(files))
	for _, file := range files {
		if filterRegexp.MatchString(file) {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles, nil
}

func IsExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}