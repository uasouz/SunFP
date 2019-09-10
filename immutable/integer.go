package immutable

import "math"

type IInteger interface {
	Add(interface{}) IInteger
	Sub(interface{}) IInteger
	Div(interface{}) IInteger
	Mul(interface{}) IInteger
	Pow(interface{}) IInteger
	Op(func(interface{}) int) IInteger
	Equal(interface{}) bool
	Set(interface{}) IInteger
}

func Integer(value interface{}) IInteger {
	switch value.(type) {
	case int:
		{
			return integer{value: value.(int), isSet: true}
		}
	case IInteger:
		{
			return integer{value: value.(integer).value, isSet: true}
		}
	}
	return integer{}
}

type integer struct {
	value int
	isSet bool
}

func (i integer) Set(value interface{}) IInteger {
	if !i.isSet {
		i.value = value.(int)
		i.isSet = true
	}
	return Integer(value)
}

func unwrapValue(value interface{}) (int, bool) {
	switch value.(type) {
	case int:
		{
			return value.(int), true
		}
	case IInteger:
		{
			return value.(integer).value, true
		}
	}
	return 0, false
}

func (i integer) Equal(value interface{}) bool {
	switch value.(type) {
	case int:
		{
			return i.value == value
		}
	case IInteger:
		{
			return i == value
		}
	}
	return false
}

func (i integer) Add(value interface{}) IInteger {
	if uValue, ok := unwrapValue(value); ok {
		return Integer(i.value + uValue)
	}
	return integer{}
}

func (i integer) Sub(value interface{}) IInteger {
	if uValue, ok := unwrapValue(value); ok {
		return Integer(i.value - uValue)
	}
	return integer{}
}

func (i integer) Div(value interface{}) IInteger {
	if uValue, ok := unwrapValue(value); ok {
		return Integer(i.value / uValue)
	}
	return integer{}
}

func (i integer) Mul(value interface{}) IInteger {
	if uValue, ok := unwrapValue(value); ok {
		return Integer(i.value * uValue)
	}
	return integer{}
}

func (i integer) Pow(value interface{}) IInteger {
	if uValue,ok := unwrapValue(value); ok {
		return Integer(int(math.Pow(float64(i.value),float64(uValue))))
	}
	return integer{}
}

func (i integer) Op(fn func(value interface{}) int) IInteger {
		return Integer(fn(i.value))
}
