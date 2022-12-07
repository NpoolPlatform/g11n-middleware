//nolint:dupl
package applang

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	constant "github.com/NpoolPlatform/g11n-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get applang connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateLang(ctx context.Context, in *applangmgrpb.LangReq) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateLang(ctx, &npool.CreateLangRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create applang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create applang: %v", err)
	}
	return info.(*npool.Lang), nil
}

func CreateLangs(ctx context.Context, in []*applangmgrpb.LangReq) ([]*npool.Lang, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateLangs(ctx, &npool.CreateLangsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create applangs: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create applangs: %v", err)
	}
	return infos.([]*npool.Lang), nil
}

func UpdateLang(ctx context.Context, in *applangmgrpb.LangReq) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateLang(ctx, &npool.UpdateLangRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update applang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update applang: %v", err)
	}
	return info.(*npool.Lang), nil
}

func GetLangs(ctx context.Context, conds *applangmgrpb.Conds, limit, offset int32) ([]*npool.Lang, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetLangs(ctx, &npool.GetLangsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get applangs: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get applangs: %v", err)
	}
	return infos.([]*npool.Lang), total, nil
}

func DeleteLang(ctx context.Context, id string) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteLang(ctx, &npool.DeleteLangRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete applang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete applang: %v", err)
	}
	return info.(*npool.Lang), nil
}
