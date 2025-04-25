-- name: AddAuction :one
insert into auctions
(seller_id, product_name, product_desc, auc_mode, auc_status, starting_price, target_price)
values
($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAllAuctionsByUser :many
SELECT * FROM auctions
WHERE seller_id = $1;

-- name: GetAuctionByName :one
SELECT * FROM auctions
WHERE product_name = $1 LIMIT 1;