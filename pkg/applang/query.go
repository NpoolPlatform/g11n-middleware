package applang

import (
	"context"
	"fmt"

	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	"github.com/NpoolPlatform/g11n-manager/pkg/db"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"

	entapplang "github.com/NpoolPlatform/g11n-manager/pkg/db/ent/applang"
	entlang "github.com/NpoolPlatform/g11n-manager/pkg/db/ent/lang"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

func GetLang(ctx context.Context, id string) (*npool.Lang, error) {
	infos := []*npool.Lang{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppLang.
			Query().
			Where(
				entapplang.ID(uuid.MustParse(id)),
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

func GetLangs(ctx context.Context, conds *applangmgrpb.Conds, offset, limit int32) ([]*npool.Lang, uint32, error) {
	infos := []*npool.Lang{}
	total := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppLang.
			Query()

		if conds.ID != nil {
			stm.
				Where(
					entapplang.ID(uuid.MustParse(conds.GetID().GetValue())),
				)
		}
		if conds.AppID != nil {
			stm.
				Where(
					entapplang.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
		}
		if conds.LangID != nil {
			stm.
				Where(
					entapplang.LangID(uuid.MustParse(conds.GetLangID().GetValue())),
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

func GetManyLangs(ctx context.Context, ids []string) ([]*npool.Lang, error) {
	infos := []*npool.Lang{}

	langIDs := []uuid.UUID{}
	for _, id := range ids {
		langIDs = append(langIDs, uuid.MustParse(id))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppLang.
			Query().
			Where(
				entapplang.IDIn(langIDs...),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func join(stm *ent.AppLangQuery) *ent.AppLangSelect {
	return stm.
		Select(
			entapplang.FieldID,
			entapplang.FieldAppID,
			entapplang.FieldLangID,
			entapplang.FieldMain,
			entapplang.FieldCreatedAt,
			entapplang.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entlang.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entapplang.FieldLangID),
					t1.C(entlang.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entlang.FieldLang), "lang"),
					sql.As(t1.C(entlang.FieldLogo), "logo"),
					sql.As(t1.C(entlang.FieldName), "name"),
					sql.As(t1.C(entlang.FieldShort), "short"),
				)
		})
}
