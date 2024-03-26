table "emails" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "address" {
    null = false
    type = character_varying
  }
  column "verified" {
    null    = false
    type    = boolean
    default = false
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "user_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "emails_users_emails" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  index "emails_address_key" {
    unique  = true
    columns = [column.address]
  }
}
table "ent_types" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "type" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  index "ent_types_type_key" {
    unique  = true
    columns = [column.type]
  }
}
table "identities" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "provider_id" {
    null = false
    type = character_varying
  }
  column "provider_name" {
    null = false
    type = character_varying
  }
  column "data" {
    null = true
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "email_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "identities_emails_identities" {
    columns     = [column.email_id]
    ref_columns = [table.emails.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
}
table "jwks" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "key_data" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
}
table "passcodes" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "ttl" {
    null = false
    type = integer
  }
  column "code" {
    null = false
    type = character_varying
  }
  column "try_count" {
    null = false
    type = integer
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "email_id" {
    null = true
    type = uuid
  }
  column "user_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "passcodes_emails_passcodes" {
    columns     = [column.email_id]
    ref_columns = [table.emails.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  foreign_key "passcodes_users_passcodes" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
}
table "primary_emails" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "email_id" {
    null = true
    type = uuid
  }
  column "user_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "primary_emails_emails_primary_email" {
    columns     = [column.email_id]
    ref_columns = [table.emails.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  foreign_key "primary_emails_users_primary_email" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
  index "primary_emails_email_id_key" {
    unique  = true
    columns = [column.email_id]
  }
  index "primary_emails_user_id_key" {
    unique  = true
    columns = [column.user_id]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
}
table "webauthn_credential_transports" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "name" {
    null = false
    type = character_varying
  }
  column "webauthn_credential_id" {
    null = true
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "webauthn_credential_transports_27503b7d5ccd2749f322d55e8d2d52ad" {
    columns     = [column.webauthn_credential_id]
    ref_columns = [table.webauthn_credentials.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
}
table "webauthn_credentials" {
  schema = schema.public
  column "id" {
    null = false
    type = character_varying
  }
  column "public_key" {
    null = false
    type = character_varying
  }
  column "attestation_type" {
    null = false
    type = character_varying
  }
  column "aaguid" {
    null = false
    type = uuid
  }
  column "sign_count" {
    null = false
    type = integer
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "name" {
    null = true
    type = character_varying
  }
  column "backup_eligible" {
    null = false
    type = boolean
  }
  column "backup_state" {
    null = false
    type = boolean
  }
  column "last_used_at" {
    null = true
    type = timestamptz
  }
  column "user_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "webauthn_credentials_users_webauthn_credentials" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
}
table "webauthn_session_data" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "challenge" {
    null = false
    type = character_varying
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "user_verification" {
    null = false
    type = character_varying
  }
  column "operation" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  index "webauthn_session_data_challenge_key" {
    unique  = true
    columns = [column.challenge]
  }
}
table "webauthn_session_data_allowed_credentials" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "credential_id" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "webauthn_session_data_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "webauthn_session_data_allowed__ddb38119f6377d2308ddbc520864318f" {
    columns     = [column.webauthn_session_data_id]
    ref_columns = [table.webauthn_session_data.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
}
schema "public" {
}
