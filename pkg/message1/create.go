package message1

import (
	"context"

	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	messagemgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/message"
)

func CreateMessage(ctx context.Context, in *messagemgrpb.MessageReq) (*npool.Message, error) {
	info, err := messagemgrcli.CreateMessage(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetMessage(ctx, info.ID)
}

func CreateMessages(ctx context.Context, in []*messagemgrpb.MessageReq) ([]*npool.Message, error) {
	infos, err := messagemgrcli.CreateMessages(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, info := range infos {
		ids = append(ids, info.ID)
	}

	return GetManyMessages(ctx, ids)
}
