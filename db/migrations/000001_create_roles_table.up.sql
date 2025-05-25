CREATE TABLE IF NOT EXISTS roles(
    id INT PRIMARY KEY auto_increment,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    CONSTRAINT unique_name UNIQUE(id)
)engine=InnoDB;
