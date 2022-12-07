package appcountry

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	appcountrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/appcountry"
)

func DeleteCountry(ctx context.Context, id string) (*npool.Country, error) {
	info, err := GetCountry(ctx, id)
	if err != nil {
		return nil, err
	}

	_, err = appcountrymgrcli.DeleteCountry(ctx, id)
	if err != nil {
		return nil, err
	}

	return info, nil
}
