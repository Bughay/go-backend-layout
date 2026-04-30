

-- name: Create :one
-- Inserts a new user
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING id, email, role, created_at;


-- name: FindByEmail :one
-- Finds a user by email
SELECT id, email, password, role, created_at 
FROM users 
WHERE email = $1;



-- name: FindByID :one
-- Finds a user by ID
SELECT id, email, role, created_at 
FROM users 
WHERE id = $1;

