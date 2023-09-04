// Code generated with struct_equal_generator; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

import (
	"encoding/json"
	"testing"

	"github.com/go-faker/faker/v4"

	jsoniter "github.com/json-iterator/go"
)

func TestStickTableEqual(t *testing.T) {
	samples := []struct {
		a, b StickTable
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTable
		var result StickTable
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		byteJSON, err := json.Marshal(sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = json.Unmarshal(byteJSON, &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		samples = append(samples, struct {
			a, b StickTable
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Equal(sample.b)
		if !result {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTable to be equal, but it is not %s %s", a, b)
		}
	}
}

func TestStickTableEqualFalse(t *testing.T) {
	samples := []struct {
		a, b StickTable
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTable
		var result StickTable
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = faker.FakeData(&result)
		if err != nil {
			t.Errorf(err.Error())
		}
		result.Process = Ptr(*sample.Process + 1)
		result.Size = Ptr(*sample.Size + 1)
		result.Used = Ptr(*sample.Used + 1)
		samples = append(samples, struct {
			a, b StickTable
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Equal(sample.b)
		if result {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTable to be different, but it is not %s %s", a, b)
		}
	}
}

func TestStickTableDiff(t *testing.T) {
	samples := []struct {
		a, b StickTable
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTable
		var result StickTable
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		byteJSON, err := json.Marshal(sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = json.Unmarshal(byteJSON, &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		samples = append(samples, struct {
			a, b StickTable
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Diff(sample.b)
		if len(result) != 0 {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTable to be equal, but it is not %s %s, %v", a, b, result)
		}
	}
}

func TestStickTableDiffFalse(t *testing.T) {
	samples := []struct {
		a, b StickTable
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTable
		var result StickTable
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = faker.FakeData(&result)
		if err != nil {
			t.Errorf(err.Error())
		}
		result.Process = Ptr(*sample.Process + 1)
		result.Size = Ptr(*sample.Size + 1)
		result.Used = Ptr(*sample.Used + 1)
		samples = append(samples, struct {
			a, b StickTable
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Diff(sample.b)
		if len(result) != 6 {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTable to be different in 6 cases, but it is not (%d) %s %s", len(result), a, b)
		}
	}
}

func TestStickTableFieldEqual(t *testing.T) {
	samples := []struct {
		a, b StickTableField
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTableField
		var result StickTableField
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		byteJSON, err := json.Marshal(sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = json.Unmarshal(byteJSON, &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		samples = append(samples, struct {
			a, b StickTableField
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Equal(sample.b)
		if !result {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTableField to be equal, but it is not %s %s", a, b)
		}
	}
}

func TestStickTableFieldEqualFalse(t *testing.T) {
	samples := []struct {
		a, b StickTableField
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTableField
		var result StickTableField
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = faker.FakeData(&result)
		if err != nil {
			t.Errorf(err.Error())
		}
		result.Period = sample.Period + 1
		samples = append(samples, struct {
			a, b StickTableField
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Equal(sample.b)
		if result {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTableField to be different, but it is not %s %s", a, b)
		}
	}
}

func TestStickTableFieldDiff(t *testing.T) {
	samples := []struct {
		a, b StickTableField
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTableField
		var result StickTableField
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		byteJSON, err := json.Marshal(sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = json.Unmarshal(byteJSON, &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		samples = append(samples, struct {
			a, b StickTableField
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Diff(sample.b)
		if len(result) != 0 {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTableField to be equal, but it is not %s %s, %v", a, b, result)
		}
	}
}

func TestStickTableFieldDiffFalse(t *testing.T) {
	samples := []struct {
		a, b StickTableField
	}{}
	for i := 0; i < 2; i++ {
		var sample StickTableField
		var result StickTableField
		err := faker.FakeData(&sample)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = faker.FakeData(&result)
		if err != nil {
			t.Errorf(err.Error())
		}
		result.Period = sample.Period + 1
		samples = append(samples, struct {
			a, b StickTableField
		}{sample, result})
	}

	for _, sample := range samples {
		result := sample.a.Diff(sample.b)
		if len(result) != 3 {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			a, err := json.Marshal(&sample.a)
			if err != nil {
				t.Errorf(err.Error())
			}
			b, err := json.Marshal(&sample.b)
			if err != nil {
				t.Errorf(err.Error())
			}
			t.Errorf("Expected StickTableField to be different in 3 cases, but it is not (%d) %s %s", len(result), a, b)
		}
	}
}
