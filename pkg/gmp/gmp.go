// Copyright 2022 zenpk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gmp

import (
	"errors"
	"os"
)

// GetModPath returns the path contains go.mod
func GetModPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir += "/"
	for true {
		files, err := os.ReadDir(dir)
		if err != nil {
			return "", err
		}
		for _, file := range files {
			if file.Name() == "go.mod" {
				return dir, nil
			}
		}
		dir += "../"
	}
	return "", errors.New("failed")
}
