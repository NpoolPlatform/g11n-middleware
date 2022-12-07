package message1

import (
	"context"

	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	messagemgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/message"
)

func UpdateMessage(ctx context.Context, in *messagemgrpb.MessageReq) (*npool.Message, error) {
	info, err := messagemgrcli.UpdateMessage(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetMessage(ctx, info.ID)
}
