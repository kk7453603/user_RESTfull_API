# Golang Gin clean architecture RESTfull API

**_Stack: gin+mongodb_**

## Правила чистой архитектуры:
- Независимость от фреймворков. Архитектура не зависит от существования какой-либо библиотеки функционального программного обеспечения. Это позволяет вам использовать такие платформы в качестве инструментов, вместо того, чтобы загонять вашу систему в их ограниченные ограничения.
- Тестируемость. Бизнес-правила можно тестировать без пользовательского интерфейса, базы данных, веб-сервера или любого другого внешнего элемента.
- Независимость от пользовательского интерфейса. Пользовательский интерфейс можно легко изменить, не меняя остальную часть системы. Веб-интерфейс можно заменить, например, консольным интерфейсом без изменения бизнес-правил.
- Независимость от базы данных. Вы можете заменить Oracle или SQL Server на Mongo, BigTable, CouchDB или что-то еще. Ваши бизнес-правила не привязаны к базе данных.
- Независимость от каких-либо внешних факторов. На самом деле ваши бизнес-правила просто ничего не знают о внешнем мире.
---
## Описание проекта:
REST API со специальной системой аутентификации на основе JWT. Основная функциональность заключается в создании закладок и управлении ими.

### Структура
4 Базовых слоя:
- Models layer
- Repository layer
- UseCase layer
- Delivery layer

## **API Endpoints:**

**POST /auth/sign-up**<br>
Создает нового пользователя<br>
Пример входных данных:
```json
{
	"username": "Vasya",
	"password": "cleanArch"
} 
```

**POST /auth/sign-in**<br>
Создает запрос на получение токена JWT на основе учетных данных пользователя<br>
Пример входных данных:
```json
{
	"username": "Vasya",
	"password": "cleanArch"
} 
```
Пример ответа:
```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

**POST /api/bookmarks**<br>
Создает новую закладку<br>
Пример входных данных:
```json
{
	"url": "https://go.dev",
	"title": "Go Clean Architecture example"
} 
```

**GET /api/bookmarks**<br>
Возвращает все закладки пользователя<br>
Пример входных данных:
```json
{
	"bookmarks": [
            {
                "id": "5da2d8aae9b63715ddfae856",
                "url": "https://go.dev",
                "title": "Go Clean Architecture example"
            }
    ]
} 
```

**DELETE /api/bookmarks**<br>
Удаляет закладку по ID<br>
Пример входных данных:
```json
{
	"id": "5da2d8aae9b63715ddfae856"
} 
```

# **Запуск проекта**
Исрользуйте ```make run``` для создания и запуска Docker-контейнеров с самим приложением и экземпляром mongodb