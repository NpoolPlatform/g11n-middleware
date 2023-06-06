package applang

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	lang "github.com/NpoolPlatform/g11n-middleware/pkg/mw/lang"
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
		ID:     uuid.NewString(),
		AppID:  uuid.NewString(),
		LangID: uuid.NewString(),
		Lang:   "test mw lang" + uuid.NewString(),
		Name:   "test mw lang" + uuid.NewString(),
		Logo:   "test mw logo" + uuid.NewString(),
		Short:  "test mw short" + uuid.NewString(),
		Main:   true,
	}
)

func setup(t *testing.T) func(*testing.T) {
	lh, err := lang.NewHandler(
		context.Background(),
		lang.WithID(&ret.LangID),
		lang.WithLang(&ret.Lang),
		lang.WithName(&ret.Name),
		lang.WithLogo(&ret.Logo),
		lang.WithShort(&ret.Short),
	)
	assert.Nil(t, err)
	assert.NotNil(t, lh)
	lang1, err := lh.CreateLang(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, lang1)

	return func(*testing.T) {
		_, _ = lh.DeleteLang(context.Background())
	}
}

func creatLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.AppID),
		WithLangID(&ret.LangID),
		WithMain(&ret.Main),
	)
	assert.Nil(t, err)

	info, err := handler.CreateLang(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateLang(t *testing.T) {
	ret.Main = false
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.AppID),
		WithLangID(&ret.LangID),
		WithMain(&ret.Main),
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
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
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

	teardown := setup(t)
	defer teardown(t)

	t.Run("creatLang", creatLang)
	t.Run("updateLang", updateLang)
	t.Run("getLang", getLang)
	t.Run("getLangs", getLangs)
	t.Run("deleteLang", deleteLang)
}
