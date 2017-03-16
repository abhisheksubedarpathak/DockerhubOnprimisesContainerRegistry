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

package api

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingEmail(t *testing.T) {
	fmt.Println("Testing ping email server")
	assert := assert.New(t)
	apiTest := newHarborAPI()

	//case 1: ping email server without admin role
	code, _, err := apiTest.PingEmail(*testUser, nil)
	if err != nil {
		t.Errorf("failed to test ping email server: %v", err)
		return
	}

	assert.Equal(401, code, "the status code of ping email server with non-admin user should be 401")

	//case 2: bad request
	settings := `{
		"email_host":     ""
	}`

	code, _, err = apiTest.PingEmail(*admin, []byte(settings))
	if err != nil {
		t.Errorf("failed to test ping email server: %v", err)
		return
	}

	assert.Equal(400, code, "the status code of ping email server should be 400")

	//case 3: secure connection with admin role
	settings = `{
		"email_host":     "smtp.gmail.com",
		"email_port":     465,
		"email_identity": "",
		"email_username": "wrong_username",
		"email_ssl":      true
	}`

	code, body, err := apiTest.PingEmail(*admin, []byte(settings))
	if err != nil {
		t.Errorf("failed to test ping email server: %v", err)
		return
	}

	assert.Equal(400, code, "the status code of ping email server should be 400")

	if !strings.Contains(body, "535") {
		t.Errorf("unexpected error: %s does not contains 535", body)
		return
	}
}
