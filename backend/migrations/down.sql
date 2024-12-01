DROP TRIGGER IF EXISTS create_categories_trigger ON users;
DROP FUNCTION IF EXISTS create_default_categories;

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_transactions_user_id;
DROP INDEX IF EXISTS idx_transactions_transaction_date;
DROP INDEX IF EXISTS idx_transactions_category_id;
DROP INDEX IF EXISTS idx_categories_user_id;