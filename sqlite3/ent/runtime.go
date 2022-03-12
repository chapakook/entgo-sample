// Code generated by entc, DO NOT EDIT.

package ent

import (
	"sqlite3/ent/account"
	"sqlite3/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescMoney is the schema descriptor for money field.
	accountDescMoney := accountFields[1].Descriptor()
	// account.DefaultMoney holds the default value on creation for the money field.
	account.DefaultMoney = accountDescMoney.Default.(int)
}
