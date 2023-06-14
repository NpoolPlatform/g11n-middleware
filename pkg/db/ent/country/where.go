// Code generated by ent, DO NOT EDIT.

package country

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountry), v))
	})
}

// Flag applies equality check predicate on the "flag" field. It's identical to FlagEQ.
func Flag(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlag), v))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Short applies equality check predicate on the "short" field. It's identical to ShortEQ.
func Short(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShort), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountry), v))
	})
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCountry), v))
	})
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCountry), v...))
	})
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCountry), v...))
	})
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCountry), v))
	})
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCountry), v))
	})
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCountry), v))
	})
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCountry), v))
	})
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCountry), v))
	})
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCountry), v))
	})
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCountry), v))
	})
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCountry)))
	})
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCountry)))
	})
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCountry), v))
	})
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCountry), v))
	})
}

// FlagEQ applies the EQ predicate on the "flag" field.
func FlagEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlag), v))
	})
}

// FlagNEQ applies the NEQ predicate on the "flag" field.
func FlagNEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlag), v))
	})
}

// FlagIn applies the In predicate on the "flag" field.
func FlagIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFlag), v...))
	})
}

// FlagNotIn applies the NotIn predicate on the "flag" field.
func FlagNotIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFlag), v...))
	})
}

// FlagGT applies the GT predicate on the "flag" field.
func FlagGT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFlag), v))
	})
}

// FlagGTE applies the GTE predicate on the "flag" field.
func FlagGTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFlag), v))
	})
}

// FlagLT applies the LT predicate on the "flag" field.
func FlagLT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFlag), v))
	})
}

// FlagLTE applies the LTE predicate on the "flag" field.
func FlagLTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFlag), v))
	})
}

// FlagContains applies the Contains predicate on the "flag" field.
func FlagContains(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFlag), v))
	})
}

// FlagHasPrefix applies the HasPrefix predicate on the "flag" field.
func FlagHasPrefix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFlag), v))
	})
}

// FlagHasSuffix applies the HasSuffix predicate on the "flag" field.
func FlagHasSuffix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFlag), v))
	})
}

// FlagIsNil applies the IsNil predicate on the "flag" field.
func FlagIsNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFlag)))
	})
}

// FlagNotNil applies the NotNil predicate on the "flag" field.
func FlagNotNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFlag)))
	})
}

// FlagEqualFold applies the EqualFold predicate on the "flag" field.
func FlagEqualFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFlag), v))
	})
}

// FlagContainsFold applies the ContainsFold predicate on the "flag" field.
func FlagContainsFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFlag), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeIsNil applies the IsNil predicate on the "code" field.
func CodeIsNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCode)))
	})
}

// CodeNotNil applies the NotNil predicate on the "code" field.
func CodeNotNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCode)))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// ShortEQ applies the EQ predicate on the "short" field.
func ShortEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShort), v))
	})
}

// ShortNEQ applies the NEQ predicate on the "short" field.
func ShortNEQ(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShort), v))
	})
}

// ShortIn applies the In predicate on the "short" field.
func ShortIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldShort), v...))
	})
}

// ShortNotIn applies the NotIn predicate on the "short" field.
func ShortNotIn(vs ...string) predicate.Country {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldShort), v...))
	})
}

// ShortGT applies the GT predicate on the "short" field.
func ShortGT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShort), v))
	})
}

// ShortGTE applies the GTE predicate on the "short" field.
func ShortGTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShort), v))
	})
}

// ShortLT applies the LT predicate on the "short" field.
func ShortLT(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShort), v))
	})
}

// ShortLTE applies the LTE predicate on the "short" field.
func ShortLTE(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShort), v))
	})
}

// ShortContains applies the Contains predicate on the "short" field.
func ShortContains(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldShort), v))
	})
}

// ShortHasPrefix applies the HasPrefix predicate on the "short" field.
func ShortHasPrefix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldShort), v))
	})
}

// ShortHasSuffix applies the HasSuffix predicate on the "short" field.
func ShortHasSuffix(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldShort), v))
	})
}

// ShortIsNil applies the IsNil predicate on the "short" field.
func ShortIsNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShort)))
	})
}

// ShortNotNil applies the NotNil predicate on the "short" field.
func ShortNotNil() predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShort)))
	})
}

// ShortEqualFold applies the EqualFold predicate on the "short" field.
func ShortEqualFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldShort), v))
	})
}

// ShortContainsFold applies the ContainsFold predicate on the "short" field.
func ShortContainsFold(v string) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldShort), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Country) predicate.Country {
	return predicate.Country(func(s *sql.Selector) {
		p(s.Not())
	})
}
