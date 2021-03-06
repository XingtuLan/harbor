// Copyright Project Harbor Authors
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

package metadata

import (
	"testing"
)

func TestCfgMetaData_InitFromArray(t *testing.T) {
	testArray := []Item{
		{Scope: SystemScope, Group: BasicGroup, EnvKey: "HARBOR_ADMIN_PASSWORD", DefaultValue: "", Name: "admin_initial_password", ItemType: &PasswordType{}, Editable: true},
		{Scope: SystemScope, Group: BasicGroup, EnvKey: "ADMIRAL_URL", DefaultValue: "NA", Name: "admiral_url", ItemType: &StringType{}, Editable: false},
		{Scope: UserScope, Group: BasicGroup, EnvKey: "AUTH_MODE", DefaultValue: "db_auth", Name: "auth_mode", ItemType: &StringType{}, Editable: false},
		{Scope: SystemScope, Group: BasicGroup, EnvKey: "CFG_EXPIRATION", DefaultValue: "5", Name: "cfg_expiration", ItemType: &StringType{}, Editable: false},
		{Scope: SystemScope, Group: BasicGroup, EnvKey: "CHART_REPOSITORY_URL", DefaultValue: "http://chartmuseum:9999", Name: "chart_repository_url", ItemType: &StringType{}, Editable: false},
	}
	curInst := Instance()
	curInst.initFromArray(testArray)

	if len(metaDataInstance.metaMap) != 5 {
		t.Errorf("Can not initial metadata, size %v", len(metaDataInstance.metaMap))
	}
	item, ok := curInst.GetByName("admin_initial_password")
	if ok == false {
		t.Errorf("Can not get admin_initial_password metadata")
	}
	if item.Name != "admin_initial_password" {
		t.Errorf("Can not get admin_initial_password metadata")
	}

}

func TestCfgMetaData_Init(t *testing.T) {
	curInst := Instance()
	curInst.init()
	if len(metaDataInstance.metaMap) < 60 {
		t.Errorf("Can not initial metadata, size %v", len(metaDataInstance.metaMap))
	}
}
