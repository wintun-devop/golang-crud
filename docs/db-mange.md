```
CREATE DATABASE go_crud;
```

```
CREATE TABLE products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    model TEXT NOT NULL UNIQUE,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
```
- function for update timestamp
```
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
```
- Insert a product
```
INSERT INTO products (name, model, price)
VALUES ('Laptop', 'XPS 13', 1299.99);
```
- Update a product
```
UPDATE products
SET price = 1199.99
WHERE id = 1;
```