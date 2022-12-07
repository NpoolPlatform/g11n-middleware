package appcountry

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/g11n-middleware/pkg/testinit"

	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	countrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/country"
	countrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

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
	countryCountry = uuid.NewString()
	countryFlag    = uuid.NewString()
	countryCode    = uuid.NewString()
	countryShort   = uuid.NewString()
)

var countryReq = &countrymgrpb.CountryReq{
	Country: &countryCountry,
	Flag:    &countryFlag,
	Code:    &countryCode,
	Short:   &countryShort,
}

var ret = &npool.Country{
	AppID:   uuid.NewString(),
	Country: countryCountry,
	Flag:    countryFlag,
	Code:    countryCode,
	Short:   countryShort,
}

var req = &appcountrymgrpb.CountryReq{
	AppID: &ret.AppID,
}

func create(t *testing.T) {
	country, err := countrymgrcli.CreateCountry(context.Background(), countryReq)
	assert.Nil(t, err)

	req.CountryID = &country.ID
	ret.CountryID = country.ID

	info, err := CreateCountry(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, ret)
	}
}

func TestCountry(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
}
