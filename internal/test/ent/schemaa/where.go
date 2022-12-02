// Code generated by ent, DO NOT EDIT.

package schemaa

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ogen-go/ent2ogen/internal/test/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Int64 applies equality check predicate on the "int64" field. It's identical to Int64EQ.
func Int64(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt64), v))
	})
}

// StringBindtoFoobar applies equality check predicate on the "string_bindto_foobar" field. It's identical to StringBindtoFoobarEQ.
func StringBindtoFoobar(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringOptionalNullable applies equality check predicate on the "string_optional_nullable" field. It's identical to StringOptionalNullableEQ.
func StringOptionalNullable(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringOptionalNullable), v))
	})
}

// OptionalNullableBool applies equality check predicate on the "optional_nullable_bool" field. It's identical to OptionalNullableBoolEQ.
func OptionalNullableBool(v bool) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOptionalNullableBool), v))
	})
}

// Int64EQ applies the EQ predicate on the "int64" field.
func Int64EQ(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt64), v))
	})
}

// Int64NEQ applies the NEQ predicate on the "int64" field.
func Int64NEQ(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInt64), v))
	})
}

// Int64In applies the In predicate on the "int64" field.
func Int64In(vs ...int64) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInt64), v...))
	})
}

// Int64NotIn applies the NotIn predicate on the "int64" field.
func Int64NotIn(vs ...int64) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInt64), v...))
	})
}

// Int64GT applies the GT predicate on the "int64" field.
func Int64GT(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInt64), v))
	})
}

// Int64GTE applies the GTE predicate on the "int64" field.
func Int64GTE(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInt64), v))
	})
}

// Int64LT applies the LT predicate on the "int64" field.
func Int64LT(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInt64), v))
	})
}

// Int64LTE applies the LTE predicate on the "int64" field.
func Int64LTE(v int64) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInt64), v))
	})
}

// StringBindtoFoobarEQ applies the EQ predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarEQ(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarNEQ applies the NEQ predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarNEQ(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarIn applies the In predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarIn(vs ...string) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStringBindtoFoobar), v...))
	})
}

// StringBindtoFoobarNotIn applies the NotIn predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarNotIn(vs ...string) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStringBindtoFoobar), v...))
	})
}

// StringBindtoFoobarGT applies the GT predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarGT(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarGTE applies the GTE predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarGTE(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarLT applies the LT predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarLT(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarLTE applies the LTE predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarLTE(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarContains applies the Contains predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarContains(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarHasPrefix applies the HasPrefix predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarHasPrefix(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarHasSuffix applies the HasSuffix predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarHasSuffix(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarEqualFold applies the EqualFold predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarEqualFold(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringBindtoFoobarContainsFold applies the ContainsFold predicate on the "string_bindto_foobar" field.
func StringBindtoFoobarContainsFold(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringBindtoFoobar), v))
	})
}

// StringOptionalNullableEQ applies the EQ predicate on the "string_optional_nullable" field.
func StringOptionalNullableEQ(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableNEQ applies the NEQ predicate on the "string_optional_nullable" field.
func StringOptionalNullableNEQ(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableIn applies the In predicate on the "string_optional_nullable" field.
func StringOptionalNullableIn(vs ...string) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStringOptionalNullable), v...))
	})
}

// StringOptionalNullableNotIn applies the NotIn predicate on the "string_optional_nullable" field.
func StringOptionalNullableNotIn(vs ...string) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStringOptionalNullable), v...))
	})
}

// StringOptionalNullableGT applies the GT predicate on the "string_optional_nullable" field.
func StringOptionalNullableGT(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableGTE applies the GTE predicate on the "string_optional_nullable" field.
func StringOptionalNullableGTE(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableLT applies the LT predicate on the "string_optional_nullable" field.
func StringOptionalNullableLT(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableLTE applies the LTE predicate on the "string_optional_nullable" field.
func StringOptionalNullableLTE(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableContains applies the Contains predicate on the "string_optional_nullable" field.
func StringOptionalNullableContains(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableHasPrefix applies the HasPrefix predicate on the "string_optional_nullable" field.
func StringOptionalNullableHasPrefix(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableHasSuffix applies the HasSuffix predicate on the "string_optional_nullable" field.
func StringOptionalNullableHasSuffix(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableIsNil applies the IsNil predicate on the "string_optional_nullable" field.
func StringOptionalNullableIsNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringOptionalNullable)))
	})
}

// StringOptionalNullableNotNil applies the NotNil predicate on the "string_optional_nullable" field.
func StringOptionalNullableNotNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringOptionalNullable)))
	})
}

// StringOptionalNullableEqualFold applies the EqualFold predicate on the "string_optional_nullable" field.
func StringOptionalNullableEqualFold(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringOptionalNullable), v))
	})
}

// StringOptionalNullableContainsFold applies the ContainsFold predicate on the "string_optional_nullable" field.
func StringOptionalNullableContainsFold(v string) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringOptionalNullable), v))
	})
}

// OptionalNullableBoolEQ applies the EQ predicate on the "optional_nullable_bool" field.
func OptionalNullableBoolEQ(v bool) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOptionalNullableBool), v))
	})
}

// OptionalNullableBoolNEQ applies the NEQ predicate on the "optional_nullable_bool" field.
func OptionalNullableBoolNEQ(v bool) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOptionalNullableBool), v))
	})
}

// OptionalNullableBoolIsNil applies the IsNil predicate on the "optional_nullable_bool" field.
func OptionalNullableBoolIsNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOptionalNullableBool)))
	})
}

// OptionalNullableBoolNotNil applies the NotNil predicate on the "optional_nullable_bool" field.
func OptionalNullableBoolNotNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOptionalNullableBool)))
	})
}

// JsontypeStringsOptionalIsNil applies the IsNil predicate on the "jsontype_strings_optional" field.
func JsontypeStringsOptionalIsNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldJsontypeStringsOptional)))
	})
}

// JsontypeStringsOptionalNotNil applies the NotNil predicate on the "jsontype_strings_optional" field.
func JsontypeStringsOptionalNotNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldJsontypeStringsOptional)))
	})
}

// RequiredEnumEQ applies the EQ predicate on the "required_enum" field.
func RequiredEnumEQ(v RequiredEnum) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequiredEnum), v))
	})
}

// RequiredEnumNEQ applies the NEQ predicate on the "required_enum" field.
func RequiredEnumNEQ(v RequiredEnum) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRequiredEnum), v))
	})
}

// RequiredEnumIn applies the In predicate on the "required_enum" field.
func RequiredEnumIn(vs ...RequiredEnum) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRequiredEnum), v...))
	})
}

// RequiredEnumNotIn applies the NotIn predicate on the "required_enum" field.
func RequiredEnumNotIn(vs ...RequiredEnum) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRequiredEnum), v...))
	})
}

// OptionalNullableEnumEQ applies the EQ predicate on the "optional_nullable_enum" field.
func OptionalNullableEnumEQ(v OptionalNullableEnum) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOptionalNullableEnum), v))
	})
}

// OptionalNullableEnumNEQ applies the NEQ predicate on the "optional_nullable_enum" field.
func OptionalNullableEnumNEQ(v OptionalNullableEnum) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOptionalNullableEnum), v))
	})
}

// OptionalNullableEnumIn applies the In predicate on the "optional_nullable_enum" field.
func OptionalNullableEnumIn(vs ...OptionalNullableEnum) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOptionalNullableEnum), v...))
	})
}

// OptionalNullableEnumNotIn applies the NotIn predicate on the "optional_nullable_enum" field.
func OptionalNullableEnumNotIn(vs ...OptionalNullableEnum) predicate.SchemaA {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOptionalNullableEnum), v...))
	})
}

// OptionalNullableEnumIsNil applies the IsNil predicate on the "optional_nullable_enum" field.
func OptionalNullableEnumIsNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOptionalNullableEnum)))
	})
}

// OptionalNullableEnumNotNil applies the NotNil predicate on the "optional_nullable_enum" field.
func OptionalNullableEnumNotNil() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOptionalNullableEnum)))
	})
}

// HasEdgeSchemabUniqueRequired applies the HasEdge predicate on the "edge_schemab_unique_required" edge.
func HasEdgeSchemabUniqueRequired() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabUniqueRequiredTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EdgeSchemabUniqueRequiredTable, EdgeSchemabUniqueRequiredColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgeSchemabUniqueRequiredWith applies the HasEdge predicate on the "edge_schemab_unique_required" edge with a given conditions (other predicates).
func HasEdgeSchemabUniqueRequiredWith(preds ...predicate.SchemaB) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabUniqueRequiredInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EdgeSchemabUniqueRequiredTable, EdgeSchemabUniqueRequiredColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgeSchemabUniqueRequiredBindtoBs applies the HasEdge predicate on the "edge_schemab_unique_required_bindto_bs" edge.
func HasEdgeSchemabUniqueRequiredBindtoBs() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabUniqueRequiredBindtoBsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EdgeSchemabUniqueRequiredBindtoBsTable, EdgeSchemabUniqueRequiredBindtoBsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgeSchemabUniqueRequiredBindtoBsWith applies the HasEdge predicate on the "edge_schemab_unique_required_bindto_bs" edge with a given conditions (other predicates).
func HasEdgeSchemabUniqueRequiredBindtoBsWith(preds ...predicate.SchemaB) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabUniqueRequiredBindtoBsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EdgeSchemabUniqueRequiredBindtoBsTable, EdgeSchemabUniqueRequiredBindtoBsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgeSchemabUniqueOptional applies the HasEdge predicate on the "edge_schemab_unique_optional" edge.
func HasEdgeSchemabUniqueOptional() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabUniqueOptionalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EdgeSchemabUniqueOptionalTable, EdgeSchemabUniqueOptionalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgeSchemabUniqueOptionalWith applies the HasEdge predicate on the "edge_schemab_unique_optional" edge with a given conditions (other predicates).
func HasEdgeSchemabUniqueOptionalWith(preds ...predicate.SchemaB) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabUniqueOptionalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EdgeSchemabUniqueOptionalTable, EdgeSchemabUniqueOptionalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgeSchemab applies the HasEdge predicate on the "edge_schemab" edge.
func HasEdgeSchemab() predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgeSchemabTable, EdgeSchemabColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgeSchemabWith applies the HasEdge predicate on the "edge_schemab" edge with a given conditions (other predicates).
func HasEdgeSchemabWith(preds ...predicate.SchemaB) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgeSchemabInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgeSchemabTable, EdgeSchemabColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SchemaA) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SchemaA) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
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
func Not(p predicate.SchemaA) predicate.SchemaA {
	return predicate.SchemaA(func(s *sql.Selector) {
		p(s.Not())
	})
}
