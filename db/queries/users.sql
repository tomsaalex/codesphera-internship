-- name: AddUser :one
insert into users
(email, fullname, pass_hash, pass_salt)
values
($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
set email = $2,
fullname = $3,
pass_hash = $4,
pass_salt = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
