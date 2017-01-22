// Package match contains special functions recognized by the dont template matcher.
package match

// NoInterveningStatements matches code
// where the preceding and succeeding statements
// occur with no intervening statements.
func NoInterveningStatements() {}

// Used matches code where x is used.
// TODO: Exactly what kinds of things can it match?
// At a minimum needs to match idents and method expressions.
func Used(x interface{}) {}

// Unused matches code in which x is never used.
func Unused(x interface{}) {}

// PossiblyUnused matches code where there is a non-panic
// way to exit the current function without using x.
// Use it in places to find code that should be doing defer <some code that uses x>.
func PossiblyUnused(x interface{}) {}

// NotType matches code in which val does not have the same concrete type as typ.
func NotType(val, typ interface{}) {}
