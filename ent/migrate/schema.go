// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssetTablesColumns holds the columns for the "asset_tables" table.
	AssetTablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sheet", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "section", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// AssetTablesTable holds the schema information for the "asset_tables" table.
	AssetTablesTable = &schema.Table{
		Name:       "asset_tables",
		Columns:    AssetTablesColumns,
		PrimaryKey: []*schema.Column{AssetTablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "asset_tables_users_asset_tables",
				Columns:    []*schema.Column{AssetTablesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmailsColumns holds the columns for the "emails" table.
	EmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// EmailsTable holds the schema information for the "emails" table.
	EmailsTable = &schema.Table{
		Name:       "emails",
		Columns:    EmailsColumns,
		PrimaryKey: []*schema.Column{EmailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "emails_users_emails",
				Columns:    []*schema.Column{EmailsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FvSessionsColumns holds the columns for the "fv_sessions" table.
	FvSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "access_token", Type: field.TypeString},
		{Name: "expires_in", Type: field.TypeInt32},
		{Name: "issued_at", Type: field.TypeString},
		{Name: "token_type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// FvSessionsTable holds the schema information for the "fv_sessions" table.
	FvSessionsTable = &schema.Table{
		Name:       "fv_sessions",
		Columns:    FvSessionsColumns,
		PrimaryKey: []*schema.Column{FvSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "fv_sessions_users_fv_session",
				Columns:    []*schema.Column{FvSessionsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// IdentitiesColumns holds the columns for the "identities" table.
	IdentitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider_id", Type: field.TypeString},
		{Name: "provider_name", Type: field.TypeString},
		{Name: "data", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email_id", Type: field.TypeUUID, Nullable: true},
	}
	// IdentitiesTable holds the schema information for the "identities" table.
	IdentitiesTable = &schema.Table{
		Name:       "identities",
		Columns:    IdentitiesColumns,
		PrimaryKey: []*schema.Column{IdentitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "identities_emails_identities",
				Columns:    []*schema.Column{IdentitiesColumns[6]},
				RefColumns: []*schema.Column{EmailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JwksColumns holds the columns for the "jwks" table.
	JwksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true, SchemaType: map[string]string{"postgres": "serial"}},
		{Name: "key_data", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// JwksTable holds the schema information for the "jwks" table.
	JwksTable = &schema.Table{
		Name:       "jwks",
		Columns:    JwksColumns,
		PrimaryKey: []*schema.Column{JwksColumns[0]},
	}
	// PasscodesColumns holds the columns for the "passcodes" table.
	PasscodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "ttl", Type: field.TypeInt32},
		{Name: "code", Type: field.TypeString},
		{Name: "try_count", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// PasscodesTable holds the schema information for the "passcodes" table.
	PasscodesTable = &schema.Table{
		Name:       "passcodes",
		Columns:    PasscodesColumns,
		PrimaryKey: []*schema.Column{PasscodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "passcodes_emails_passcodes",
				Columns:    []*schema.Column{PasscodesColumns[6]},
				RefColumns: []*schema.Column{EmailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "passcodes_users_passcodes",
				Columns:    []*schema.Column{PasscodesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PrimaryEmailsColumns holds the columns for the "primary_emails" table.
	PrimaryEmailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// PrimaryEmailsTable holds the schema information for the "primary_emails" table.
	PrimaryEmailsTable = &schema.Table{
		Name:       "primary_emails",
		Columns:    PrimaryEmailsColumns,
		PrimaryKey: []*schema.Column{PrimaryEmailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "primary_emails_emails_primary_email",
				Columns:    []*schema.Column{PrimaryEmailsColumns[3]},
				RefColumns: []*schema.Column{EmailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "primary_emails_users_primary_email",
				Columns:    []*schema.Column{PrimaryEmailsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WebauthnCredentialsColumns holds the columns for the "webauthn_credentials" table.
	WebauthnCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "public_key", Type: field.TypeString},
		{Name: "attestation_type", Type: field.TypeString},
		{Name: "aaguid", Type: field.TypeUUID},
		{Name: "sign_count", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "backup_eligible", Type: field.TypeBool},
		{Name: "backup_state", Type: field.TypeBool},
		{Name: "last_used_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// WebauthnCredentialsTable holds the schema information for the "webauthn_credentials" table.
	WebauthnCredentialsTable = &schema.Table{
		Name:       "webauthn_credentials",
		Columns:    WebauthnCredentialsColumns,
		PrimaryKey: []*schema.Column{WebauthnCredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "webauthn_credentials_users_webauthn_credentials",
				Columns:    []*schema.Column{WebauthnCredentialsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WebauthnCredentialTransportsColumns holds the columns for the "webauthn_credential_transports" table.
	WebauthnCredentialTransportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "webauthn_credential_id", Type: field.TypeString, Nullable: true},
	}
	// WebauthnCredentialTransportsTable holds the schema information for the "webauthn_credential_transports" table.
	WebauthnCredentialTransportsTable = &schema.Table{
		Name:       "webauthn_credential_transports",
		Columns:    WebauthnCredentialTransportsColumns,
		PrimaryKey: []*schema.Column{WebauthnCredentialTransportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "webauthn_credential_transports_webauthn_credentials_webauthn_credential_transports",
				Columns:    []*schema.Column{WebauthnCredentialTransportsColumns[2]},
				RefColumns: []*schema.Column{WebauthnCredentialsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WebauthnSessionDataColumns holds the columns for the "webauthn_session_data" table.
	WebauthnSessionDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "challenge", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "user_verification", Type: field.TypeString},
		{Name: "operation", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// WebauthnSessionDataTable holds the schema information for the "webauthn_session_data" table.
	WebauthnSessionDataTable = &schema.Table{
		Name:       "webauthn_session_data",
		Columns:    WebauthnSessionDataColumns,
		PrimaryKey: []*schema.Column{WebauthnSessionDataColumns[0]},
	}
	// WebauthnSessionDataAllowedCredentialsColumns holds the columns for the "webauthn_session_data_allowed_credentials" table.
	WebauthnSessionDataAllowedCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "credential_id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "webauthn_session_data_id", Type: field.TypeUUID, Nullable: true},
	}
	// WebauthnSessionDataAllowedCredentialsTable holds the schema information for the "webauthn_session_data_allowed_credentials" table.
	WebauthnSessionDataAllowedCredentialsTable = &schema.Table{
		Name:       "webauthn_session_data_allowed_credentials",
		Columns:    WebauthnSessionDataAllowedCredentialsColumns,
		PrimaryKey: []*schema.Column{WebauthnSessionDataAllowedCredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "webauthn_session_data_allowed_credentials_webauthn_session_data_webauthn_session_data_allowed_credentials",
				Columns:    []*schema.Column{WebauthnSessionDataAllowedCredentialsColumns[4]},
				RefColumns: []*schema.Column{WebauthnSessionDataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssetTablesTable,
		EmailsTable,
		FvSessionsTable,
		IdentitiesTable,
		JwksTable,
		PasscodesTable,
		PrimaryEmailsTable,
		UsersTable,
		WebauthnCredentialsTable,
		WebauthnCredentialTransportsTable,
		WebauthnSessionDataTable,
		WebauthnSessionDataAllowedCredentialsTable,
	}
)

func init() {
	AssetTablesTable.ForeignKeys[0].RefTable = UsersTable
	EmailsTable.ForeignKeys[0].RefTable = UsersTable
	FvSessionsTable.ForeignKeys[0].RefTable = UsersTable
	IdentitiesTable.ForeignKeys[0].RefTable = EmailsTable
	PasscodesTable.ForeignKeys[0].RefTable = EmailsTable
	PasscodesTable.ForeignKeys[1].RefTable = UsersTable
	PrimaryEmailsTable.ForeignKeys[0].RefTable = EmailsTable
	PrimaryEmailsTable.ForeignKeys[1].RefTable = UsersTable
	WebauthnCredentialsTable.ForeignKeys[0].RefTable = UsersTable
	WebauthnCredentialTransportsTable.ForeignKeys[0].RefTable = WebauthnCredentialsTable
	WebauthnSessionDataAllowedCredentialsTable.ForeignKeys[0].RefTable = WebauthnSessionDataTable
}
