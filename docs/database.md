## Схема базы данных

### 1. Таблица пользователей

| Поле        | Тип данных                     | Ограничения                                   | Описание                                   |
|-------------|---------------------------------|-----------------------------------------------|--------------------------------------------|
| `id`        | `SERIAL`                        | `PRIMARY KEY`                                 | Уникальный идентификатор пользователя.     |
| `created_at`| `TIMESTAMP WITH TIME ZONE`     | `DEFAULT CURRENT_TIMESTAMP`                   | Дата и время создания записи.             |
| `updated_at`| `TIMESTAMP WITH TIME ZONE`     | `DEFAULT CURRENT_TIMESTAMP`                   | Дата и время последнего обновления записи.|
| `deleted_at`| `TIMESTAMP WITH TIME ZONE`     | -                                             | Дата и время удаления записи (логическое удаление). |
| `name`      | `VARCHAR(255)`                 | `NOT NULL`                                    | Имя пользователя.                          |
| `password`  | `VARCHAR(255)`                 | `NOT NULL`                                    | Пароль пользователя.                       |

### 2. Таблица категорий

| Поле        | Тип данных                     | Ограничения                                   | Описание                                   |
|-------------|---------------------------------|-----------------------------------------------|--------------------------------------------|
| `id`        | `SERIAL`                        | `PRIMARY KEY`                                 | Уникальный идентификатор категории.        |
| `created_at`| `TIMESTAMP WITH TIME ZONE`     | `DEFAULT CURRENT_TIMESTAMP`                   | Дата и время создания записи.             |
| `updated_at`| `TIMESTAMP WITH TIME ZONE`     | `DEFAULT CURRENT_TIMESTAMP`                   | Дата и время последнего обновления записи.|
| `deleted_at`| `TIMESTAMP WITH TIME ZONE`     | -                                             | Дата и время удаления записи (логическое удаление). |
| `user_id`   | `INTEGER`                      | `REFERENCES users(id) NOT NULL`              | Идентификатор пользователя, которому принадлежит категория. |
| `name`      | `VARCHAR(255)`                 | `NOT NULL`, `UNIQUE (user_id, name)`        | Название категории.                        |

### 3. Таблица транзакций

| Поле               | Тип данных                     | Ограничения                                   | Описание                                   |
|--------------------|---------------------------------|-----------------------------------------------|--------------------------------------------|
| `id`               | `SERIAL`                        | `PRIMARY KEY`                                 | Уникальный идентификатор транзакции.      |
| `created_at`       | `TIMESTAMP WITH TIME ZONE`     | `DEFAULT CURRENT_TIMESTAMP`                   | Дата и время создания записи.             |
| `updated_at`       | `TIMESTAMP WITH TIME ZONE`     | `DEFAULT CURRENT_TIMESTAMP`                   | Дата и время последнего обновления записи.|
| `deleted_at`       | `TIMESTAMP WITH TIME ZONE`     | -                                             | Дата и время удаления записи (логическое удаление). |
| `user_id`          | `INTEGER`                      | `REFERENCES users(id) NOT NULL`              | Идентификатор пользователя, которому принадлежит транзакция. |
| `amount`           | `BIGINT`                       | `NOT NULL`                                    | Сумма транзакции.                         |
| `category_id`      | `INTEGER`                      | `REFERENCES categories(id)`                   | Идентификатор категории, к которой относится транзакция. |
| `transaction_date` | `DATE NOT NULL DEFAULT CURRENT_DATE`  | -                                      | Дата транзакции (по умолчанию текущая дата). |
| `type`             | `VARCHAR(10)`                  | `NOT NULL`, CHECK (type IN ('income', 'expense'))  | Тип транзакции: доход или расход. |

### 4. Индексы

- **Индекс по user_id в таблице транзакций**:
  ```sql
  CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions (user_id);
  ```

- **Индекс по transaction_date в таблице транзакций**:
  ```sql
  CREATE INDEX IF NOT EXISTS idx_transactions_transaction_date ON transactions (transaction_date);
  ```

- **Индекс по category_id в таблице транзакций**:
  ```sql
  CREATE INDEX IF NOT EXISTS idx_transactions_category_id ON transactions (category_id);
  ```

- **Индекс по user_id в таблице категорий**:
  ```sql
  CREATE INDEX IF NOT EXISTS idx_categories_user_id ON categories (user_id);
  ```

### 5. Триггер для создания стандартных категорий

#### Функция триггера

```sql
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
```

#### Создание триггера

```sql
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = 'create_categories_trigger') THEN
    CREATE TRIGGER create_categories_trigger
    AFTER INSERT ON users
    FOR EACH ROW
    EXECUTE PROCEDURE create_default_categories();
  END IF;
END;
$$ LANGUAGE plpgsql;
```