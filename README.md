# ✅ TODO CRUD API - Golang

Sebuah RESTful API menggunakan Go dengan fitur:

- CRUD (Create, Read, Update, Delete) untuk Todo
- Autentikasi menggunakan JWT
- 3 entitas utama: **roles**, **users**, **todos**
- Logging menggunakan Logrus

---

## 📦 Stack Teknologi

- `Go (1.21+)`
- `database/sql` + `go-sql-driver/mysql`
- `github.com/golang-jwt/jwt/v5`
- `github.com/julienschmidt/httprouter`
- `github.com/sirupsen/logrus`

---

## 🗃️ Struktur Tabel

### `roles` Table


```.sql
CREATE TABLE roles (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

```

### `users` Table

```
CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(200) NOT NULL,
  username VARCHAR(200) NOT NULL,
  role_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_users_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);
```

### `todos` Table

```CREATE TABLE todos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_todos_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);
```

---

## 🔐 Autentikasi JWT

- User login akan menerima token JWT
- Sertakan header pada request ke endpoint terproteksi:


```Authorization: Bearer <token>```

---

## 🚀 Daftar Endpoint API

| Method | Endpoint         | Deskripsi                      |
|--------|------------------|--------------------------------|
| POST   | `/auth/register` | Registrasi user                |
| POST   | `/auth/login`    | Login dan dapatkan token JWT   |
| GET    | `/todos`         | Ambil semua todo (by user)     |
| POST   | `/todos`         | Tambahkan todo baru            |
| PUT    | `/todos/:id`     | Update todo berdasarkan ID     |
| DELETE | `/todos/:id`     | Hapus todo berdasarkan ID      |

---

## ▶️ Cara Menjalankan

### 1. Clone Repository

```
git clone https://github.com/username/todo-api-go.git
cd todo-api-go
```

### 2. Konfigurasi `.env`

```
DB_DSN=root:password@tcp(127.0.0.1:3306)/tododb
JWT_SECRET=your_jwt_secret
```

### 3. Jalankan Server

```
go run main.go
```

Server akan berjalan di: `http://localhost:8080`

---

## 🪵 Logging

Semua error, akses penting, dan aktivitas sensitif dicatat menggunakan `logrus`.

---

## ✨ Fitur Selanjutnya (opsional)

- Refresh token
- Pagination
- Role-based access control (RBAC)
- Unit test & integration test

---

## 🤝 Kontribusi

Pull Request sangat diterima!  
Silakan fork repo, buat branch baru, dan ajukan PR.

---

## 📄 Lisensi

MIT License.  
Silakan digunakan secara bebas untuk keperluan pribadi dan komersial.

