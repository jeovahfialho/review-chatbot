CREATE TABLE interactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    interaction_time TIMESTAMP,
    message TEXT,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);