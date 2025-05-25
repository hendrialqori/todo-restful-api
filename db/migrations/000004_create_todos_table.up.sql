CREATE TABLE IF NOT EXISTS todos(
    id INT PRIMARY KEY auto_increment,
    title VARCHAR(200) NOT NULL,
    status ENUM("DONE","PENDING") NOT NULL DEFAULT "PENDING",
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    CONSTRAINT todos_user_id FOREIGN KEY (user_id) REFERENCES users (id)
)ENGINE=InnoDB;