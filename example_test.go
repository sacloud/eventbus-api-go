// Copyright 2025- The sacloud/eventbus-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventbus_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	eventbus "github.com/sacloud/eventbus-api-go"
	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
)

var requriedEnvs = []string{
	"SAKURA_SIMPLE_NOTIFICATION_GROUP_ID",
	"SAKURACLOUD_ACCESS_TOKEN",
	"SAKURACLOUD_ACCESS_TOKEN_SECRET",
}

func checkEnvs() {
	for _, env := range requriedEnvs {
		if os.Getenv(env) == "" {
			panic(env + " is not set")
		}
	}
}

func ExampleProcessConfigurationAPI() {
	checkEnvs()

	client, err := eventbus.InitClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	pcOp := eventbus.NewProcessConfigurationOp(client)
	groupId := os.Getenv("SAKURA_SIMPLE_NOTIFICATION_GROUP_ID")

	resCreate, err := pcOp.Create(ctx, v1.ProcessConfigurationRequestSettings{
		Name: "SDK Test", Description: "SDK Testの概要",
		Settings: eventbus.CreateSimpleNotificationSettings(groupId, "メッセージ"),
	})
	if err != nil {
		panic(err)
	}
	pcId := strconv.FormatInt(resCreate.ID, 10)

	resList, err := pcOp.List(ctx)
	if err != nil {
		panic(err)
	}
	for _, pc := range resList {
		if pc.ID == resCreate.ID {
			fmt.Println(pc.Name)
		}
	}

	_, err = pcOp.Update(ctx, pcId, v1.ProcessConfigurationRequestSettings{
		Name: "SDK Test 2", Description: "SDK Test 2の概要",
		Settings: eventbus.CreateSimpleNotificationSettings(groupId, "メッセージ2"),
	})
	if err != nil {
		panic(err)
	}

	resRead, err := pcOp.Read(ctx, pcId)
	if err != nil {
		panic(err)
	}
	fmt.Println(resRead.Name)

	err = pcOp.UpdateSecret(ctx, pcId, v1.ProcessConfigurationSecret{
		AccessToken:       os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
	})
	if err != nil {
		panic(err)
	}

	resDelete, err := pcOp.Delete(ctx, pcId)
	if err != nil {
		panic(err)
	}
	fmt.Println(resDelete.Name)
	// Output:
	// SDK Test
	// SDK Test 2
	// SDK Test 2
}

func ExampleScheduleAPI() {
	checkEnvs()

	client, err := eventbus.InitClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	pcOp := eventbus.NewProcessConfigurationOp(client)
	schedOp := eventbus.NewScheduleOp(client)
	groupId := os.Getenv("SAKURA_SIMPLE_NOTIFICATION_GROUP_ID")

	// テスト用の実行設定の生成
	pc, err := pcOp.Create(ctx, v1.ProcessConfigurationRequestSettings{
		Name: "SDK Test", Description: "SDK Testの概要",
		Settings: eventbus.CreateSimpleNotificationSettings(groupId, "メッセージ"),
	})
	if err != nil {
		panic(err)
	}
	pcId := strconv.FormatInt(pc.ID, 10)

	resCreate, err := schedOp.Create(ctx, v1.ScheduleRequestSettings{
		Name: "SDK Test", Description: "SDK Testの概要",
		Settings: v1.ScheduleSettings{
			ProcessConfigurationID: pcId,
			RecurringStep:          5,
			RecurringUnit:          "min",
			StartsAt:               time.Now().UnixMilli(),
		},
	})
	if err != nil {
		panic(err)
	}
	schedId := strconv.FormatInt(resCreate.ID, 10)

	resList, err := schedOp.List(ctx)
	if err != nil {
		panic(err)
	}
	for _, sched := range resList {
		if sched.ID == resCreate.ID {
			fmt.Println(sched.Name)
			fmt.Println(sched.Settings.RecurringUnit)
		}
	}

	_, err = schedOp.Update(ctx, schedId, v1.ScheduleRequestSettings{
		Name: "SDK Test 2", Description: "SDK Test 2の概要",
		Settings: v1.ScheduleSettings{
			ProcessConfigurationID: pcId,
			RecurringStep:          1,
			RecurringUnit:          "hour",
			StartsAt:               time.Now().UnixMilli(),
		},
	})
	if err != nil {
		panic(err)
	}

	resRead, err := schedOp.Read(ctx, schedId)
	if err != nil {
		panic(err)
	}
	fmt.Println(resRead.Name)
	fmt.Println(resRead.Settings.RecurringUnit)

	resDelete, err := schedOp.Delete(ctx, schedId)
	if err != nil {
		panic(err)
	}

	_, err = pcOp.Delete(ctx, pcId)
	if err != nil {
		panic(err)
	}

	fmt.Println(resDelete.Name)
	// Output:
	// SDK Test
	// min
	// SDK Test 2
	// hour
	// SDK Test 2
}
