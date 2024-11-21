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
	"github.com/haproxytech/client-native/v5/config-parser/errors"
	"github.com/haproxytech/client-native/v5/config-parser/types"
)

type ReplaceHeader struct {
	Name       string
	MatchRegex string
	ReplaceFmt string
	Cond       string
	CondTest   string
	Comment    string
}

func (f *ReplaceHeader) Parse(parts []string, parserType types.ParserType, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 5 {
		command, condition := common.SplitRequest(parts[2:])
		if len(command) < 3 {
			return errors.ErrInvalidData
		}
		f.Name = command[0]
		f.MatchRegex = command[1]
		f.ReplaceFmt = strings.Join(command[2:], "")
		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return errors.ErrInvalidData
}

func (f *ReplaceHeader) String() string {
	condition := ""
	if f.Cond != "" {
		condition = fmt.Sprintf(" %s %s", f.Cond, f.CondTest)
	}
	return fmt.Sprintf("replace-header %s %s %s%s", f.Name, f.MatchRegex, f.ReplaceFmt, condition)
}

func (f *ReplaceHeader) GetComment() string {
	return f.Comment
}
