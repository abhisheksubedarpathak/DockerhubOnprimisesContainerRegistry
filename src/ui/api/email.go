// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"net"
	"net/http"
	"strconv"

	"github.com/vmware/harbor/src/common/utils/email"
	"github.com/vmware/harbor/src/common/utils/log"
	"github.com/vmware/harbor/src/ui/config"
)

const (
	pingEmailTimeout = 60
)

// EmailAPI ...
type EmailAPI struct {
	BaseController
}

// Prepare ...
func (e *EmailAPI) Prepare() {
	e.BaseController.Prepare()
	if !e.SecurityCtx.IsAuthenticated() {
		e.HandleUnauthorized()
		return
	}

	if !e.SecurityCtx.IsSysAdmin() {
		e.HandleForbidden(e.SecurityCtx.GetUsername())
		return
	}
}

// Ping tests connection and authentication with email server
func (e *EmailAPI) Ping() {
	var host, username, password, identity string
	var port int
	var ssl, insecure bool
	body := e.Ctx.Input.CopyBody(1 << 32)
	if body == nil || len(body) == 0 {
		cfg, err := config.Email()
		if err != nil {
			log.Errorf("failed to get email configurations: %v", err)
			e.CustomAbort(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError))
		}
		host = cfg.Host
		port = cfg.Port
		username = cfg.Username
		password = cfg.Password
		identity = cfg.Identity
		ssl = cfg.SSL
		insecure = cfg.Insecure
	} else {
		settings := &struct {
			Host     string  `json:"email_host"`
			Port     *int    `json:"email_port"`
			Username string  `json:"email_username"`
			Password *string `json:"email_password"`
			SSL      bool    `json:"email_ssl"`
			Identity string  `json:"email_identity"`
			Insecure bool    `json:"email_insecure"`
		}{}
		e.DecodeJSONReq(&settings)

		if len(settings.Host) == 0 || settings.Port == nil {
			e.CustomAbort(http.StatusBadRequest, "empty host or port")
		}

		if settings.Password == nil {
			cfg, err := config.Email()
			if err != nil {
				log.Errorf("failed to get email configurations: %v", err)
				e.CustomAbort(http.StatusInternalServerError,
					http.StatusText(http.StatusInternalServerError))
			}

			settings.Password = &cfg.Password
		}

		host = settings.Host
		port = *settings.Port
		username = settings.Username
		password = *settings.Password
		identity = settings.Identity
		ssl = settings.SSL
		insecure = settings.Insecure
	}

	addr := net.JoinHostPort(host, strconv.Itoa(port))
	if err := email.Ping(addr, identity, username,
		password, pingEmailTimeout, ssl, insecure); err != nil {
		log.Debugf("ping %s failed: %v", addr, err)
		e.CustomAbort(http.StatusBadRequest, err.Error())
	}
}
