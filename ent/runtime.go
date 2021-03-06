// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/GavukaAlexandr/ent-sandbox/ent/pet"
	"github.com/GavukaAlexandr/ent-sandbox/ent/schema"
	"github.com/GavukaAlexandr/ent-sandbox/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[1].Descriptor()
	// pet.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	pet.AgeValidator = petDescAge.Validators[0].(func(int) error)
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.DefaultID holds the default value on creation for the id field.
	pet.DefaultID = petDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
