package message

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/g11n-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	applang "github.com/NpoolPlatform/g11n-middleware/pkg/mw/applang"
	lang "github.com/NpoolPlatform/g11n-middleware/pkg/mw/lang"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
		Lang:      "test lang" + uuid.NewString(),
		LangID:    uuid.NewString(),
		MessageID: uuid.NewString(),
		Message:   "test message" + uuid.NewString(),
		Disabled:  false,
		GetIndex:  0,
	}
	langName    = "test lang name" + uuid.NewString()
	langLogo    = "test lang logo" + uuid.NewString()
	langShort   = "test lang short" + uuid.NewString()
	appLangID   = uuid.NewString()
	appLangMain = false
)

func setupMessage(t *testing.T) func(*testing.T) {
	lh, err := lang.NewHandler(
		context.Background(),
		lang.WithID(&ret.LangID),
		lang.WithLang(&ret.Lang),
		lang.WithName(&langName),
		lang.WithLogo(&langLogo),
		lang.WithShort(&langShort),
	)
	assert.Nil(t, err)
	assert.NotNil(t, lh)
	lang1, err := lh.CreateLang(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, lang1)

	ah, err := applang.NewHandler(
		context.Background(),
		applang.WithID(&appLangID),
		applang.WithAppID(&ret.AppID),
		applang.WithLangID(&ret.LangID),
		applang.WithMain(&appLangMain),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	applang1, err := ah.CreateLang(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, applang1)

	return func(t *testing.T) {
		_, _ = lh.DeleteLang(context.Background())
		_, _ = ah.DeleteLang(context.Background())
	}
}

func createMessage(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithLangID(&ret.LangID),
		WithMessageID(&ret.MessageID),
		WithMessage(&ret.Message),
		WithGetIndex(&ret.GetIndex),
		WithDisabled(&ret.Disabled),
	)
	assert.Nil(t, err)

	info, err := handler.CreateMessage(context.Background())
	if assert.Nil(t, err) {
		ret.Lang = info.Lang
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateMessage(t *testing.T) {
	ret.Message = "change message" + uuid.NewString()
	ret.MessageID = "change messageID " + uuid.NewString()
	ret.GetIndex = 8
	ret.Disabled = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(&ret.AppID),
		WithLangID(&ret.LangID),
		WithMessageID(&ret.MessageID),
		WithMessage(&ret.Message),
		WithGetIndex(&ret.GetIndex),
		WithDisabled(&ret.Disabled),
	)
	assert.Nil(t, err)
	info, err := handler.UpdateMessage(context.Background())
	if assert.Nil(t, err) {
		ret.AppID = info.AppID
		ret.Message = info.Message
		ret.MessageID = info.MessageID
		ret.GetIndex = info.GetIndex
		ret.Disabled = info.Disabled
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
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
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
		WithAppID(&ret.AppID),
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
