package appcountry

import (
	"context"
	"fmt"

	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	"github.com/NpoolPlatform/g11n-manager/pkg/db"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"

	entappcountry "github.com/NpoolPlatform/g11n-manager/pkg/db/ent/appcountry"
	entcountry "github.com/NpoolPlatform/g11n-manager/pkg/db/ent/country"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

func GetCountry(ctx context.Context, id string) (*npool.Country, error) {
	infos := []*npool.Country{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppCountry.
			Query().
			Where(
				entappcountry.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func GetCountries(ctx context.Context, conds *appcountrymgrpb.Conds, offset, limit int32) ([]*npool.Country, uint32, error) {
	infos := []*npool.Country{}
	total := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppCountry.
			Query()

		if conds.ID != nil {
			stm.
				Where(
					entappcountry.ID(uuid.MustParse(conds.GetID().GetValue())),
				)
		}
		if conds.AppID != nil {
			stm.
				Where(
					entappcountry.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
		}
		if conds.CountryID != nil {
			stm.
				Where(
					entappcountry.CountryID(uuid.MustParse(conds.GetCountryID().GetValue())),
				)
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func GetManyCountries(ctx context.Context, ids []string) ([]*npool.Country, error) {
	infos := []*npool.Country{}

	countryIDs := []uuid.UUID{}
	for _, id := range ids {
		countryIDs = append(countryIDs, uuid.MustParse(id))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppCountry.
			Query().
			Where(
				entappcountry.IDIn(countryIDs...),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func join(stm *ent.AppCountryQuery) *ent.AppCountrySelect {
	return stm.
		Select(
			entappcountry.FieldID,
			entappcountry.FieldAppID,
			entappcountry.FieldCountryID,
			entappcountry.FieldCreatedAt,
			entappcountry.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entcountry.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entappcountry.FieldCountryID),
					t1.C(entcountry.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entcountry.FieldCountry), "country"),
					sql.As(t1.C(entcountry.FieldFlag), "flag"),
					sql.As(t1.C(entcountry.FieldCode), "code"),
					sql.As(t1.C(entcountry.FieldShort), "short"),
				)
		})
}
