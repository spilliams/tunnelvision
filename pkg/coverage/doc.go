// Package coverage is a helper library that seeks to provide coverage
// about a terraform module. You can call it on a terraform root directory
// (where a complete terraform configuration is held), and it will report to you
// which resources are "covered" and which are not.
//
// It does this by comparing its understanding of what's in the module or root
// (given by github.com/hashicorp/terraform-config-inspect/tfconfig) with any
// number of "actuals" representing plan output (in the format of the Plan
// struct from github.com/hashicorp/terraform-json/)
package coverage
