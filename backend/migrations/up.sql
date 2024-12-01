-- Таблица пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Таблица категорий
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    name VARCHAR(255) NOT NULL,
    UNIQUE (user_id, name)
);

-- Таблица транзакций
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    amount BIGINT NOT NULL,
    category_id INTEGER REFERENCES categories(id),
    transaction_date DATE NOT NULL DEFAULT CURRENT_DATE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('income', 'expense'))
);

-- Индексы для повышения производительности
CREATE INDEX idx_transactions_user_id ON transactions (user_id);
CREATE INDEX idx_transactions_transaction_date ON transactions (transaction_date);
CREATE INDEX idx_transactions_category_id ON transactions (category_id);
CREATE INDEX idx_categories_user_id ON categories (user_id);

-- Триггер для создания стандартных категорий
CREATE OR REPLACE FUNCTION create_default_categories()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO categories (user_id, name) VALUES
    (NEW.id, 'продукты'),
    (NEW.id, 'транспорт'),
    (NEW.id, 'развлечения');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_categories_trigger
AFTER INSERT ON users
FOR EACH ROW
EXECUTE PROCEDURE create_default_categories();