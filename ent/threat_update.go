// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hackerchai/threatbook-ip-web/ent/predicate"
	"github.com/hackerchai/threatbook-ip-web/ent/threat"
)

// ThreatUpdate is the builder for updating Threat entities.
type ThreatUpdate struct {
	config
	hooks    []Hook
	mutation *ThreatMutation
}

// Where appends a list predicates to the ThreatUpdate builder.
func (tu *ThreatUpdate) Where(ps ...predicate.Threat) *ThreatUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetIP sets the "ip" field.
func (tu *ThreatUpdate) SetIP(s string) *ThreatUpdate {
	tu.mutation.SetIP(s)
	return tu
}

// SetThreatIDInfo sets the "threat_id_info" field.
func (tu *ThreatUpdate) SetThreatIDInfo(s string) *ThreatUpdate {
	tu.mutation.SetThreatIDInfo(s)
	return tu
}

// SetDomainCount sets the "domain_count" field.
func (tu *ThreatUpdate) SetDomainCount(i int) *ThreatUpdate {
	tu.mutation.ResetDomainCount()
	tu.mutation.SetDomainCount(i)
	return tu
}

// AddDomainCount adds i to the "domain_count" field.
func (tu *ThreatUpdate) AddDomainCount(i int) *ThreatUpdate {
	tu.mutation.AddDomainCount(i)
	return tu
}

// SetTagCount sets the "tag_count" field.
func (tu *ThreatUpdate) SetTagCount(i int) *ThreatUpdate {
	tu.mutation.ResetTagCount()
	tu.mutation.SetTagCount(i)
	return tu
}

// AddTagCount adds i to the "tag_count" field.
func (tu *ThreatUpdate) AddTagCount(i int) *ThreatUpdate {
	tu.mutation.AddTagCount(i)
	return tu
}

// SetItelCount sets the "itel_count" field.
func (tu *ThreatUpdate) SetItelCount(i int) *ThreatUpdate {
	tu.mutation.ResetItelCount()
	tu.mutation.SetItelCount(i)
	return tu
}

// AddItelCount adds i to the "itel_count" field.
func (tu *ThreatUpdate) AddItelCount(i int) *ThreatUpdate {
	tu.mutation.AddItelCount(i)
	return tu
}

// SetJudge sets the "judge" field.
func (tu *ThreatUpdate) SetJudge(i int) *ThreatUpdate {
	tu.mutation.ResetJudge()
	tu.mutation.SetJudge(i)
	return tu
}

// AddJudge adds i to the "judge" field.
func (tu *ThreatUpdate) AddJudge(i int) *ThreatUpdate {
	tu.mutation.AddJudge(i)
	return tu
}

// SetPoc sets the "poc" field.
func (tu *ThreatUpdate) SetPoc(b bool) *ThreatUpdate {
	tu.mutation.SetPoc(b)
	return tu
}

// SetSource sets the "source" field.
func (tu *ThreatUpdate) SetSource(i int) *ThreatUpdate {
	tu.mutation.ResetSource()
	tu.mutation.SetSource(i)
	return tu
}

// AddSource adds i to the "source" field.
func (tu *ThreatUpdate) AddSource(i int) *ThreatUpdate {
	tu.mutation.AddSource(i)
	return tu
}

// SetCtime sets the "ctime" field.
func (tu *ThreatUpdate) SetCtime(i int64) *ThreatUpdate {
	tu.mutation.ResetCtime()
	tu.mutation.SetCtime(i)
	return tu
}

// AddCtime adds i to the "ctime" field.
func (tu *ThreatUpdate) AddCtime(i int64) *ThreatUpdate {
	tu.mutation.AddCtime(i)
	return tu
}

// Mutation returns the ThreatMutation object of the builder.
func (tu *ThreatUpdate) Mutation() *ThreatMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ThreatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ThreatUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ThreatUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ThreatUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *ThreatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   threat.Table,
			Columns: threat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: threat.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: threat.FieldIP,
		})
	}
	if value, ok := tu.mutation.ThreatIDInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: threat.FieldThreatIDInfo,
		})
	}
	if value, ok := tu.mutation.DomainCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldDomainCount,
		})
	}
	if value, ok := tu.mutation.AddedDomainCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldDomainCount,
		})
	}
	if value, ok := tu.mutation.TagCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldTagCount,
		})
	}
	if value, ok := tu.mutation.AddedTagCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldTagCount,
		})
	}
	if value, ok := tu.mutation.ItelCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldItelCount,
		})
	}
	if value, ok := tu.mutation.AddedItelCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldItelCount,
		})
	}
	if value, ok := tu.mutation.Judge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldJudge,
		})
	}
	if value, ok := tu.mutation.AddedJudge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldJudge,
		})
	}
	if value, ok := tu.mutation.Poc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: threat.FieldPoc,
		})
	}
	if value, ok := tu.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldSource,
		})
	}
	if value, ok := tu.mutation.AddedSource(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldSource,
		})
	}
	if value, ok := tu.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: threat.FieldCtime,
		})
	}
	if value, ok := tu.mutation.AddedCtime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: threat.FieldCtime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{threat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ThreatUpdateOne is the builder for updating a single Threat entity.
type ThreatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThreatMutation
}

// SetIP sets the "ip" field.
func (tuo *ThreatUpdateOne) SetIP(s string) *ThreatUpdateOne {
	tuo.mutation.SetIP(s)
	return tuo
}

// SetThreatIDInfo sets the "threat_id_info" field.
func (tuo *ThreatUpdateOne) SetThreatIDInfo(s string) *ThreatUpdateOne {
	tuo.mutation.SetThreatIDInfo(s)
	return tuo
}

// SetDomainCount sets the "domain_count" field.
func (tuo *ThreatUpdateOne) SetDomainCount(i int) *ThreatUpdateOne {
	tuo.mutation.ResetDomainCount()
	tuo.mutation.SetDomainCount(i)
	return tuo
}

// AddDomainCount adds i to the "domain_count" field.
func (tuo *ThreatUpdateOne) AddDomainCount(i int) *ThreatUpdateOne {
	tuo.mutation.AddDomainCount(i)
	return tuo
}

// SetTagCount sets the "tag_count" field.
func (tuo *ThreatUpdateOne) SetTagCount(i int) *ThreatUpdateOne {
	tuo.mutation.ResetTagCount()
	tuo.mutation.SetTagCount(i)
	return tuo
}

// AddTagCount adds i to the "tag_count" field.
func (tuo *ThreatUpdateOne) AddTagCount(i int) *ThreatUpdateOne {
	tuo.mutation.AddTagCount(i)
	return tuo
}

// SetItelCount sets the "itel_count" field.
func (tuo *ThreatUpdateOne) SetItelCount(i int) *ThreatUpdateOne {
	tuo.mutation.ResetItelCount()
	tuo.mutation.SetItelCount(i)
	return tuo
}

// AddItelCount adds i to the "itel_count" field.
func (tuo *ThreatUpdateOne) AddItelCount(i int) *ThreatUpdateOne {
	tuo.mutation.AddItelCount(i)
	return tuo
}

// SetJudge sets the "judge" field.
func (tuo *ThreatUpdateOne) SetJudge(i int) *ThreatUpdateOne {
	tuo.mutation.ResetJudge()
	tuo.mutation.SetJudge(i)
	return tuo
}

// AddJudge adds i to the "judge" field.
func (tuo *ThreatUpdateOne) AddJudge(i int) *ThreatUpdateOne {
	tuo.mutation.AddJudge(i)
	return tuo
}

// SetPoc sets the "poc" field.
func (tuo *ThreatUpdateOne) SetPoc(b bool) *ThreatUpdateOne {
	tuo.mutation.SetPoc(b)
	return tuo
}

// SetSource sets the "source" field.
func (tuo *ThreatUpdateOne) SetSource(i int) *ThreatUpdateOne {
	tuo.mutation.ResetSource()
	tuo.mutation.SetSource(i)
	return tuo
}

// AddSource adds i to the "source" field.
func (tuo *ThreatUpdateOne) AddSource(i int) *ThreatUpdateOne {
	tuo.mutation.AddSource(i)
	return tuo
}

// SetCtime sets the "ctime" field.
func (tuo *ThreatUpdateOne) SetCtime(i int64) *ThreatUpdateOne {
	tuo.mutation.ResetCtime()
	tuo.mutation.SetCtime(i)
	return tuo
}

// AddCtime adds i to the "ctime" field.
func (tuo *ThreatUpdateOne) AddCtime(i int64) *ThreatUpdateOne {
	tuo.mutation.AddCtime(i)
	return tuo
}

// Mutation returns the ThreatMutation object of the builder.
func (tuo *ThreatUpdateOne) Mutation() *ThreatMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ThreatUpdateOne) Select(field string, fields ...string) *ThreatUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Threat entity.
func (tuo *ThreatUpdateOne) Save(ctx context.Context) (*Threat, error) {
	var (
		err  error
		node *Threat
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Threat)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ThreatMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ThreatUpdateOne) SaveX(ctx context.Context) *Threat {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ThreatUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ThreatUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *ThreatUpdateOne) sqlSave(ctx context.Context) (_node *Threat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   threat.Table,
			Columns: threat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: threat.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Threat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, threat.FieldID)
		for _, f := range fields {
			if !threat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != threat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: threat.FieldIP,
		})
	}
	if value, ok := tuo.mutation.ThreatIDInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: threat.FieldThreatIDInfo,
		})
	}
	if value, ok := tuo.mutation.DomainCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldDomainCount,
		})
	}
	if value, ok := tuo.mutation.AddedDomainCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldDomainCount,
		})
	}
	if value, ok := tuo.mutation.TagCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldTagCount,
		})
	}
	if value, ok := tuo.mutation.AddedTagCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldTagCount,
		})
	}
	if value, ok := tuo.mutation.ItelCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldItelCount,
		})
	}
	if value, ok := tuo.mutation.AddedItelCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldItelCount,
		})
	}
	if value, ok := tuo.mutation.Judge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldJudge,
		})
	}
	if value, ok := tuo.mutation.AddedJudge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldJudge,
		})
	}
	if value, ok := tuo.mutation.Poc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: threat.FieldPoc,
		})
	}
	if value, ok := tuo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldSource,
		})
	}
	if value, ok := tuo.mutation.AddedSource(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: threat.FieldSource,
		})
	}
	if value, ok := tuo.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: threat.FieldCtime,
		})
	}
	if value, ok := tuo.mutation.AddedCtime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: threat.FieldCtime,
		})
	}
	_node = &Threat{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{threat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
