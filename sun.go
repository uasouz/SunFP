package sunfp

type Functor interface {
	Map(fn func(interface{}) interface{}) interface{}
}

type Chain interface {

}

type Bind interface {
	Chain
}

type Applicative interface {

}

type CataMorphism interface {
	Cata(func(Error) interface{}, func(interface{}) interface{}) interface{}
}