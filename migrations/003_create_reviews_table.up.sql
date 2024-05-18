CREATE TABLE reviews (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    product_id INT,
    rating INT,
    comments TEXT,
    review_time TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);