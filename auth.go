/*
	Go library for attlassians confluence wiki

	Copyright (C) 2017 Carsten Seeger

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	@author Carsten Seeger
	@copyright Copyright (C) 2017 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/confluence-go-api
*/

package goconfluence

import (
	"errors"
	"net/http"
	"strings"
)

// NewAPI implements API constructor
// password can be user password or user api-token
func NewAPI(url string, username string, password string) (*API, error) {
	if len(url) == 0 || len(username) == 0 || len(password) == 0 {
		return nil, errors.New("url, username or password empty")
	}

	u, err := url.ParseRequestURI(location)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}

	u.Path += "rest/api"

	a := new(API)
	a.endPoint = u
	a.password = password
	a.username = username
	a.client = &http.Client{}

	return a, nil
}

// Auth implements basic auth
func (a *API) Auth(req *http.Request) {
	req.SetBasicAuth(a.username, a.password)
}
