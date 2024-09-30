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

package parsers

import (
	"fmt"

	"github.com/haproxytech/client-native/v5/config-parser/common"
	"github.com/haproxytech/client-native/v5/config-parser/errors"
	"github.com/haproxytech/client-native/v5/config-parser/types"
)

type LuaPrependPath struct {
	data        []types.LuaPrependPath
	preComments []string // comments that appear before the actual line
}

func (l *LuaPrependPath) parse(line string, parts []string, comment string) (*types.LuaPrependPath, error) {
	if len(parts) < 2 {
		return nil, &errors.ParseError{Parser: "LuaPrependPath", Line: line}
	}
	lpp := &types.LuaPrependPath{
		Path:    parts[1],
		Comment: comment,
	}
	if len(parts) > 2 {
		lpp.Type = parts[2]
	}
	return lpp, nil
}

func (l *LuaPrependPath) Result() ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.ErrFetch
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, data := range l.data {
		typ := ""
		if data.Type != "" {
			typ = fmt.Sprintf(" %s", data.Type)
		}
		result[index] = common.ReturnResultLine{
			Data:    fmt.Sprintf("lua-prepend-path %s%s", data.Path, typ),
			Comment: data.Comment,
		}
	}
	return result, nil
}
