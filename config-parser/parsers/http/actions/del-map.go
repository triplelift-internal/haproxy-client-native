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

package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/client-native/v5/config-parser/common"
	"github.com/haproxytech/client-native/v5/config-parser/types"
)

type DelMap struct {
	FileName string
	KeyFmt   string
	Cond     string
	CondTest string
	Comment  string
}

func (f *DelMap) Parse(parts []string, parserType types.ParserType, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	f.FileName = strings.TrimPrefix(parts[1], "del-map(")
	f.FileName = strings.TrimRight(f.FileName, ")")
	if len(parts) >= 3 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) == 0 {
			return fmt.Errorf("not enough params")
		}
		f.KeyFmt = command[0]
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *DelMap) String() string {
	keyfmt := ""
	condition := ""
	if f.KeyFmt != "" {
		keyfmt = " " + f.KeyFmt
	}
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("del-map(%s)%s%s", f.FileName, keyfmt, condition)
}

func (f *DelMap) GetComment() string {
	return f.Comment
}
