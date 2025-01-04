# Unexpected Behavior with Defer and Variable Modification in Go

This example demonstrates a subtle issue related to Go's `defer` statement and variable modification.  The `defer` statement's behavior might not be intuitive in scenarios where a variable's value changes after the `defer` is defined, but before its execution.

## The Problem

The code in `bug.go` showcases the surprising behavior. We expect the deferred function to capture the initial value of `x`, which is 10. However, it reflects the modified value (5).  This is because closures in Go capture variables by reference, not value.

## The Solution

The solution (`bugSolution.go`) addresses this issue by creating a copy of the variable before the `defer` is called. This ensures the deferred function always works with the expected initial value.
