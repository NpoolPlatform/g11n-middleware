package lang

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
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
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
	ret = npool.Lang{
		ID:    uuid.NewString(),
		Lang:  uuid.NewString(),
		Logo:  uuid.NewString(),
		Name:  uuid.NewString(),
		Short: uuid.NewString(),
	}
)

func createLang(t *testing.T) {
	req := npool.LangReq{
		ID:    &ret.ID,
		Lang:  &ret.Lang,
		Logo:  &ret.Logo,
		Name:  &ret.Name,
		Short: &ret.Short,
	}
	info, err := CreateLang(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateLang(t *testing.T) {
	ret.Lang = uuid.NewString()
	ret.Short = uuid.NewString()
	ret.Name = uuid.NewString()
	ret.Logo = uuid.NewString()
	req := npool.LangReq{
		ID:    &ret.ID,
		Lang:  &ret.Lang,
		Logo:  &ret.Logo,
		Name:  &ret.Name,
		Short: &ret.Short,
	}
	info, err := UpdateLang(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLang(t *testing.T) {
	info, err := GetLang(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLangs(t *testing.T) {
	_, total, err := GetLangs(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func deleteLang(t *testing.T) {
	info, err := DeleteLang(context.Background(), &npool.LangReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetLang(context.Background(), ret.ID)
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

	t.Run("createLang", createLang)
	t.Run("updateLang", updateLang)
	t.Run("getLang", getLang)
	t.Run("getLangs", getLangs)
	t.Run("deleteLang", deleteLang)
}
