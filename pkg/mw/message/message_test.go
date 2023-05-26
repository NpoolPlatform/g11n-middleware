package message

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/g11n-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.Message{
		ID:        uuid.NewString(),
		AppID:     uuid.NewString(),
		LangID:    uuid.NewString(),
		MessageID: uuid.NewString(),
		Message:   uuid.NewString(),
		Disabled:  false,
		GetIndex:  0,
	}
)

func setupMessage(t *testing.T) func(*testing.T) {
	ah, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	message1, err := ah.CreateMessage(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, message1)

	return func(t *testing.T) {
		_, _ = ah.DeleteMessage(context.Background())
	}
}

func createMessage(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.AppID),
		WithLangID(&ret.LangID),
		WithMessageID(&ret.MessageID),
		WithMessage(&ret.Message),
		WithGetIndex(ret.GetIndex),
		WithDisabled(ret.Disabled),
	)
	assert.Nil(t, err)

	info, err := handler.CreateMessage(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateMessage(t *testing.T) {
	ret.AppID = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.AppID),
		WithLangID(&ret.LangID),
		WithMessageID(&ret.MessageID),
		WithMessage(&ret.Message),
		WithGetIndex(ret.GetIndex),
		WithDisabled(ret.Disabled),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateMessage(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getMessage(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetMessage(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getMessages(t *testing.T) {
	conds := &npool.Conds{
		// AppID: &message.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetMessages(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteMessage(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteMessage(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetMessage(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMessage(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupMessage(t)
	defer teardown(t)

	t.Run("createMessage", createMessage)
	t.Run("updateMessage", updateMessage)
	t.Run("getMessage", getMessage)
	t.Run("getMessages", getMessages)
	t.Run("deleteMessage", deleteMessage)
}
