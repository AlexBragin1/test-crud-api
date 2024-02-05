package filter

import "fmt"

const (
	DataTypeInt  = "int"
	DataTypeDate = "int64"

	OperatorEq = "eg"
	//OperatorNotEq         = "neg"
	//OperatorLowerThan     = "lt"
	//OperatorLowerThanEq   = "lte"
	//OperatorGreaterThan   = "gt"
	//OperatorGreaterThanEq = "gte"
	OperatorBetween = "between"
	//OperatorLike          = "Like"
)

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}
type options struct {
	//	isToApply bool
	limit  int
	fields []Field
}

// AddFiedls implements Options.
func (*options) AddFiedls(name string, operator string, value string, dtype string) error {
	panic("unimplemented")
}

type Options interface {
	GetLimit() int
	AddFields(name, operator, value, dtype string) error
	Fields() []Field
}

func NewOptions(limit int) Options {
	return &options{limit: limit}
}
func (o *options) GetLimit() int {
	return o.limit
}

func (o *options) AddFields(name, operator, value, dtype string) error {

	err := validateOperator(operator)
	if err != nil {
		return err
	}

	o.fields = append(o.fields, Field{Name: name,
		Value:    value,
		Operator: operator,
		Type:     dtype,
	})
	return nil

}
func (o *options) Fields() []Field {
	return o.fields
}
func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	//case OperatorNotEq:
	//case OperatorLowerThan:
	//case OperatorLowerThanEq:
	//case OperatorGreaterThan:
	//case OperatorGreaterThanEq:
	case OperatorBetween:
	default:
		return fmt.Errorf("bad operator")
	}
	return nil
}
