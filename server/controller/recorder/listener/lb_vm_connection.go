/*
 * Copyright (c) 2024 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package listener

import (
	cloudmodel "github.com/deepflowio/deepflow/server/controller/cloud/model"
	metadbmodel "github.com/deepflowio/deepflow/server/controller/db/metadb/model"
	"github.com/deepflowio/deepflow/server/controller/recorder/cache"
	"github.com/deepflowio/deepflow/server/controller/recorder/cache/diffbase"
)

type LBVMConnection struct {
	cache *cache.Cache
}

func NewLBVMConnection(c *cache.Cache) *LBVMConnection {
	listener := &LBVMConnection{
		cache: c,
	}
	return listener
}

func (c *LBVMConnection) OnUpdaterAdded(addedDBItems []*metadbmodel.LBVMConnection) {
	c.cache.AddLBVMConnections(addedDBItems)
}

func (c *LBVMConnection) OnUpdaterUpdated(cloudItem *cloudmodel.LBVMConnection, diffBase *diffbase.LBVMConnection) {
}

func (c *LBVMConnection) OnUpdaterDeleted(lcuuids []string) {
	c.cache.DeleteLBVMConnections(lcuuids)
}
