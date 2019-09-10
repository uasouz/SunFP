package test

import (
	"fmt"
	"github.com/uasouz/sunfp"
	"github.com/uasouz/sunfp/immutable"
	"testing"
)

type User struct {
	Name  string
	Level immutable.IInteger
}

func setUserLevel(user interface{}) interface{} {
	u := user.(User)
	u.Level = immutable.Integer(5)
	return u
}

func changeUser(user interface{}) interface{} {
	u := user.(User)
	u.Level = u.Level.Add(10).Pow(2).Mul(2)
	fmt.Println(u)
	return u
}

func validateUserName(user interface{}) sunfp.IMaybe {
	u := user.(User)
	if u.Name != "" {
		return sunfp.Some(u)
	}
	return sunfp.Nothing()
}

func setupUser(userOption sunfp.IMaybe) sunfp.IMaybe {
	return userOption.FlatMap(validateUserName).Map(setUserLevel).Map(changeUser)
}

func TestMaybe(t *testing.T) {
	hadara := User{Name: "Hadara"}
	hadara = setupUser(sunfp.Maybe(hadara)).Unwrap().(User)
	fmt.Print(hadara)
	if !hadara.Level.Equal(immutable.Integer(450)) {
		t.Error("Wrong Value")
	}
}
