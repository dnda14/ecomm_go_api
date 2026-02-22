-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductByID :one
select * from products
where id = $1;
