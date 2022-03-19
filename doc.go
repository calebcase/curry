// Package curry provides facilities for currying functions. The process
// converts a single function like `f(x, y, z)` into N many functions each
// taking a single argument like `h(x)(y)(z)`.
//
// The functions are named by the number of arguments and return parameters of
// the target function. For example, to wrap a function with 2 arguments and 1
// return `func F(x, y int) int` use `curry.A2R1(F)`.
package curry
