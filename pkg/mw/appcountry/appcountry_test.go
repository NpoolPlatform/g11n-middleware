package appcountry

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	country "github.com/NpoolPlatform/g11n-middleware/pkg/mw/country"
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
	ret = npool.Country{
		ID:        uuid.NewString(),
		AppID:     uuid.NewString(),
		CountryID: uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	ch, err := country.NewHandler(
		context.Background(),
		country.WithID(&ret.CountryID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ch)
	country1, err := ch.CreateCountry(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, country1)

	ah, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCountryID(&ret.CountryID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	appcountry1, err := ah.CreateCountry(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, appcountry1)

	return func(*testing.T) {
		_, _ = ah.DeleteCountry(context.Background())
		_, _ = ch.DeleteCountry(context.Background())
	}
}

func creatCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.AppID),
		WithCountryID(&ret.CountryID),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCountry(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetCountry(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCountries(t *testing.T) {
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

	infos, _, err := handler.GetCountries(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCountry(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCountry(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCountry(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("creatCountry", creatCountry)
	t.Run("getCountry", getCountry)
	t.Run("getCountries", getCountries)
	t.Run("deleteCountry", deleteCountry)
}
