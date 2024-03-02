package main

/*
	what is reflection

	Reflection in computing is the ability of a program to examine its own structure,
	particularly through types; it's a form of metaprogramming. It's also a great
	source of confusion.

	what is interface{}

	We have enjoyed the type-safety that Go has offered us in terms of functions
	that work with known types, such as string, int and our own types like
	BankAccount.

	This means that we get some documentation for free and the compiler will
	complain if you try and pass the wrong type to a function.

	You may come across scenarios though where you want to write a function
	where you don't know the type at compile time.

	Go lets us get around this with the type interface{} which you can think
	of as just any type (in fact, in Go any is an

	So walk(x interface{}, fn func(string)) will accept any value for x.
*/

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
