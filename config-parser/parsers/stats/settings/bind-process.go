/*
Copyright 2019 HAProxy Technologies

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

package stats

import (
	"fmt"

	parsers "github.com/haproxytech/client-native/v5/config-parser/parsers"
)

type BindProcess struct {
	BindProcess *parsers.BindProcess
}

func (m *BindProcess) Parse(parts []string, comment string) error {
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}

	m.BindProcess = &parsers.BindProcess{}
	_, err := m.BindProcess.Parse("", parts[1:], comment)
	if err != nil {
		return fmt.Errorf("error parsing bind-process")
	}
	return nil
}

func (m *BindProcess) String() string {
	res, _ := m.BindProcess.Result()
	if len(res) != 0 {
		return res[0].Data
	}
	return "bind-process"
}

func (m *BindProcess) GetComment() string {
	res, _ := m.BindProcess.Result()
	if len(res) != 0 {
		return res[0].Comment
	}
	return ""
}
