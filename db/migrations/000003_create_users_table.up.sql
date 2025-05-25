CREATE TABLE IF NOT EXISTS users(
    id INT PRIMARY KEY auto_increment,
    email VARCHAR(100) NOT NULL,
    username VARCHAR(200) NOT NULL,
    role_id int NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    CONSTRAINT users_role_id FOREIGN KEY (role_id) REFERENCES roles (id)
)ENGINE=InnoDB