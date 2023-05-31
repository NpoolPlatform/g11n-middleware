package country

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
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
	ret = npool.Country{
		ID:      uuid.NewString(),
		Country: uuid.NewString(),
		Flag:    uuid.NewString(),
		Code:    uuid.NewString(),
		Short:   uuid.NewString(),
	}
)

func creatCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCountry(&ret.Country),
		WithFlag(&ret.Flag),
		WithCode(&ret.Code),
		WithShort(&ret.Short),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCountry(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCountry(t *testing.T) {
	ret.Country = uuid.NewString()
	ret.Flag = uuid.NewString()
	ret.Code = uuid.NewString()
	ret.Short = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCountry(&ret.Country),
		WithFlag(&ret.Flag),
		WithCode(&ret.Code),
		WithShort(&ret.Short),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCountry(context.Background())
	if assert.Nil(t, err) {
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

func getCountrys(t *testing.T) {
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

	t.Run("creatCountry", creatCountry)
	t.Run("updateCountry", updateCountry)
	t.Run("getCountry", getCountry)
	t.Run("getCountrys", getCountrys)
	t.Run("deleteCountry", deleteCountry)
}
