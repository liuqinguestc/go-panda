// Copyright (c) 2018 Huawei Technologies Co., Ltd. All Rights Reserved.
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

package crontrigger

import(
	"github.com/opensds/go-panda/dataflow/pkg/scheduler/trigger"
	"github.com/robfig/cron"
	"github.com/micro/go-log"
	"fmt"
)


type CronTrigger struct {
	plans map[string] *cron.Cron
}

func init()  {
	ct := &CronTrigger{
		plans:make(map[string] *cron.Cron),
	}
	trigger.RegisterTrigger(trigger.TriggerTypeCron, ct)
}

func (c* CronTrigger) Add(planId, properties string, executer trigger.Executer) error {
	cn := cron.New()
	c.plans[planId] = cn
	if err := cn.AddFunc(properties, executer.Run); err != nil {
		log.Logf("Add plan(%s) in corn trigger failed", planId)
		return fmt.Errorf("Add plan(%s) in corn trigger failed", planId)
	}
	cn.Start()
	return nil
}

func (c* CronTrigger) Update(planId, properties string, executer trigger.Executer) error {
	c.Remove(planId)
	if err := c.Add(planId, properties, executer); err != nil {
		return err
	}
	return nil
}

func (c* CronTrigger)  Remove(planId string) error {
	cn,ok := c.plans[planId]
	if !ok {
		log.Logf("Specified plan(%s) is not found", planId)
		return nil
	}
	cn.Stop()
	delete(c.plans, planId)
	return nil
}
