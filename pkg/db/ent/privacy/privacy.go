// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/g11n-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AppCountryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppCountryQueryRuleFunc func(context.Context, *ent.AppCountryQuery) error

// EvalQuery return f(ctx, q).
func (f AppCountryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppCountryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppCountryQuery", q)
}

// The AppCountryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppCountryMutationRuleFunc func(context.Context, *ent.AppCountryMutation) error

// EvalMutation calls f(ctx, m).
func (f AppCountryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppCountryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppCountryMutation", m)
}

// The AppLangQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppLangQueryRuleFunc func(context.Context, *ent.AppLangQuery) error

// EvalQuery return f(ctx, q).
func (f AppLangQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppLangQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppLangQuery", q)
}

// The AppLangMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppLangMutationRuleFunc func(context.Context, *ent.AppLangMutation) error

// EvalMutation calls f(ctx, m).
func (f AppLangMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppLangMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppLangMutation", m)
}

// The CountryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CountryQueryRuleFunc func(context.Context, *ent.CountryQuery) error

// EvalQuery return f(ctx, q).
func (f CountryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CountryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CountryQuery", q)
}

// The CountryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CountryMutationRuleFunc func(context.Context, *ent.CountryMutation) error

// EvalMutation calls f(ctx, m).
func (f CountryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CountryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CountryMutation", m)
}

// The LangQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LangQueryRuleFunc func(context.Context, *ent.LangQuery) error

// EvalQuery return f(ctx, q).
func (f LangQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LangQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LangQuery", q)
}

// The LangMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LangMutationRuleFunc func(context.Context, *ent.LangMutation) error

// EvalMutation calls f(ctx, m).
func (f LangMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LangMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LangMutation", m)
}

// The MessageQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MessageQueryRuleFunc func(context.Context, *ent.MessageQuery) error

// EvalQuery return f(ctx, q).
func (f MessageQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MessageQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MessageQuery", q)
}

// The MessageMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MessageMutationRuleFunc func(context.Context, *ent.MessageMutation) error

// EvalMutation calls f(ctx, m).
func (f MessageMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MessageMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MessageMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AppCountryQuery:
		return q.Filter(), nil
	case *ent.AppLangQuery:
		return q.Filter(), nil
	case *ent.CountryQuery:
		return q.Filter(), nil
	case *ent.LangQuery:
		return q.Filter(), nil
	case *ent.MessageQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AppCountryMutation:
		return m.Filter(), nil
	case *ent.AppLangMutation:
		return m.Filter(), nil
	case *ent.CountryMutation:
		return m.Filter(), nil
	case *ent.LangMutation:
		return m.Filter(), nil
	case *ent.MessageMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
