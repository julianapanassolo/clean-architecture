CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  products TEXT ARRAY NOT NULL,
  total_amount DECIMAL NOT NULL
);