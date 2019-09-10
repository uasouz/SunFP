package sunfp

import "errors"

//IdFunction

func EitherIdentity(value interface{}) IEither {
	return value.(IEither)
}

//LEFT
type LeftDef struct {
	either
}

func Left(value interface{}) IEither {
	return LeftDef{either{value: value, isRight: false}}
}

//RIGHT
type RightDef struct {
	either
}

func Right(value interface{}) IEither {
	return RightDef{either{value: value, isRight: true}}
}

//EITHER

type either struct {
	value   interface{}
	isRight bool
}

type IEither interface {
	/* Inherited from Monad: */
	Bind(func(interface{}) IEither) IEither
	//flatMap<V>(fn: (val: T) => IEither<E, V>): IEither<E, V>;
	FlatMap(func(interface{}) IEither) IEither
	//chain<V>(fn: (val: T) => IEither<E, V>): IEither<E, V>;
	Chain(func(interface{}) IEither) IEither
	//map<V>(fn: (val: T) => V): IEither<E, V>;
	Map(func(interface{}) interface{}) IEither
	//join<V>(): IEither<E, V>; // if T is IEither<V>
	Join() IEither
	//takeLeft(m: IEither<E, T>): IEither<E, T>;
	TakeLeft(either IEither) IEither
	//takeRight(m: IEither<E, T>): IEither<E, T>;
	TakeRight(either IEither) IEither

	///* Inherited from Applicative */
	//ap<V>(eitherFn: IEither<E, (val: T) => V>): IEither<E, V>;
	Applicative(either IEither) IEither

	///* IEither specific */
	//cata<Z>(leftFn: (err: E) => Z, rightFn: (val: T) => Z): Z;
	Cata(func(interface{}) interface{}, func(interface{}) interface{}) interface{}
	//fold<Z>(leftFn: (err: E) => Z, rightFn: (val: T) => Z): Z;
	Fold(func(interface{}) interface{}, func(interface{}) interface{}) interface{}
	//catchMap<F>(fn: (err: E) => IEither<F, T>): IEither<F, T>;
	CatchMap(func(interface{}) IEither) IEither
	//swap(): IEither<T, E>;
	Swap() IEither
	//
	//bimap<Z, V>(leftFn: (err: E) => Z, rightFn: (val: T) => V): IEither<Z, V>;
	//leftMap<F>(fn: (leftVal: E) => F): IEither<F, T>;
	//
	//isRight(): boolean;
	IsRight() bool
	//isLeft(): boolean;
	IsLeft() bool
	//right(): T;
	Right() interface{}
	//left(): E;
	Left() interface{}
	//forEach(fn: (val: T) => void): void;
	ForEach(func(interface{}) interface{})
	//forEachLeft(fn: (val: E) => void): void;
	ForEachLeft(func(interface{}) interface{})
	//
	//toValidation(): Validation<E, T>;
	//toMaybe(): IMaybe<T>;
}

func (e either) Bind(fn func(interface{}) IEither) IEither {
	return e.FlatMap(fn)
}

func (e either) FlatMap(fn func(interface{}) IEither) IEither {
	if e.isRight {
		return fn(e.value)
	}
	return e
}

func (e either) Chain(fn func(interface{}) IEither) IEither {
	return e.FlatMap(fn)
}

func (e either) Map(fn func(interface{}) interface{}) IEither {
	if e.isRight {
		return Right(fn(e.value).(IEither))
	}
	return e
}

func (e either) LeftMap(fn func(interface{}) interface{}) IEither {
	if e.IsLeft() {
		return Left(fn(e.value))
	}
	return e
}

func (e either) Join() IEither {
	return e.FlatMap(EitherIdentity)
}

func (e either) TakeLeft(either IEither) IEither {
	panic("implement me")
}

func (e either) TakeRight(either IEither) IEither {
	panic("implement me")
}

func (e either) Applicative(either IEither) IEither {
	if e.isRight {
		return either.Right().(func(interface{}) IEither)(e.Right())
	}
	return e
}

func (e either) Cata(leftFn func(interface{}) interface{}, rightFn func(interface{}) interface{}) interface{} {
	if e.IsLeft() {
		return leftFn(e.value)
	}
	return rightFn(e.value)
}

func (e either) Fold(leftFn func(interface{}) interface{},rightFn func(interface{}) interface{}) interface{} {
	return e.Cata(leftFn,rightFn)
}

//Review
func (e either) CatchMap(fn func(interface{}) IEither) IEither {
	if e.IsRight() {
		return e
	}
	return fn(e.value)
}

func (e either) Swap() IEither {
	if e.IsRight() {
		return Left(e.value)
	}
	return Right(e.value)
}

func (e either) IsRight() bool {
	return e.isRight
}

func (e either) IsLeft() bool {
	return !e.isRight
}

func (e either) Right() interface{} {
	if e.isRight {
		return e.value
	}
	return errors.New("cannot call Right() from Left.")
}

func (e either) Left() interface{} {
	if !e.isRight {
		return e.value
	}
	return errors.New("cannot call Left() from Right.")
}

func (e either) ForEach(fn func(interface{}) interface{}) {
	e.Cata(noop, fn)
}

func (e either) ForEachLeft(fn func(interface{}) interface{}) {
	e.Cata(fn, noop)
}
