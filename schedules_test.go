package eventbus_test

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	eventbus "github.com/sacloud/eventbus-api-go"
	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
	"github.com/sacloud/packages-go/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScheduleAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN",
		"SAKURACLOUD_ACCESS_TOKEN_SECRET", "SAKURACLOUD_SIMPLE_NOTIFICATION_GROUP_ID")(t)

	client, err := eventbus.NewClient()
	require.NoError(t, err)

	ctx := context.Background()
	pcOp := eventbus.NewProcessConfigurationOp(client)
	schedOp := eventbus.NewScheduleOp(client)
	groupId := os.Getenv("SAKURACLOUD_SIMPLE_NOTIFICATION_GROUP_ID")

	pc, err := pcOp.Create(ctx, v1.ProcessConfigurationRequestSettings{
		Name: "SDK Test", Description: "SDK Testの概要", Tags: []string{"tag1", "tag2"},
		Settings: eventbus.CreateSimpleNotificationSettings(groupId, "メッセージ"),
	})
	require.NoError(t, err)
	pcId := strconv.FormatInt(pc.ID, 10)
	defer func() {
		_ = pcOp.Delete(ctx, pcId)
	}()

	resCreate, err := schedOp.Create(ctx, v1.ScheduleRequestSettings{
		Name: "SDK Test", Description: "SDK Testの概要",
		Settings: v1.ScheduleSettings{
			ProcessConfigurationID: pcId,
			RecurringStep:          5,
			RecurringUnit:          "min",
			StartsAt:               time.Now().UnixMilli(),
		},
	})
	require.NoError(t, err)
	schedId := strconv.FormatInt(resCreate.ID, 10)

	resList, err := schedOp.List(ctx)
	assert.NoError(t, err)
	found := false
	for _, sched := range resList {
		if sched.ID == resCreate.ID {
			found = true
			assert.Equal(t, "SDK Test", sched.Name)
			assert.Equal(t, v1.CreateScheduleRequestRecurringUnit("min"), sched.Settings.RecurringUnit)
		}
	}
	assert.True(t, found, "Created Schedule not found in list")

	_, err = schedOp.Update(ctx, schedId, v1.ScheduleRequestSettings{
		Name: "SDK Test 2", Description: "SDK Test 2の概要", Tags: []string{"tag1", "tag2"},
		Settings: v1.ScheduleSettings{
			ProcessConfigurationID: pcId,
			RecurringStep:          1,
			RecurringUnit:          "hour",
			StartsAt:               time.Now().UnixMilli(),
		},
	})
	assert.NoError(t, err)

	resRead, err := schedOp.Read(ctx, schedId)
	assert.NoError(t, err)
	assert.Equal(t, "SDK Test 2", resRead.Name)
	assert.Equal(t, v1.CreateScheduleRequestRecurringUnit("hour"), resRead.Settings.RecurringUnit)
	assert.Equal(t, []string{"tag1", "tag2"}, resRead.Tags)

	err = schedOp.Delete(ctx, schedId)
	require.NoError(t, err)
}
