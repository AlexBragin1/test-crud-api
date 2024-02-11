package filter

import "fmt"

const (
	DataTypeInt  = "int"
	DataTypeDate = "int64"
	

	OperatorEq            = "eg"
	OperatorNotEq         = "neg"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "gte"
	OperatorBetween       = ":"
)

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

func NewField() *Field {
	return &Field{}
}

func (f *Field) AddFields(name, value, operator, dtype string) error {

	f.Name = name
	f.Value = value
	f.Operator = operator
	f.Type = dtype

	err := validateOperator(operator)
	if err != nil {
		return err
	}

	return nil

}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThan:
	case OperatorGreaterThanEq:
	case OperatorBetween:
	default:
		return fmt.Errorf("bad operator")
	}
	return nil
}
