// Code generated by entc, DO NOT EDIT.

package tag

import (
	"TH0400/repo/ent/predicate"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TagName applies equality check predicate on the "tag_name" field. It's identical to TagNameEQ.
func TagName(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagName), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// Baned applies equality check predicate on the "baned" field. It's identical to BanedEQ.
func Baned(v bool) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaned), v))
	})
}

// TagNameEQ applies the EQ predicate on the "tag_name" field.
func TagNameEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagName), v))
	})
}

// TagNameNEQ applies the NEQ predicate on the "tag_name" field.
func TagNameNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagName), v))
	})
}

// TagNameIn applies the In predicate on the "tag_name" field.
func TagNameIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTagName), v...))
	})
}

// TagNameNotIn applies the NotIn predicate on the "tag_name" field.
func TagNameNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTagName), v...))
	})
}

// TagNameGT applies the GT predicate on the "tag_name" field.
func TagNameGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTagName), v))
	})
}

// TagNameGTE applies the GTE predicate on the "tag_name" field.
func TagNameGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTagName), v))
	})
}

// TagNameLT applies the LT predicate on the "tag_name" field.
func TagNameLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTagName), v))
	})
}

// TagNameLTE applies the LTE predicate on the "tag_name" field.
func TagNameLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTagName), v))
	})
}

// TagNameContains applies the Contains predicate on the "tag_name" field.
func TagNameContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTagName), v))
	})
}

// TagNameHasPrefix applies the HasPrefix predicate on the "tag_name" field.
func TagNameHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTagName), v))
	})
}

// TagNameHasSuffix applies the HasSuffix predicate on the "tag_name" field.
func TagNameHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTagName), v))
	})
}

// TagNameEqualFold applies the EqualFold predicate on the "tag_name" field.
func TagNameEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTagName), v))
	})
}

// TagNameContainsFold applies the ContainsFold predicate on the "tag_name" field.
func TagNameContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTagName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// BanedEQ applies the EQ predicate on the "baned" field.
func BanedEQ(v bool) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaned), v))
	})
}

// BanedNEQ applies the NEQ predicate on the "baned" field.
func BanedNEQ(v bool) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaned), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
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
func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		p(s.Not())
	})
}