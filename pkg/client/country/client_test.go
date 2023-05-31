package country

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/g11n-middleware/pkg/testinit"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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

func createCountry(t *testing.T) {
	req := npool.CountryReq{
		ID:      &ret.ID,
		Country: &ret.Country,
		Flag:    &ret.Flag,
		Code:    &ret.Code,
		Short:   &ret.Short,
	}
	info, err := CreateCountry(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCountry(t *testing.T) {
	ret.Country = uuid.NewString()
	ret.Short = uuid.NewString()
	ret.Flag = uuid.NewString()
	ret.Code = uuid.NewString()
	req := npool.CountryReq{
		ID:      &ret.ID,
		Country: &ret.Country,
		Flag:    &ret.Flag,
		Code:    &ret.Code,
		Short:   &ret.Short,
	}
	info, err := UpdateCountry(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCountry(t *testing.T) {
	info, err := GetCountry(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCountries(t *testing.T) {
	_, total, err := GetCountries(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func deleteCountry(t *testing.T) {
	info, err := DeleteCountry(context.Background(), &npool.CountryReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetCountry(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCountry", createCountry)
	t.Run("updateCountry", updateCountry)
	t.Run("getCountry", getCountry)
	t.Run("getCountries", getCountries)
	t.Run("deleteCountry", deleteCountry)
}
