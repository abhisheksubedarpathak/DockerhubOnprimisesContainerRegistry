/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// Package config provide methods to get the configurations reqruied by code in src/common
package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetOfKeyFileProvider(t *testing.T) {
	path := "/tmp/secretkey"
	key := "secretkey"

	if err := ioutil.WriteFile(path, []byte(key), 0777); err != nil {
		t.Errorf("failed to write to file %s: %v", path, err)
		return
	}
	defer os.Remove(path)

	provider := NewKeyFileProvider(path)
	k, err := provider.Get()
	if err != nil {
		t.Errorf("failed to get key from the file provider: %v", err)
		return
	}

	if k != key {
		t.Errorf("unexpected key: %s != %s", k, key)
		return
	}
}
