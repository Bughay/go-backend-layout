-- name: CreateProduct :one
-- Inserts products to database
INSERT INTO products (user_id, name, description, price, stock)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, name, description, price, stock, created_at, updated_at;

-- name: FindAllProducts :many
-- Finds all products
WITH product_count AS (SELECT COUNT(*) FROM products)
SELECT 
    id, 
    user_id, 
    name, 
    description, 
    price, 
    stock, 
    created_at, 
    updated_at,
    (SELECT * FROM product_count) AS total_count
FROM products
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: FindProductByID :one
-- Finds products by ID
SELECT id, user_id, name, description, price, stock, created_at, updated_at 
FROM products 
WHERE id = $1;

-- name: UpdateProduct :one
-- Updates product
UPDATE products
SET 
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    price = COALESCE(sqlc.narg('price'), price),
    stock = COALESCE(sqlc.narg('stock'), stock),
    updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING id, user_id, name, description, price, stock, created_at, updated_at;

-- name: DeleteProduct :execrows
-- Delete a product
DELETE FROM products WHERE id = $1;