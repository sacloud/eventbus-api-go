package eventbus_test

import (
	"context"
	"os"
	"strconv"
	"testing"

	client "github.com/sacloud/api-client-go"
	eventbus "github.com/sacloud/eventbus-api-go"
	v1 "github.com/sacloud/eventbus-api-go/apis/v1"
	"github.com/sacloud/packages-go/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessConfigurationAPI(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN",
		"SAKURACLOUD_ACCESS_TOKEN_SECRET", "SAKURACLOUD_SIMPLE_NOTIFICATION_GROUP_ID")(t)

	client, err := eventbus.NewClient(client.WithDisableProfile(true))
	require.NoError(t, err)

	ctx := context.Background()
	pcOp := eventbus.NewProcessConfigurationOp(client)
	groupId := os.Getenv("SAKURACLOUD_SIMPLE_NOTIFICATION_GROUP_ID")

	resCreate, err := pcOp.Create(ctx, v1.ProcessConfigurationRequestSettings{
		Name: "SDK Test", Description: "SDK Testの概要",
		Settings: eventbus.CreateSimpleNotificationSettings(groupId, "メッセージ"),
	})
	require.NoError(t, err)
	pcId := strconv.FormatInt(resCreate.ID, 10)

	resList, err := pcOp.List(ctx)
	assert.NoError(t, err)
	found := false
	for _, pc := range resList {
		if pc.ID == resCreate.ID {
			found = true
			assert.Equal(t, "SDK Test", pc.Name)
		}
	}
	assert.True(t, found, "Created ProcessConfiguration not found in list")

	_, err = pcOp.Update(ctx, pcId, v1.ProcessConfigurationRequestSettings{
		Name: "SDK Test 2", Description: "SDK Test 2の概要",
		Settings: eventbus.CreateSimpleNotificationSettings(groupId, "メッセージ2"),
	})
	assert.NoError(t, err)

	resRead, err := pcOp.Read(ctx, pcId)
	assert.NoError(t, err)
	assert.Equal(t, "SDK Test 2", resRead.Name)

	err = pcOp.UpdateSecret(ctx, pcId, v1.ProcessConfigurationSecret{
		AccessToken:       os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
	})
	assert.NoError(t, err)

	err = pcOp.Delete(ctx, pcId)
	require.NoError(t, err)
}
