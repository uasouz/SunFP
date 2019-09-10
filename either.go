package sunfp

type Left struct {
	Either
}

type Right struct{
	Either
}

type Error struct {

}

type Either interface {
	/* Inherited from Monad: */
	Bind(func(interface{}) Either) Either
	//flatMap<V>(fn: (val: T) => Either<E, V>): Either<E, V>;
	FlatMap(func(interface{}) Either) Either
	//chain<V>(fn: (val: T) => Either<E, V>): Either<E, V>;
	Chain(func(interface{}) Either) Either
	//map<V>(fn: (val: T) => V): Either<E, V>;
	Map(func(interface{}) Either) Either
	//join<V>(): Either<E, V>; // if T is Either<V>
	Join() Either
	//takeLeft(m: Either<E, T>): Either<E, T>;
	TakeLeft(either Either) Either
	//takeRight(m: Either<E, T>): Either<E, T>;
	TakeRight(either Either) Either

	///* Inherited from Applicative */
	//ap<V>(eitherFn: Either<E, (val: T) => V>): Either<E, V>;
	Applicative(either Either) Either

	///* Either specific */
	//cata<Z>(leftFn: (err: E) => Z, rightFn: (val: T) => Z): Z;
	Cata(func(interface{}) interface{}, func(interface{}) interface{}) interface{}
	//fold<Z>(leftFn: (err: E) => Z, rightFn: (val: T) => Z): Z;
	Fold(func(interface{}) interface{}, func(interface{}) interface{}) interface{}
	//catchMap<F>(fn: (err: E) => Either<F, T>): Either<F, T>;
	CatchMap(func(interface{}) interface{}) Either
	//swap(): Either<T, E>;
	Swap() Either
	//
	//bimap<Z, V>(leftFn: (err: E) => Z, rightFn: (val: T) => V): Either<Z, V>;
	//leftMap<F>(fn: (leftVal: E) => F): Either<F, T>;
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
	ForEach(func(interface{}))
	//forEachLeft(fn: (val: E) => void): void;
	ForEachLeft(func(interface{}))
	//
	//toValidation(): Validation<E, T>;
	//toMaybe(): IMaybe<T>;
}
