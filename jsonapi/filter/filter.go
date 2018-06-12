package filter

import (
	"encoding/json"
	"errors"
	"fmt"
)

var dictOper = map[string]string{
	// Unary operators.
	"is_null":     "IS NULL",
	"is_not_null": "IS NOT NULL",
	// Binary operators.
	"==":             "=",
	"eq":             "=",
	"equals":         "=",
	"equal_to":       "=",
	"!=":             "!=",
	"ne":             "!=",
	"neq":            "!=",
	"not_equal_to":   "!=",
	"does_not_equal": "!=",
	">":              ">",
	"gt":             ">",
	"<":              "<",
	"lt":             "<",
	">=":             ">=",
	"ge":             ">=",
	"gte":            ">=",
	"geq":            ">=",
	"<=":             "<=",
	"le":             "<=",
	"lte":            "<=",
	"leq":            "<=",
	//"<<": inet_is_contained_by,
	//"<<=": inet_is_contained_by_or_equals,
	//">>": inet_contains,
	//">>=": inet_contains_or_equals,
	//"<>": inet_not_equals,
	//"&&": inet_contains_or_is_contained_by,
	"ilike": "ILIKE",
	"like":  "LIKE",
	//"not_like": not_like,
	"in":     "IN ",
	"not_in": "NOT IN",
	// (Binary) relationship operators.
	//"has": has,
	//"any": any_,
}

// FieldFilter -
type FieldFilter struct {
	Field    string
	Operator string
	Argument string
}

// NewFieldFilter -
func NewFieldFilter(field string, oper string, arg string) *FieldFilter {
	return &FieldFilter{
		Field:    field,
		Operator: oper,
		Argument: arg,
	}
}

// CreateFilter -
func CreateFilter(filters string) (string, error) {
	result := ""
	list := []map[string]interface{}{}

	if err := json.Unmarshal([]byte(filters), &list); err != nil {
		return result, err
	}

	for _, e := range list {
		var err error

		if _, ok := e["name"]; ok {
			if result, err = field(e, result); err != nil {
				return result, err
			}
			continue
		}

		if _, ok := e["and"]; ok {
			if result, err = and(e, result); err != nil {
				return result, err
			}
			continue
		}

		if _, ok := e["or"]; ok {
			if result, err = or(e, result); err != nil {
				return result, err
			}
			continue
		}

		if _, ok := e["not"]; ok {
			if result, err = not(e, result); err != nil {
				return result, err
			}
			continue
		}

		return result, errors.New("Undefined key attribute")
	}

	return result, nil
}

func field(v interface{}, r string) (string, error) {
	var f *FieldFilter
	var err error

	if f, err = fieldFilterFrom(v); err != nil {
		return r, err
	}
	return r + fmt.Sprintf(" AND %s %s %s", f.Field, f.Operator, f.Argument), nil
}

func fieldFilterFrom(v interface{}) (*FieldFilter, error) {
	f := new(FieldFilter)

	if filter, ok := v.(map[string]interface{}); ok {
		if name, ok := filter["name"].(string); ok {
			f.Field = name
		} else {
			return f, errors.New("Error: get 'name' from field filter")
		}

		if oper, ok := filter["op"].(string); ok {
			if f.Operator, ok = dictOper[oper]; !ok {
				return f, errors.New("Error: get 'op' from operation dictionary")
			}
		} else {
			return f, errors.New("Error: get 'op' from field filter")
		}

		if val, ok := filter["val"].(string); ok {
			f.Argument = val
		} else {
			return f, errors.New("Error: get 'name' from field filter")
		}
	}
	return f, errors.New("Field filter must be JSON-object")
}

func and(v interface{}, r string) (string, error) {
	return r, nil
}

func or(v interface{}, r string) (string, error) {
	return r, nil
}

func not(v interface{}, r string) (string, error) {
	return r, nil
}
