package appcountry

import (
	"context"

	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	appcountrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/appcountry"
)

func CreateCountry(ctx context.Context, in *appcountrymgrpb.CountryReq) (*npool.Country, error) {
	info, err := appcountrymgrcli.CreateCountry(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetCountry(ctx, info.ID)
}

func CreateCountries(ctx context.Context, in []*appcountrymgrpb.CountryReq) ([]*npool.Country, error) {
	infos, err := appcountrymgrcli.CreateCountries(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, info := range infos {
		ids = append(ids, info.ID)
	}

	return GetManyCountries(ctx, ids)
}
