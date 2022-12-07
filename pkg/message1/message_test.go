package message1

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/g11n-middleware/pkg/testinit"

	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	langmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

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
	langLang  = uuid.NewString()
	langLogo  = uuid.NewString()
	langName  = uuid.NewString()
	langShort = uuid.NewString()
)

var langReq = &langmgrpb.LangReq{
	Lang:  &langLang,
	Logo:  &langLogo,
	Name:  &langName,
	Short: &langShort,
}

var ret = &npool.Message{
	AppID:     uuid.NewString(),
	Lang:      langLang,
	MessageID: uuid.NewString(),
	Message:   uuid.NewString(),
	GetIndex:  1,
}

var req = &messagemgrpb.MessageReq{
	AppID:     &ret.AppID,
	MessageID: &ret.MessageID,
	Message:   &ret.Message,
	GetIndex:  &ret.GetIndex,
}

func create(t *testing.T) {
	lang, err := langmgrcli.CreateLang(context.Background(), langReq)
	assert.Nil(t, err)

	req.LangID = &lang.ID
	ret.LangID = lang.ID

	info, err := CreateMessage(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		req.ID = &info.ID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	disabled := true

	req.Disabled = &disabled
	ret.Disabled = disabled

	info, err := UpdateMessage(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func delete1(t *testing.T) {
	info, err := DeleteMessage(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	_, err = GetMessage(context.Background(), ret.ID)
	assert.NotNil(t, err)
}

func TestLang(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("update", update)
	t.Run("delete", delete1)
}
