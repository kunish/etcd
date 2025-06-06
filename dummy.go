// Copyright 2016 The etcd Authors
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

package main_test

// MainTest package makes sure these packages stay as dependencies of the root
// module (e.g. for sake of 'bom' generation).
// Thanks to this 'go mod tidy' is not removing that dependencies from go.mod.
import (
	_ "go.etcd.io/etcd/etcdctl/v3/ctlv3/command" // keep
	_ "go.etcd.io/etcd/etcdutl/v3/etcdutl"       // keep
	_ "go.etcd.io/etcd/tests/v3/integration"     // keep
)
