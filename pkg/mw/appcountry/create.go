package appcountry

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	appcountrycrud "github.com/NpoolPlatform/g11n-middleware/pkg/crud/appcountry"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	"github.com/google/uuid"
)

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}
	if h.CountryID == nil {
		return nil, fmt.Errorf("invalid countryid")
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appcountrycrud.SetQueryConds(
			cli.AppCountry.Query(),
			&appcountrycrud.Conds{
				AppID:     &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
				CountryID: &cruder.Cond{Op: cruder.EQ, Val: *h.CountryID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			h.ID = &info.ID
			return nil
		}
		if _, err := appcountrycrud.CreateSet(
			cli.AppCountry.Create(),
			&appcountrycrud.Req{
				ID:        h.ID,
				AppID:     &h.AppID,
				CountryID: h.CountryID,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCountry(ctx)
}

func (h *Handler) CreateCountries(ctx context.Context) ([]*npool.Country, error) {
	ids := []uuid.UUID{}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = *req.ID
			}
			if _, err := appcountrycrud.CreateSet(
				cli.AppCountry.Create(),
				&appcountrycrud.Req{
					ID:        &id,
					AppID:     req.AppID,
					CountryID: req.CountryID,
				},
			).Save(ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &appcountrycrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	infos, _, err := h.GetCountries(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
