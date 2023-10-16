//nolint:nolintlint,dupl
package message

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	servicename "github.com/NpoolPlatform/g11n-middleware/pkg/servicename"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func CreateMessage(ctx context.Context, req *npool.MessageReq) (*npool.Message, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateMessage(ctx, &npool.CreateMessageRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Message), nil
}

func CreateMessages(ctx context.Context, reqs []*npool.MessageReq) ([]*npool.Message, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateMessages(ctx, &npool.CreateMessagesRequest{
			Infos: reqs,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Message), nil
}

func UpdateMessage(ctx context.Context, req *npool.MessageReq) (*npool.Message, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateMessage(ctx, &npool.UpdateMessageRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Message), nil
}

func GetMessage(ctx context.Context, id string) (*npool.Message, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetMessage(ctx, &npool.GetMessageRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Message), nil
}

func GetMessages(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Message, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetMessages(ctx, &npool.GetMessagesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Message), total, nil
}

func GetMessageOnly(ctx context.Context, conds *npool.Conds) (*npool.Message, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetMessages(ctx, &npool.GetMessagesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2, //nolint
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Message)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Message)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.Message)[0], nil
}

func DeleteMessage(ctx context.Context, req *npool.MessageReq) (*npool.Message, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteMessage(ctx, &npool.DeleteMessageRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Message), nil
}

func ExistMessageConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistMessageConds(ctx, &npool.ExistMessageCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get message: %v", err)
	}
	return infos.(bool), nil
}
