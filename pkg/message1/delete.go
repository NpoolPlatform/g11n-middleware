package message1

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	messagemgrcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
)

func DeleteMessage(ctx context.Context, id string) (*npool.Message, error) {
	info, err := GetMessage(ctx, id)
	if err != nil {
		return nil, err
	}

	_, err = messagemgrcli.DeleteMessage(ctx, id)
	if err != nil {
		return nil, err
	}

	return info, nil
}
