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

func (h *createHandler) validate(req *countrycrud.Req) error {
	if req.Country == nil || *req.Country == "" {
		return fmt.Errorf("invalid country")
	}
	if req.Flag == nil || *req.Flag == "" {
		return fmt.Errorf("invalid flag")
	}
	if req.Code == nil || *req.Code == "" {
		return fmt.Errorf("invalid code")
	}
	if req.Short == nil || *req.Short == "" {
		return fmt.Errorf("invalid short")
	}
	return nil
}

func (h *createHandler) createCountry(ctx context.Context, cli *ent.Client, req *countrycrud.Req) (*npool.Country, error) {
	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixCreateCountry,
		*req.Country,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &countrycrud.Conds{
		Country: &cruder.Cond{Op: cruder.EQ, Val: *req.Country},
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
	if req.EntID == nil {
		req.EntID = &id
	}

	info, err := countrycrud.CreateSet(
		cli.Country.Create(),
		&countrycrud.Req{
			EntID:   req.EntID,
			Country: req.Country,
			Flag:    req.Flag,
			Code:    req.Code,
			Short:   req.Short,
		},
	).Save(ctx)
	if err != nil {
		return nil, err
	}

	h.ID = &info.ID
	h.EntID = &info.EntID

	return nil, nil
}

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	handler := &createHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		req := &countrycrud.Req{
			EntID:   h.EntID,
			Country: h.Country,
			Flag:    h.Flag,
			Code:    h.Code,
			Short:   h.Short,
		}
		if err := handler.validate(req); err != nil {
			return err
		}
		h.Conds = &countrycrud.Conds{
			Country: &cruder.Cond{Op: cruder.EQ, Val: *h.Country},
		}
		exist, err := h.ExistCountryConds(ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("country exist")
		}
		info, err := handler.createCountry(ctx, cli, req)
		if err != nil {
			return err
		}
		if info != nil {
			id, err := uuid.Parse(info.GetEntID())
			if err != nil {
				return err
			}
			h.EntID = &id
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
			if err := handler.validate(req); err != nil {
				return err
			}
			if req.EntID != nil {
				handler.EntID = req.EntID
			}
			info, err := handler.createCountry(ctx, cli, req)
			if err != nil {
				return err
			}
			if info != nil {
				id, err := uuid.Parse(info.GetEntID())
				if err != nil {
					return err
				}
				h.EntID = &id
			}
			ids = append(ids, *h.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &countrycrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCountries(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
