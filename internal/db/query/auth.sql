-- name: CreateUser :exec
INSERT INTO cvz_users (
    id, email, password, role
) VALUES (
    $1, $2, $3, $4
);

-- name: FindUser :one
SELECT * FROM cvz_users 
WHERE email = $1 LIMIT 1;