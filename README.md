# Сервис по работе с информацией о подписках

REST-сервис для агрегации данных об онлайн-подписках пользователей. Доступны CRUDL-операции с подписками, а также получение суммарной стоимости всех подписок с фильтрацией по пользователю и названию сервиса, предоставляющего подписку.


## Технический стек

Golang, PostgreSQL, Docker, Swagger



## Локальный запуск

Клонировать проект

```bash
  git clone https://github.com/gratefultolord/users-subscriptions.git
```

Перейти в директорию проекта

```bash
  cd users-subscriptions
```

Создать .env файл

```bash
  touch .env
```

Заполнить .env файл

```env
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=sub_db
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
```

Запустить докер

```bash
  docker compose up --build
```


## Тесты
В рамках проекта были написаны интеграционные тесты.

Чтобы запустить введите следующую команду в терминале

```bash
  go test ./...
```


## API

Спецификация API будет доступна после запуска проекта на `http://localhost:8080/index.html`

Внизу некоторые примеры запросов

#### Получить информацию списка подписок

```http
  GET /subscriptions
```

#### Получить информацию о подписке

```http
  GET /susbcriptions/{subscriptionId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `subscriptionId`      | `integer` | **Обязательный параметр**. Id подписки |



## 🚀 Обо мне
Golang-разработчик, работаю в Авито