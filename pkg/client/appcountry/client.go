//nolint:dupl
package appcountry

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get appcountry connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateCountry(ctx context.Context, in *appcountrymgrpb.CountryReq) (*npool.Country, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCountry(ctx, &npool.CreateCountryRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create appcountry: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create appcountry: %v", err)
	}
	return info.(*npool.Country), nil
}

func CreateCountries(ctx context.Context, in []*appcountrymgrpb.CountryReq) ([]*npool.Country, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCountries(ctx, &npool.CreateCountriesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create appcountries: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create appcountries: %v", err)
	}
	return infos.([]*npool.Country), nil
}

func GetCountries(ctx context.Context, conds *appcountrymgrpb.Conds, offset, limit int32) ([]*npool.Country, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCountries(ctx, &npool.GetCountriesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcountries: %v", err)
		}
		total = resp.Total
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get appcountries: %v", err)
	}
	return infos.([]*npool.Country), total, nil
}

func GetCountryOnly(ctx context.Context, conds *appcountrymgrpb.Conds) (*npool.Country, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCountryOnly(ctx, &npool.GetCountryOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get appcountryonly: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get appcountryonly: %v", err)
	}
	return info.(*npool.Country), nil
}

func DeleteCountry(ctx context.Context, id string) (*npool.Country, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteCountry(ctx, &npool.DeleteCountryRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete appcountry: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete appcountry: %v", err)
	}
	return info.(*npool.Country), nil
}
