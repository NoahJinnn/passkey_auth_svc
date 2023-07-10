// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/income"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/institution"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/transaction"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Connection is the client for interacting with the Connection builders.
	Connection *ConnectionClient
	// Income is the client for interacting with the Income builders.
	Income *IncomeClient
	// Institution is the client for interacting with the Institution builders.
	Institution *InstitutionClient
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Connection = NewConnectionClient(c.config)
	c.Income = NewIncomeClient(c.config)
	c.Institution = NewInstitutionClient(c.config)
	c.Transaction = NewTransactionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Account:     NewAccountClient(cfg),
		Connection:  NewConnectionClient(cfg),
		Income:      NewIncomeClient(cfg),
		Institution: NewInstitutionClient(cfg),
		Transaction: NewTransactionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Account:     NewAccountClient(cfg),
		Connection:  NewConnectionClient(cfg),
		Income:      NewIncomeClient(cfg),
		Institution: NewInstitutionClient(cfg),
		Transaction: NewTransactionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Connection.Use(hooks...)
	c.Income.Use(hooks...)
	c.Institution.Use(hooks...)
	c.Transaction.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Account.Intercept(interceptors...)
	c.Connection.Intercept(interceptors...)
	c.Income.Intercept(interceptors...)
	c.Institution.Intercept(interceptors...)
	c.Transaction.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AccountMutation:
		return c.Account.mutate(ctx, m)
	case *ConnectionMutation:
		return c.Connection.mutate(ctx, m)
	case *IncomeMutation:
		return c.Income.mutate(ctx, m)
	case *InstitutionMutation:
		return c.Institution.mutate(ctx, m)
	case *TransactionMutation:
		return c.Transaction.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `account.Intercept(f(g(h())))`.
func (c *AccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.Account = append(c.inters.Account, interceptors...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uuid.UUID) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uuid.UUID) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uuid.UUID) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uuid.UUID) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstitution queries the institution edge of a Account.
func (c *AccountClient) QueryInstitution(a *Account) *InstitutionQuery {
	query := (&InstitutionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(institution.Table, institution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, account.InstitutionTable, account.InstitutionColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTransactions queries the transactions edge of a Account.
func (c *AccountClient) QueryTransactions(a *Account) *TransactionQuery {
	query := (&TransactionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.TransactionsTable, account.TransactionsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// Interceptors returns the client interceptors.
func (c *AccountClient) Interceptors() []Interceptor {
	return c.inters.Account
}

func (c *AccountClient) mutate(ctx context.Context, m *AccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Account mutation op: %q", m.Op())
	}
}

// ConnectionClient is a client for the Connection schema.
type ConnectionClient struct {
	config
}

// NewConnectionClient returns a client for the Connection from the given config.
func NewConnectionClient(c config) *ConnectionClient {
	return &ConnectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `connection.Hooks(f(g(h())))`.
func (c *ConnectionClient) Use(hooks ...Hook) {
	c.hooks.Connection = append(c.hooks.Connection, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `connection.Intercept(f(g(h())))`.
func (c *ConnectionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Connection = append(c.inters.Connection, interceptors...)
}

// Create returns a builder for creating a Connection entity.
func (c *ConnectionClient) Create() *ConnectionCreate {
	mutation := newConnectionMutation(c.config, OpCreate)
	return &ConnectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Connection entities.
func (c *ConnectionClient) CreateBulk(builders ...*ConnectionCreate) *ConnectionCreateBulk {
	return &ConnectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Connection.
func (c *ConnectionClient) Update() *ConnectionUpdate {
	mutation := newConnectionMutation(c.config, OpUpdate)
	return &ConnectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConnectionClient) UpdateOne(co *Connection) *ConnectionUpdateOne {
	mutation := newConnectionMutation(c.config, OpUpdateOne, withConnection(co))
	return &ConnectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConnectionClient) UpdateOneID(id uuid.UUID) *ConnectionUpdateOne {
	mutation := newConnectionMutation(c.config, OpUpdateOne, withConnectionID(id))
	return &ConnectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Connection.
func (c *ConnectionClient) Delete() *ConnectionDelete {
	mutation := newConnectionMutation(c.config, OpDelete)
	return &ConnectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ConnectionClient) DeleteOne(co *Connection) *ConnectionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ConnectionClient) DeleteOneID(id uuid.UUID) *ConnectionDeleteOne {
	builder := c.Delete().Where(connection.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConnectionDeleteOne{builder}
}

// Query returns a query builder for Connection.
func (c *ConnectionClient) Query() *ConnectionQuery {
	return &ConnectionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeConnection},
		inters: c.Interceptors(),
	}
}

// Get returns a Connection entity by its id.
func (c *ConnectionClient) Get(ctx context.Context, id uuid.UUID) (*Connection, error) {
	return c.Query().Where(connection.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConnectionClient) GetX(ctx context.Context, id uuid.UUID) *Connection {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstitution queries the institution edge of a Connection.
func (c *ConnectionClient) QueryInstitution(co *Connection) *InstitutionQuery {
	query := (&InstitutionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(connection.Table, connection.FieldID, id),
			sqlgraph.To(institution.Table, institution.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, connection.InstitutionTable, connection.InstitutionColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConnectionClient) Hooks() []Hook {
	return c.hooks.Connection
}

// Interceptors returns the client interceptors.
func (c *ConnectionClient) Interceptors() []Interceptor {
	return c.inters.Connection
}

func (c *ConnectionClient) mutate(ctx context.Context, m *ConnectionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ConnectionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ConnectionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ConnectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ConnectionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Connection mutation op: %q", m.Op())
	}
}

// IncomeClient is a client for the Income schema.
type IncomeClient struct {
	config
}

// NewIncomeClient returns a client for the Income from the given config.
func NewIncomeClient(c config) *IncomeClient {
	return &IncomeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `income.Hooks(f(g(h())))`.
func (c *IncomeClient) Use(hooks ...Hook) {
	c.hooks.Income = append(c.hooks.Income, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `income.Intercept(f(g(h())))`.
func (c *IncomeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Income = append(c.inters.Income, interceptors...)
}

// Create returns a builder for creating a Income entity.
func (c *IncomeClient) Create() *IncomeCreate {
	mutation := newIncomeMutation(c.config, OpCreate)
	return &IncomeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Income entities.
func (c *IncomeClient) CreateBulk(builders ...*IncomeCreate) *IncomeCreateBulk {
	return &IncomeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Income.
func (c *IncomeClient) Update() *IncomeUpdate {
	mutation := newIncomeMutation(c.config, OpUpdate)
	return &IncomeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IncomeClient) UpdateOne(i *Income) *IncomeUpdateOne {
	mutation := newIncomeMutation(c.config, OpUpdateOne, withIncome(i))
	return &IncomeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IncomeClient) UpdateOneID(id uuid.UUID) *IncomeUpdateOne {
	mutation := newIncomeMutation(c.config, OpUpdateOne, withIncomeID(id))
	return &IncomeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Income.
func (c *IncomeClient) Delete() *IncomeDelete {
	mutation := newIncomeMutation(c.config, OpDelete)
	return &IncomeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IncomeClient) DeleteOne(i *Income) *IncomeDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IncomeClient) DeleteOneID(id uuid.UUID) *IncomeDeleteOne {
	builder := c.Delete().Where(income.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IncomeDeleteOne{builder}
}

// Query returns a query builder for Income.
func (c *IncomeClient) Query() *IncomeQuery {
	return &IncomeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIncome},
		inters: c.Interceptors(),
	}
}

// Get returns a Income entity by its id.
func (c *IncomeClient) Get(ctx context.Context, id uuid.UUID) (*Income, error) {
	return c.Query().Where(income.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IncomeClient) GetX(ctx context.Context, id uuid.UUID) *Income {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstitution queries the institution edge of a Income.
func (c *IncomeClient) QueryInstitution(i *Income) *InstitutionQuery {
	query := (&InstitutionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(income.Table, income.FieldID, id),
			sqlgraph.To(institution.Table, institution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, income.InstitutionTable, income.InstitutionColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IncomeClient) Hooks() []Hook {
	return c.hooks.Income
}

// Interceptors returns the client interceptors.
func (c *IncomeClient) Interceptors() []Interceptor {
	return c.inters.Income
}

func (c *IncomeClient) mutate(ctx context.Context, m *IncomeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IncomeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IncomeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IncomeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IncomeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Income mutation op: %q", m.Op())
	}
}

// InstitutionClient is a client for the Institution schema.
type InstitutionClient struct {
	config
}

// NewInstitutionClient returns a client for the Institution from the given config.
func NewInstitutionClient(c config) *InstitutionClient {
	return &InstitutionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `institution.Hooks(f(g(h())))`.
func (c *InstitutionClient) Use(hooks ...Hook) {
	c.hooks.Institution = append(c.hooks.Institution, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `institution.Intercept(f(g(h())))`.
func (c *InstitutionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Institution = append(c.inters.Institution, interceptors...)
}

// Create returns a builder for creating a Institution entity.
func (c *InstitutionClient) Create() *InstitutionCreate {
	mutation := newInstitutionMutation(c.config, OpCreate)
	return &InstitutionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Institution entities.
func (c *InstitutionClient) CreateBulk(builders ...*InstitutionCreate) *InstitutionCreateBulk {
	return &InstitutionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Institution.
func (c *InstitutionClient) Update() *InstitutionUpdate {
	mutation := newInstitutionMutation(c.config, OpUpdate)
	return &InstitutionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstitutionClient) UpdateOne(i *Institution) *InstitutionUpdateOne {
	mutation := newInstitutionMutation(c.config, OpUpdateOne, withInstitution(i))
	return &InstitutionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstitutionClient) UpdateOneID(id uuid.UUID) *InstitutionUpdateOne {
	mutation := newInstitutionMutation(c.config, OpUpdateOne, withInstitutionID(id))
	return &InstitutionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Institution.
func (c *InstitutionClient) Delete() *InstitutionDelete {
	mutation := newInstitutionMutation(c.config, OpDelete)
	return &InstitutionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InstitutionClient) DeleteOne(i *Institution) *InstitutionDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *InstitutionClient) DeleteOneID(id uuid.UUID) *InstitutionDeleteOne {
	builder := c.Delete().Where(institution.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstitutionDeleteOne{builder}
}

// Query returns a query builder for Institution.
func (c *InstitutionClient) Query() *InstitutionQuery {
	return &InstitutionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeInstitution},
		inters: c.Interceptors(),
	}
}

// Get returns a Institution entity by its id.
func (c *InstitutionClient) Get(ctx context.Context, id uuid.UUID) (*Institution, error) {
	return c.Query().Where(institution.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstitutionClient) GetX(ctx context.Context, id uuid.UUID) *Institution {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryConnection queries the connection edge of a Institution.
func (c *InstitutionClient) QueryConnection(i *Institution) *ConnectionQuery {
	query := (&ConnectionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(institution.Table, institution.FieldID, id),
			sqlgraph.To(connection.Table, connection.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, institution.ConnectionTable, institution.ConnectionColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccounts queries the accounts edge of a Institution.
func (c *InstitutionClient) QueryAccounts(i *Institution) *AccountQuery {
	query := (&AccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(institution.Table, institution.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, institution.AccountsTable, institution.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIncomes queries the incomes edge of a Institution.
func (c *InstitutionClient) QueryIncomes(i *Institution) *IncomeQuery {
	query := (&IncomeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(institution.Table, institution.FieldID, id),
			sqlgraph.To(income.Table, income.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, institution.IncomesTable, institution.IncomesColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstitutionClient) Hooks() []Hook {
	return c.hooks.Institution
}

// Interceptors returns the client interceptors.
func (c *InstitutionClient) Interceptors() []Interceptor {
	return c.inters.Institution
}

func (c *InstitutionClient) mutate(ctx context.Context, m *InstitutionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&InstitutionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&InstitutionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&InstitutionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&InstitutionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Institution mutation op: %q", m.Op())
	}
}

// TransactionClient is a client for the Transaction schema.
type TransactionClient struct {
	config
}

// NewTransactionClient returns a client for the Transaction from the given config.
func NewTransactionClient(c config) *TransactionClient {
	return &TransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transaction.Hooks(f(g(h())))`.
func (c *TransactionClient) Use(hooks ...Hook) {
	c.hooks.Transaction = append(c.hooks.Transaction, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `transaction.Intercept(f(g(h())))`.
func (c *TransactionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Transaction = append(c.inters.Transaction, interceptors...)
}

// Create returns a builder for creating a Transaction entity.
func (c *TransactionClient) Create() *TransactionCreate {
	mutation := newTransactionMutation(c.config, OpCreate)
	return &TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transaction entities.
func (c *TransactionClient) CreateBulk(builders ...*TransactionCreate) *TransactionCreateBulk {
	return &TransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transaction.
func (c *TransactionClient) Update() *TransactionUpdate {
	mutation := newTransactionMutation(c.config, OpUpdate)
	return &TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionClient) UpdateOne(t *Transaction) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransaction(t))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionClient) UpdateOneID(id uuid.UUID) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransactionID(id))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transaction.
func (c *TransactionClient) Delete() *TransactionDelete {
	mutation := newTransactionMutation(c.config, OpDelete)
	return &TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionClient) DeleteOne(t *Transaction) *TransactionDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TransactionClient) DeleteOneID(id uuid.UUID) *TransactionDeleteOne {
	builder := c.Delete().Where(transaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionDeleteOne{builder}
}

// Query returns a query builder for Transaction.
func (c *TransactionClient) Query() *TransactionQuery {
	return &TransactionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTransaction},
		inters: c.Interceptors(),
	}
}

// Get returns a Transaction entity by its id.
func (c *TransactionClient) Get(ctx context.Context, id uuid.UUID) (*Transaction, error) {
	return c.Query().Where(transaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionClient) GetX(ctx context.Context, id uuid.UUID) *Transaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a Transaction.
func (c *TransactionClient) QueryAccount(t *Transaction) *AccountQuery {
	query := (&AccountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transaction.AccountTable, transaction.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TransactionClient) Hooks() []Hook {
	return c.hooks.Transaction
}

// Interceptors returns the client interceptors.
func (c *TransactionClient) Interceptors() []Interceptor {
	return c.inters.Transaction
}

func (c *TransactionClient) mutate(ctx context.Context, m *TransactionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Transaction mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Account, Connection, Income, Institution, Transaction []ent.Hook
	}
	inters struct {
		Account, Connection, Income, Institution, Transaction []ent.Interceptor
	}
)
