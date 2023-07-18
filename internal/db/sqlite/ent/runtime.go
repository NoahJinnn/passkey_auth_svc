// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/income"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/institution"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/manualitem"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/schema"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/todo"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/transaction"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[3].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[4].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	connectionFields := schema.Connection{}.Fields()
	_ = connectionFields
	// connectionDescCreatedAt is the schema descriptor for created_at field.
	connectionDescCreatedAt := connectionFields[3].Descriptor()
	// connection.DefaultCreatedAt holds the default value on creation for the created_at field.
	connection.DefaultCreatedAt = connectionDescCreatedAt.Default.(func() time.Time)
	// connectionDescUpdatedAt is the schema descriptor for updated_at field.
	connectionDescUpdatedAt := connectionFields[4].Descriptor()
	// connection.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	connection.DefaultUpdatedAt = connectionDescUpdatedAt.Default.(func() time.Time)
	// connection.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	connection.UpdateDefaultUpdatedAt = connectionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// connectionDescID is the schema descriptor for id field.
	connectionDescID := connectionFields[0].Descriptor()
	// connection.DefaultID holds the default value on creation for the id field.
	connection.DefaultID = connectionDescID.Default.(func() uuid.UUID)
	incomeFields := schema.Income{}.Fields()
	_ = incomeFields
	// incomeDescCreatedAt is the schema descriptor for created_at field.
	incomeDescCreatedAt := incomeFields[3].Descriptor()
	// income.DefaultCreatedAt holds the default value on creation for the created_at field.
	income.DefaultCreatedAt = incomeDescCreatedAt.Default.(func() time.Time)
	// incomeDescUpdatedAt is the schema descriptor for updated_at field.
	incomeDescUpdatedAt := incomeFields[4].Descriptor()
	// income.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	income.DefaultUpdatedAt = incomeDescUpdatedAt.Default.(func() time.Time)
	// income.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	income.UpdateDefaultUpdatedAt = incomeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// incomeDescID is the schema descriptor for id field.
	incomeDescID := incomeFields[0].Descriptor()
	// income.DefaultID holds the default value on creation for the id field.
	income.DefaultID = incomeDescID.Default.(func() uuid.UUID)
	institutionFields := schema.Institution{}.Fields()
	_ = institutionFields
	// institutionDescCreatedAt is the schema descriptor for created_at field.
	institutionDescCreatedAt := institutionFields[3].Descriptor()
	// institution.DefaultCreatedAt holds the default value on creation for the created_at field.
	institution.DefaultCreatedAt = institutionDescCreatedAt.Default.(func() time.Time)
	// institutionDescUpdatedAt is the schema descriptor for updated_at field.
	institutionDescUpdatedAt := institutionFields[4].Descriptor()
	// institution.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	institution.DefaultUpdatedAt = institutionDescUpdatedAt.Default.(func() time.Time)
	// institution.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	institution.UpdateDefaultUpdatedAt = institutionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// institutionDescID is the schema descriptor for id field.
	institutionDescID := institutionFields[0].Descriptor()
	// institution.DefaultID holds the default value on creation for the id field.
	institution.DefaultID = institutionDescID.Default.(func() uuid.UUID)
	manualitemFields := schema.ManualItem{}.Fields()
	_ = manualitemFields
	// manualitemDescProviderName is the schema descriptor for provider_name field.
	manualitemDescProviderName := manualitemFields[1].Descriptor()
	// manualitem.ProviderNameValidator is a validator for the "provider_name" field. It is called by the builders before save.
	manualitem.ProviderNameValidator = manualitemDescProviderName.Validators[0].(func(string) error)
	// manualitemDescCategory is the schema descriptor for category field.
	manualitemDescCategory := manualitemFields[4].Descriptor()
	// manualitem.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	manualitem.CategoryValidator = manualitemDescCategory.Validators[0].(func(string) error)
	// manualitemDescDescription is the schema descriptor for description field.
	manualitemDescDescription := manualitemFields[5].Descriptor()
	// manualitem.DefaultDescription holds the default value on creation for the description field.
	manualitem.DefaultDescription = manualitemDescDescription.Default.(string)
	// manualitemDescCreatedAt is the schema descriptor for created_at field.
	manualitemDescCreatedAt := manualitemFields[7].Descriptor()
	// manualitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	manualitem.DefaultCreatedAt = manualitemDescCreatedAt.Default.(func() time.Time)
	// manualitemDescUpdatedAt is the schema descriptor for updated_at field.
	manualitemDescUpdatedAt := manualitemFields[8].Descriptor()
	// manualitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	manualitem.DefaultUpdatedAt = manualitemDescUpdatedAt.Default.(func() time.Time)
	// manualitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	manualitem.UpdateDefaultUpdatedAt = manualitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// manualitemDescID is the schema descriptor for id field.
	manualitemDescID := manualitemFields[0].Descriptor()
	// manualitem.DefaultID holds the default value on creation for the id field.
	manualitem.DefaultID = manualitemDescID.Default.(func() uuid.UUID)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescListId is the schema descriptor for listId field.
	todoDescListId := todoFields[1].Descriptor()
	// todo.DefaultListId holds the default value on creation for the listId field.
	todo.DefaultListId = todoDescListId.Default.(int)
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[2].Descriptor()
	// todo.DefaultText holds the default value on creation for the text field.
	todo.DefaultText = todoDescText.Default.(string)
	// todoDescCompleted is the schema descriptor for completed field.
	todoDescCompleted := todoFields[3].Descriptor()
	// todo.DefaultCompleted holds the default value on creation for the completed field.
	todo.DefaultCompleted = todoDescCompleted.Default.(bool)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[4].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[5].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	// todo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	todo.UpdateDefaultUpdatedAt = todoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() uuid.UUID)
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescCreatedAt is the schema descriptor for created_at field.
	transactionDescCreatedAt := transactionFields[3].Descriptor()
	// transaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	transaction.DefaultCreatedAt = transactionDescCreatedAt.Default.(func() time.Time)
	// transactionDescUpdatedAt is the schema descriptor for updated_at field.
	transactionDescUpdatedAt := transactionFields[4].Descriptor()
	// transaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	transaction.DefaultUpdatedAt = transactionDescUpdatedAt.Default.(func() time.Time)
	// transaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	transaction.UpdateDefaultUpdatedAt = transactionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// transactionDescID is the schema descriptor for id field.
	transactionDescID := transactionFields[0].Descriptor()
	// transaction.DefaultID holds the default value on creation for the id field.
	transaction.DefaultID = transactionDescID.Default.(func() uuid.UUID)
}
