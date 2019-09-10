package sunfp

import (
	"reflect"
)

//MAYBE
func Maybe(value interface{}) IMaybe{
	if value == nil {
		return Nothing()
	}
	return Some(value)
}

type IMaybe interface {
	FlatMap(fn func(value interface{}) IMaybe) IMaybe
	Of(of interface{}) IMaybe
	Map(fn func(interface{}) interface{}) IMaybe
	Or(or interface{}) interface{}
	IsPresent() bool
	IsNil() bool
	IsSome() bool
	IsNothing() bool
	Unwrap() interface{}
	Type() reflect.Type
	Kind() reflect.Kind
	IsType(t reflect.Type) bool
	IsKind(t reflect.Kind) bool
}

//NOTHING
type NothingDef struct {
}

func Nothing() IMaybe {
	var noValue interface{}
	return NothingDef{}.Of(noValue)
}


func (maybe NothingDef) Of(of interface{}) IMaybe {
	return NothingDef{}
}

func (maybe NothingDef) FlatMap(fn func(value interface{}) IMaybe) IMaybe {
	return Nothing()
}

func (maybe NothingDef) Map(fn func(interface{}) interface{}) IMaybe {
	return Nothing()
}

// Or Check the value wrapped by NothingDef, if it's nil then return a given fallback value
func (maybe NothingDef) Or(or interface{}) interface{} {
		return or
}

// IsPresent Check is it present(not nil)
func (maybe NothingDef) IsPresent() bool {
	return false
}

// IsNil Check is it nil
func (maybe NothingDef) IsNil() bool {
	return true
}

func (maybe NothingDef) IsSome() bool {
	return true
}

func (maybe NothingDef) IsNothing() bool {
	return false
}

// Unwrap Unwrap the wrapped value of NothingDef
func (maybe NothingDef) Unwrap() interface{} {
		return "<Nothing>"
}

// Type Get its Type
func (maybe NothingDef) Type() reflect.Type {
	if maybe.IsNil() {
		return reflect.TypeOf(nil)
	}
	return reflect.TypeOf(Nothing())
}

// Kind Get its Kind
func (maybe NothingDef) Kind() reflect.Kind {
	return reflect.ValueOf(Nothing()).Kind()
}

// IsType Check is its Type equal to the given one
func (maybe NothingDef) IsType(t reflect.Type) bool {
	return maybe.Type() == t
}

// IsKind Check is its Kind equal to the given one
func (maybe NothingDef) IsKind(t reflect.Kind) bool {
	return maybe.Kind() == t
}


//SOME
type SomeDef struct {
	values interface{}
	//valueType reflect.Type
}

func Some(of interface{}) IMaybe {
	return SomeDef{}.Of(of)
}

//func IMaybe(of interface{}) *IMaybe {
//	return IMaybe{}.of(of)
//}

func (maybe SomeDef) Of(of interface{}) IMaybe {
	return SomeDef{values:of}
}

func (maybe SomeDef) FlatMap(fn func(value interface{}) IMaybe) IMaybe {
	return fn(maybe.values)
}

func (maybe SomeDef) Map(fn func(interface{}) interface{}) IMaybe {
	return Some(fn(maybe.values))
}

// Or Check the value wrapped by SomeDef, if it's nil then return a given fallback value
func (maybe SomeDef) Or(or interface{}) interface{} {
	if maybe.IsNil() {
		return or
	}

	return maybe.values
}

// IsPresent Check is it present(not nil)
func (maybe SomeDef) IsPresent() bool {
	return !(maybe.IsNil())
}

// IsNil Check is it nil
func (maybe SomeDef) IsNil() bool {
	val := reflect.ValueOf(maybe.values)

	if maybe.Kind() == reflect.Ptr {
		return val.IsNil()
	}
	return !val.IsValid()
}

func (maybe SomeDef) IsSome() bool {
	return true
}

func (maybe SomeDef) IsNothing() bool {
	return false
}

// Unwrap Unwrap the wrapped value of SomeDef
func (maybe SomeDef) Unwrap() interface{} {
	if maybe.IsNil() || maybe.IsNothing() {
		return "<Nothing>"
	}

	return maybe.values
}

// Type Get its Type
func (maybe SomeDef) Type() reflect.Type {
	if maybe.IsNil() {
		return reflect.TypeOf(nil)
	}
	return reflect.TypeOf(maybe.values)
}

// Kind Get its Kind
func (maybe SomeDef) Kind() reflect.Kind {
	return reflect.ValueOf(maybe.values).Kind()
}

// IsType Check is its Type equal to the given one
func (maybe SomeDef) IsType(t reflect.Type) bool {
	return maybe.Type() == t
}

// IsKind Check is its Kind equal to the given one
func (maybe SomeDef) IsKind(t reflect.Kind) bool {
	return maybe.Kind() == t
}
