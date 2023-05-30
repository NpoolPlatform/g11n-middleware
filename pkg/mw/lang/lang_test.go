package lang

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/g11n-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	ret = npool.Lang{
		ID:    uuid.NewString(),
		Lang:  uuid.NewString(),
		Logo:  uuid.NewString(),
		Name:  uuid.NewString(),
		Short: uuid.NewString(),
	}
)

func creatLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithLang(&ret.Lang),
		WithLogo(&ret.Logo),
		WithName(&ret.Name),
		WithShort(&ret.Short),
	)
	assert.Nil(t, err)

	info, err := handler.CreateLang(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateLang(t *testing.T) {
	ret.Lang = uuid.NewString()
	ret.Logo = uuid.NewString()
	ret.Name = uuid.NewString()
	ret.Short = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithLang(&ret.Lang),
		WithLogo(&ret.Logo),
		WithName(&ret.Name),
		WithShort(&ret.Short),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateLang(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetLang(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLangs(t *testing.T) {
	conds := &npool.Conds{
		ID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetLangs(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteLang(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetLang(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestLang(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("creatLang", creatLang)
	t.Run("updateLang", updateLang)
	t.Run("getLang", getLang)
	t.Run("getLangs", getLangs)
	t.Run("deleteLang", deleteLang)
}
