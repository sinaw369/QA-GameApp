CREATE TABLE users (
            id int PRIMARY KEY AUTO_INCREMENT,
            name varchar(255) NOT NULL,
            phone_number VARCHAR(255) NOT NULL unique ,
            password VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);