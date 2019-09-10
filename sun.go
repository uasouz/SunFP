package sunfp

//Does nothing
func Noop(){

}

//Does nothing into a fancy way
func noop(interface{}) interface{} { return "" }


type Functor interface {
	Map(fn func(interface{}) interface{}) interface{}
}

type Chain interface {
}

type Bind interface {
	Chain
}

/* Applicative allows applying wrapped functions to wrapped elements */
// https://github.com/fantasyland/fantasy-land#applicative
type Applicative interface {
	//ap<V>(afn: Applicative<(val: T) => V>): Applicative<V>
}

type CataMorphism interface {
	Cata(func(interface{}) interface{}, func(interface{}) interface{}) interface{}
}

type IFree interface {
	//	bind<V>(fn: (val: A) => Free<V>): Free<V>;
	//flatMap<V>(fn: (val: A) => Free<V>): Free<V>;
	//chain<V>(fn: (val: A) => Free<V>): Free<V>;
	//join<V>(): Free<V>; // only if A = Free<V> on the same functor
	//map<V>(fn: (val: A) => V): Free<V>;
	//takeLeft<X>(other: Free<X>): Free<A>;
	//takeRight<B>(other: Free<B>): Free<B>;
	//
	///* Free-specific: */
	//// evaluates a single layer
	//resume<FFA>(): Either<FFA, A>;
	//// runs to completion using given extraction function:
	//go<FFA>(extract: (sus: FFA) => Free<A>): A;
}

func foldRight(fn func(), list []interface{}, acc interface{}) {
	//return func (innerList []interface{}, innerAcc interface{}) {
	panic("implement me")
	//}(list,acc)
}

func Identity(value interface{}) interface{} {
	return value
}
