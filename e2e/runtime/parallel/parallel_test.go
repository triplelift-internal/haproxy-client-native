// Copyright 2021 HAProxy Technologies
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
//go:build integration
// +build integration

package parallel_test

import (
	"sync"
)

func (s *ParallelRuntime) TestParallel() {
	runtime, err := s.client.Runtime()
	if err != nil {
		s.FailNow(err.Error())
	}
	var wg sync.WaitGroup
	for range 3000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			info, err := runtime.GetInfo()
			s.Assert().Greater(len(info), 0, "info is empty")
			s.Assert().NotNil(info[0].Info, "information is nil", info[0].Error)
			s.Assert().Equal(info[0].RuntimeAPI, s.socketPath, "runtime not correct, runtime is ", info[0].RuntimeAPI)
			s.Assert().Contains(info[0].Info.Version, s.haproxyVersion, "version not correct, version is ", info[0].Info.Version)
			if err != nil {
				s.FailNow(err.Error())
			}
		}()
	}
	wg.Wait()
}
