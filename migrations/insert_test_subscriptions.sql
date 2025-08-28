-- Вставка тестовых данных --
INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) VALUES
('Yandex Plus', 199, '11111111-1111-1111-1111-111111111111', NOW() - INTERVAL '30 days', NOW() + INTERVAL '30 days'),
('Netflix', 899, '22222222-2222-2222-2222-222222222222', NOW() - INTERVAL '15 days', NOW() + INTERVAL '15 days'),
('Spotify', 499, '33333333-3333-3333-3333-333333333333', NOW() - INTERVAL '60 days', NULL),
('IVI', 399, '44444444-4444-4444-4444-444444444444', NOW(), NOW() + INTERVAL '1 month'),
('Okko', 299, '55555555-5555-5555-5555-555555555555', NOW() - INTERVAL '10 days', NOW() + INTERVAL '20 days');
