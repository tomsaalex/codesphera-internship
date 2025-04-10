-- name: GetProduct :one
SELECT * FROM products
WHERE name = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: CreateProduct :one
INSERT INTO products 
(name, description, price)
VALUES
($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
	set description = $2,
	price = $3,
	issold=$4
WHERE name = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE name = $1;
