// Code generated by entc, DO NOT EDIT.

package useruser

import (
	"TH0400/repo/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IdolID applies equality check predicate on the "idol_id" field. It's identical to IdolIDEQ.
func IdolID(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdolID), v))
	})
}

// FollowerID applies equality check predicate on the "follower_id" field. It's identical to FollowerIDEQ.
func FollowerID(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFollowerID), v))
	})
}

// IdolIDEQ applies the EQ predicate on the "idol_id" field.
func IdolIDEQ(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdolID), v))
	})
}

// IdolIDNEQ applies the NEQ predicate on the "idol_id" field.
func IdolIDNEQ(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdolID), v))
	})
}

// IdolIDIn applies the In predicate on the "idol_id" field.
func IdolIDIn(vs ...int) predicate.UserUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdolID), v...))
	})
}

// IdolIDNotIn applies the NotIn predicate on the "idol_id" field.
func IdolIDNotIn(vs ...int) predicate.UserUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdolID), v...))
	})
}

// IdolIDGT applies the GT predicate on the "idol_id" field.
func IdolIDGT(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdolID), v))
	})
}

// IdolIDGTE applies the GTE predicate on the "idol_id" field.
func IdolIDGTE(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdolID), v))
	})
}

// IdolIDLT applies the LT predicate on the "idol_id" field.
func IdolIDLT(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdolID), v))
	})
}

// IdolIDLTE applies the LTE predicate on the "idol_id" field.
func IdolIDLTE(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdolID), v))
	})
}

// FollowerIDEQ applies the EQ predicate on the "follower_id" field.
func FollowerIDEQ(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFollowerID), v))
	})
}

// FollowerIDNEQ applies the NEQ predicate on the "follower_id" field.
func FollowerIDNEQ(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFollowerID), v))
	})
}

// FollowerIDIn applies the In predicate on the "follower_id" field.
func FollowerIDIn(vs ...int) predicate.UserUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFollowerID), v...))
	})
}

// FollowerIDNotIn applies the NotIn predicate on the "follower_id" field.
func FollowerIDNotIn(vs ...int) predicate.UserUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFollowerID), v...))
	})
}

// FollowerIDGT applies the GT predicate on the "follower_id" field.
func FollowerIDGT(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFollowerID), v))
	})
}

// FollowerIDGTE applies the GTE predicate on the "follower_id" field.
func FollowerIDGTE(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFollowerID), v))
	})
}

// FollowerIDLT applies the LT predicate on the "follower_id" field.
func FollowerIDLT(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFollowerID), v))
	})
}

// FollowerIDLTE applies the LTE predicate on the "follower_id" field.
func FollowerIDLTE(v int) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFollowerID), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.UserUser) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.UserUser) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
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
func Not(p predicate.UserUser) predicate.UserUser {
	return predicate.UserUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
