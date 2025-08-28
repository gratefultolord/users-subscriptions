-- Убрать перед отправкой --
DROP TABLE IF EXISTS subscriptions;

-- Таблица для хранения подписок --
CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    user_id UUID NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP
);