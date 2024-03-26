table "audit_logs" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "type" {
    null = false
    type = character_varying(255)
  }
  column "error" {
    null = true
    type = character_varying(255)
  }
  column "meta_http_request_id" {
    null = false
    type = character_varying(255)
  }
  column "meta_source_ip" {
    null = false
    type = character_varying(255)
  }
  column "meta_user_agent" {
    null = false
    type = character_varying(255)
  }
  column "actor_user_id" {
    null = true
    type = uuid
  }
  column "actor_email" {
    null = true
    type = character_varying(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  index "audit_logs_actor_email_idx" {
    columns = [column.actor_email]
  }
  index "audit_logs_actor_user_id_idx" {
    columns = [column.actor_user_id]
  }
  index "audit_logs_meta_http_request_id_idx" {
    columns = [column.meta_http_request_id]
  }
  index "audit_logs_meta_source_ip_idx" {
    columns = [column.meta_source_ip]
  }
  index "audit_logs_type_idx" {
    columns = [column.type]
  }
}
table "emails" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "user_id" {
    null = true
    type = uuid
  }
  column "address" {
    null = false
    type = character_varying(255)
  }
  column "verified" {
    null = false
    type = boolean
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "emails_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "emails_address_idx" {
    unique  = true
    columns = [column.address]
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
    type = character_varying(255)
  }
  column "provider_name" {
    null = false
    type = character_varying(255)
  }
  column "data" {
    null = true
    type = text
  }
  column "email_id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "identities_email_id_fkey" {
    columns     = [column.email_id]
    ref_columns = [table.emails.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "identities_provider_id_provider_name_idx" {
    unique  = true
    columns = [column.provider_id, column.provider_name]
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
    type = text
  }
  column "created_at" {
    null = false
    type = timestamp
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
  column "user_id" {
    null = false
    type = uuid
  }
  column "ttl" {
    null = false
    type = integer
  }
  column "code" {
    null = false
    type = character_varying(255)
  }
  column "try_count" {
    null = false
    type = integer
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "email_id" {
    null = true
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "passcodes_emails_id_fk" {
    columns     = [column.email_id]
    ref_columns = [table.emails.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "passcodes_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
table "password_credentials" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "password" {
    null = false
    type = character_varying(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "password_credentials_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "password_credentials_user_id_idx" {
    unique  = true
    columns = [column.user_id]
  }
}
table "primary_emails" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "email_id" {
    null = false
    type = uuid
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "primary_emails_email_id_fkey" {
    columns     = [column.email_id]
    ref_columns = [table.emails.column.id]
    on_update   = CASCADE
    on_delete   = RESTRICT
  }
  foreign_key "primary_emails_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "primary_emails_email_id_idx" {
    unique  = true
    columns = [column.email_id]
  }
  index "primary_emails_user_id_idx" {
    unique  = true
    columns = [column.user_id]
  }
}
table "schema_migration" {
  schema = schema.public
  column "version" {
    null = false
    type = character_varying(14)
  }
  primary_key {
    columns = [column.version]
  }
  index "schema_migration_version_idx" {
    unique  = true
    columns = [column.version]
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
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
}
table "webauthn_credential_transports" {
  schema = schema.public
  column "id" {
    null = false
    type = character_varying(255)
  }
  column "name" {
    null = false
    type = character_varying(255)
  }
  column "webauthn_credential_id" {
    null = false
    type = character_varying(255)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "webauthn_credential_transports_webauthn_credential_id_fkey" {
    columns     = [column.webauthn_credential_id]
    ref_columns = [table.webauthn_credentials.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "webauthn_credential_transports_name_webauthn_credential_id_idx" {
    unique  = true
    columns = [column.name, column.webauthn_credential_id]
  }
}
table "webauthn_credentials" {
  schema = schema.public
  column "id" {
    null = false
    type = character_varying(255)
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "public_key" {
    null = false
    type = text
  }
  column "attestation_type" {
    null = false
    type = character_varying(255)
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
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  column "name" {
    null = true
    type = character_varying(255)
  }
  column "backup_eligible" {
    null    = false
    type    = boolean
    default = false
  }
  column "backup_state" {
    null    = false
    type    = boolean
    default = false
  }
  column "last_used_at" {
    null = true
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "webauthn_credentials_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
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
    type = character_varying(255)
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "user_verification" {
    null = false
    type = character_varying(255)
  }
  column "operation" {
    null = false
    type = character_varying(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  index "webauthn_session_data_challenge_idx" {
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
    type = character_varying(255)
  }
  column "webauthn_session_data_id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "updated_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "webauthn_session_data_allowed_cre_webauthn_session_data_id_fkey" {
    columns     = [column.webauthn_session_data_id]
    ref_columns = [table.webauthn_session_data.column.id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
schema "public" {
}
