-- Create "jwks" table
CREATE TABLE "jwks" ("id" serial NOT NULL, "key_data" character varying NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "emails" table
CREATE TABLE "emails" ("id" uuid NOT NULL, "address" character varying NOT NULL, "verified" boolean NOT NULL DEFAULT false, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "user_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "emails_users_emails" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "emails_address_key" to table: "emails"
CREATE UNIQUE INDEX "emails_address_key" ON "emails" ("address");
-- Create "primary_emails" table
CREATE TABLE "primary_emails" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "email_id" uuid NULL, "user_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "primary_emails_emails_primary_email" FOREIGN KEY ("email_id") REFERENCES "emails" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "primary_emails_users_primary_email" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "primary_emails_email_id_key" to table: "primary_emails"
CREATE UNIQUE INDEX "primary_emails_email_id_key" ON "primary_emails" ("email_id");
-- Create index "primary_emails_user_id_key" to table: "primary_emails"
CREATE UNIQUE INDEX "primary_emails_user_id_key" ON "primary_emails" ("user_id");
-- Create "webauthn_credentials" table
CREATE TABLE "webauthn_credentials" ("id" character varying NOT NULL, "public_key" character varying NOT NULL, "attestation_type" character varying NOT NULL, "aaguid" uuid NOT NULL, "sign_count" integer NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NULL, "backup_eligible" boolean NOT NULL, "backup_state" boolean NOT NULL, "last_used_at" timestamptz NULL, "user_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "webauthn_credentials_users_webauthn_credentials" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "webauthn_credential_transports" table
CREATE TABLE "webauthn_credential_transports" ("id" uuid NOT NULL, "name" character varying NOT NULL, "webauthn_credential_id" character varying NULL, PRIMARY KEY ("id"), CONSTRAINT "webauthn_credential_transports_webauthn_credentials_webauthn_cr" FOREIGN KEY ("webauthn_credential_id") REFERENCES "webauthn_credentials" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "webauthn_session_data" table
CREATE TABLE "webauthn_session_data" ("id" uuid NOT NULL, "challenge" character varying NOT NULL, "user_id" uuid NOT NULL, "user_verification" character varying NOT NULL, "operation" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "webauthn_session_data_challenge_key" to table: "webauthn_session_data"
CREATE UNIQUE INDEX "webauthn_session_data_challenge_key" ON "webauthn_session_data" ("challenge");
-- Create "webauthn_session_data_allowed_credentials" table
CREATE TABLE "webauthn_session_data_allowed_credentials" ("id" uuid NOT NULL, "credential_id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "webauthn_session_data_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "webauthn_session_data_allowed_credentials_webauthn_session_data" FOREIGN KEY ("webauthn_session_data_id") REFERENCES "webauthn_session_data" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "identities" table
CREATE TABLE "identities" ("id" uuid NOT NULL, "provider_id" character varying NOT NULL, "provider_name" character varying NOT NULL, "data" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "email_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "identities_emails_identities" FOREIGN KEY ("email_id") REFERENCES "emails" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "passcodes" table
CREATE TABLE "passcodes" ("id" uuid NOT NULL, "ttl" integer NOT NULL, "code" character varying NOT NULL, "try_count" integer NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "email_id" uuid NULL, "user_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "passcodes_emails_passcodes" FOREIGN KEY ("email_id") REFERENCES "emails" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "passcodes_users_passcodes" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
