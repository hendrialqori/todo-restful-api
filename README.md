# Todo RESTful API - Golang
A simple yet production-ready RESTful API built with **Golang**, **HttpRouter**, **MySQL**, and **JWT authentication**.  
The project includes automatic database migrations and Swagger-powered API documentation.

## 🚀 Tech Stack

| Layer          | Technology |
|----------------|------------|
| Language       | Go 1.24 +  |
| Router         | [httprouter](https://github.com/julienschmidt/httprouter) |
| Database       | MySQL |
| Migrations     | [golang-migrate](https://github.com/golang-migrate/migrate) |
| Auth           | [golang-jwt](https://github.com/golang-jwt/jwt) |
| Docs           | [swaggo/swag](https://github.com/swaggo/swag) (Swagger UI) |

---

## 📦 Features

- **Role, User, Todo** full CRUD
- Centralized middleware: JWT auth
- MySQL migrations (`migrate up/down`)
- Swagger UI with live “Authorize” form

---

## 🔐 JWT Authenticate

```Authorization: Bearer <token>```

---

## 🚀 Endpoint list API
You can access **/swagger** at your route to see all existing endpoints.

![Swagger Preview](assets/swagger.png)

---

## 🚀 Command prompt
Run the command using [Makefile](https://makefiletutorial.com/) for details see the table below

| Command               | Description |
|-----------------------|------------|
| make go               | Run go server |
| make swag             | Generate swagger docs |
| make swag-format      | Formater swagger tag |
| make migrate-up       | db migrations up |
| make migrate-down     | db migrations down |
| make migrate-force    | force version for migration |
| make migrate-down     | drop all migrations |

---


MIT License.  
Please feel free to use for personal and commercial purposes.

