// Copyright (C) 2023  Tricorder Observability
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package grafana

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

type AuthToken struct {
	client http.Client
}

func NewAuthToken() *AuthToken {
	return &AuthToken{
		client: http.Client{},
	}
}

// AuthKeyResult corresponds to the returned auth key after authenticating with Grafana.
// This is used for invoking Grafana APIs at corresponding API path.
type AuthKeyResult struct {
	// API name, essentially a URL path.
	Name string `json:"name"`
	Key  string `json:"key"`
}

// GetToken returns a auth token for a particular API path on Grafana (specified in apiPath argument).
func (g *AuthToken) GetToken(apiPath string) (*AuthKeyResult, error) {
	data := make(map[string]interface{})
	data["name"] = apiPath
	const adminRole = "Admin"
	data["role"] = adminRole

	bytesData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", CreateAuthKeysURI, bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(BasicAuth)))
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err == nil {
		var out AuthKeyResult
		err = json.Unmarshal(body, &out)
		if err != nil {
			return nil, err
		}
		return &out, nil
	}
	return nil, err
}
