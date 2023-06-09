package country

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	countrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/country"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCountry(ctx context.Context, cli *ent.Client) (*npool.Country, error) {
	if h.Country == nil {
		return nil, fmt.Errorf("country invalid")
	}
	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixCreateCountry,
		*h.Country,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &countrycrud.Conds{
		Country: &cruder.Cond{Op: cruder.EQ, Val: *h.Country},
	}
	h.Limit = 2
	exist, _, err := h.GetCountries(ctx)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return exist[0], nil
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	info, err := countrycrud.CreateSet(
		cli.Country.Create(),
		&countrycrud.Req{
			ID:      h.ID,
			Country: h.Country,
			Flag:    h.Flag,
			Code:    h.Code,
			Short:   h.Short,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID

	return nil, nil
}

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := handler.createCountry(ctx, cli)
		if err != nil {
			return err
		}
		if info != nil {
			id, err := uuid.Parse(info.GetID())
			if err != nil {
				return err
			}
			h.ID = &id
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCountry(ctx)
}

func (h *Handler) CreateCountries(ctx context.Context) ([]*npool.Country, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			handler.ID = nil
			handler.Country = req.Country
			handler.Flag = req.Flag
			handler.Code = req.Code
			handler.Short = req.Short
			info, err := handler.createCountry(ctx, cli)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.GetID())
				if err != nil {
					return err
				}
				h.ID = &id
			}
			ids = append(ids, *h.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &countrycrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCountries(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
