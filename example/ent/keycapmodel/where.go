// Code generated by ent, DO NOT EDIT.

package keycapmodel

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ogen-go/ent2ogen/example/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Profile applies equality check predicate on the "profile" field. It's identical to ProfileEQ.
func Profile(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfile), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.KeycapModel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.KeycapModel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ProfileEQ applies the EQ predicate on the "profile" field.
func ProfileEQ(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfile), v))
	})
}

// ProfileNEQ applies the NEQ predicate on the "profile" field.
func ProfileNEQ(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProfile), v))
	})
}

// ProfileIn applies the In predicate on the "profile" field.
func ProfileIn(vs ...string) predicate.KeycapModel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProfile), v...))
	})
}

// ProfileNotIn applies the NotIn predicate on the "profile" field.
func ProfileNotIn(vs ...string) predicate.KeycapModel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProfile), v...))
	})
}

// ProfileGT applies the GT predicate on the "profile" field.
func ProfileGT(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProfile), v))
	})
}

// ProfileGTE applies the GTE predicate on the "profile" field.
func ProfileGTE(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProfile), v))
	})
}

// ProfileLT applies the LT predicate on the "profile" field.
func ProfileLT(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProfile), v))
	})
}

// ProfileLTE applies the LTE predicate on the "profile" field.
func ProfileLTE(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProfile), v))
	})
}

// ProfileContains applies the Contains predicate on the "profile" field.
func ProfileContains(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProfile), v))
	})
}

// ProfileHasPrefix applies the HasPrefix predicate on the "profile" field.
func ProfileHasPrefix(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProfile), v))
	})
}

// ProfileHasSuffix applies the HasSuffix predicate on the "profile" field.
func ProfileHasSuffix(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProfile), v))
	})
}

// ProfileEqualFold applies the EqualFold predicate on the "profile" field.
func ProfileEqualFold(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProfile), v))
	})
}

// ProfileContainsFold applies the ContainsFold predicate on the "profile" field.
func ProfileContainsFold(v string) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProfile), v))
	})
}

// MaterialEQ applies the EQ predicate on the "material" field.
func MaterialEQ(v Material) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaterial), v))
	})
}

// MaterialNEQ applies the NEQ predicate on the "material" field.
func MaterialNEQ(v Material) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaterial), v))
	})
}

// MaterialIn applies the In predicate on the "material" field.
func MaterialIn(vs ...Material) predicate.KeycapModel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaterial), v...))
	})
}

// MaterialNotIn applies the NotIn predicate on the "material" field.
func MaterialNotIn(vs ...Material) predicate.KeycapModel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KeycapModel(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaterial), v...))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KeycapModel) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KeycapModel) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
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
func Not(p predicate.KeycapModel) predicate.KeycapModel {
	return predicate.KeycapModel(func(s *sql.Selector) {
		p(s.Not())
	})
}
