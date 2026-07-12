// Package limits contains request limit configuration.
//
// A nil Config or a zero field uses the built-in default of 1000 items. Set a
// positive field value to override the default; lower values configure stricter
// limits. Zero does not disable an operation.
package limits
