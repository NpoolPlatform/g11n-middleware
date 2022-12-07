package applang

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/g11n-middleware/pkg/testinit"

	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

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

var ret = &npool.Lang{
	AppID: uuid.NewString(),
	Lang:  langLang,
	Logo:  langLogo,
	Name:  langName,
	Short: langShort,
}

var req = &applangmgrpb.LangReq{
	AppID: &ret.AppID,
}

func create(t *testing.T) {
	lang, err := langmgrcli.CreateLang(context.Background(), langReq)
	assert.Nil(t, err)

	req.LangID = &lang.ID
	ret.LangID = lang.ID

	info, err := CreateLang(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, ret)
	}
}

func TestLang(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
}
