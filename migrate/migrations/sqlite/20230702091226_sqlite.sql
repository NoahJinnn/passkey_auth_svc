-- Create "providers" table
CREATE TABLE `providers` (`id` uuid NOT NULL, `user_id` uuid NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, PRIMARY KEY (`id`));
