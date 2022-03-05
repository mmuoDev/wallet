CREATE TABLE wallet(
   id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   previous_balance INT,
   current_balance int,
   account_id VARCHAR(200) UNIQUE,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);