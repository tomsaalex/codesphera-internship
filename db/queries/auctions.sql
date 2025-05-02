-- name: GetAuctions :many
SELECT * FROM auction_details WHERE
-- TODO: Pretty sure there's SQL injection in that LIKE....
((sqlc.narg(product_name)::text IS NULL OR product_name LIKE '%' || sqlc.narg(product_name)::text || '%') OR
(sqlc.narg(product_desc)::text IS NULL OR product_desc LIKE '%' || sqlc.narg(product_desc)::text || '%')) AND
(category_name = sqlc.narg(category_name) OR sqlc.narg(category_name) IS NULL)
ORDER BY
    -- Sorting for product_name
    CASE WHEN @order_by = 'product_name' AND NOT @reverse THEN product_name END ASC,
    CASE WHEN @order_by = 'product_name' AND @reverse THEN product_name END DESC,

    -- Sorting for created_at
    CASE WHEN @order_by = 'created_at' AND NOT @reverse THEN created_at END ASC,
    CASE WHEN @order_by = 'created_at' AND @reverse THEN created_at END DESC
-- limit and offset are NOT good for pagination, but let's ignore that for now
LIMIT @page_size OFFSET @skipped_pages;

-- name: AddAuction :one
insert into auctions
(seller_id, category_id, product_name, product_desc, auc_mode, auc_status, starting_price, target_price, created_at)
values
($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetAllAuctionsByUser :many
SELECT * FROM auction_details 
WHERE seller_id = $1;

-- name: GetAuctionByName :one
SELECT * FROM auction_details
WHERE product_name = $1 LIMIT 1;