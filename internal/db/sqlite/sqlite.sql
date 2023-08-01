-- Create "accounts" table
CREATE TABLE `accounts` (`id` uuid NOT NULL, `provider_name` text NOT NULL DEFAULT '', `data` text NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create "connections" table
CREATE TABLE `connections` (`id` uuid NOT NULL, `provider_name` text NOT NULL DEFAULT '', `data` text NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create "incomes" table
CREATE TABLE `incomes` (`id` uuid NOT NULL, `provider_name` text NOT NULL DEFAULT '', `data` text NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create "institutions" table
CREATE TABLE `institutions` (`id` uuid NOT NULL, `provider_name` text NOT NULL DEFAULT '', `data` text NOT NULL DEFAULT '', PRIMARY KEY (`id`));
-- Create "manual_items" table
CREATE TABLE `manual_items` (`id` uuid NOT NULL, `provider_name` text NOT NULL DEFAULT 'manual', `item_table_id` text NOT NULL DEFAULT 'asset', `type` text NOT NULL DEFAULT 'cash', `category` text NOT NULL DEFAULT 'asset', `description` text NULL DEFAULT '', `value` real NOT NULL DEFAULT 1.401298464324817e-45, PRIMARY KEY (`id`));
-- Create "todos" table
CREATE TABLE `todos` (`id` uuid NOT NULL, `list_id` integer NOT NULL DEFAULT 0, `text` text NOT NULL DEFAULT '', `completed` bool NOT NULL DEFAULT false, PRIMARY KEY (`id`));
-- Create "transactions" table
CREATE TABLE `transactions` (`id` uuid NOT NULL, `provider_name` text NOT NULL DEFAULT '', `data` text NOT NULL DEFAULT '', PRIMARY KEY (`id`));
