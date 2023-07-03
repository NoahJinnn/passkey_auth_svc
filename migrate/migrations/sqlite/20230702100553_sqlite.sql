-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_providers" table
CREATE TABLE `new_providers` (`id` uuid NOT NULL, `user_id` uuid NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, PRIMARY KEY (`id`));
-- Copy rows from old table "providers" to new temporary table "new_providers"
INSERT INTO `new_providers` (`id`, `user_id`, `created_at`, `updated_at`) SELECT `id`, `user_id`, `created_at`, `updated_at` FROM `providers`;
-- Drop "providers" table after copying rows
DROP TABLE `providers`;
-- Rename temporary table "new_providers" to "providers"
ALTER TABLE `new_providers` RENAME TO `providers`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
