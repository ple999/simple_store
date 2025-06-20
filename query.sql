-- name: GetAllCountries :many
select * from countries;

-- name: GetAllCountriesByPage :many
select * from countries
offset $1
fetch next $2 rows only;

-- name: GetCountryByID :one
select * from countries
where country_id=$1;

-- name: GetAllUsers :many
select * from users;

-- name: GetUserByEmail :one
select salt,password from users
where email=$1;

-- name: GetAllMerchants :many
select * from merchants;

-- name: GetAllProducts :many
select * from products;

-- name: GetAllOrders :many
select * from order_header;

-- name: InsertCountry :one
insert into countries(
    country_code,
    country_name,
    continent_name
) values(
    $1,
    $2,
    $3
) 
returning *;

-- name: InsertUser :one
insert into users(
    first_name,
    last_name,
    password,
    salt,
    email,
    country_id
) values(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
returning *;

-- name: InsertMerchant :one
insert into merchants(
    merchant_first_name,
    merchant_last_name,
    country_id,
    user_id
) values(
    $1,
    $2,
    $3,
    $4
)
returning *;

-- name: InsertProductStatus :one
insert into product_statuses(
    status_name
)values(
    $1
)
returning *;


-- name: InsertProduct :one
insert into products(
    product_name,
    product_price,
    product_status_id,
    merchant_id
) values(
    $1,
    $2,
    $3,
    $4
)
returning *;

-- name: InsertOrderStatus :one
insert into order_statuses(
    order_status_name
) values(
    $1
)
returning *;

-- name: InsertOrderHeader :one
insert into order_header(
    order_date,
    order_status_id,
    user_id,
    merchant_id
) values(
    $1,
    $2,
    $3,
    $4
)
returning *;

-- name: InsertOrderLines :one
insert into order_lines(
    product_id,
    quantity,
    order_header_id
) values(
    $1,
    $2,
    $3
)
returning *;

-- name: UpdateCountry :one
update countries
set country_code=$1,
country_name=$2,
continent_name=$3
where country_id=$4
returning *;

-- name: UpdateUser :one
update users
set first_name=$1,
last_name=$2,
country_id=$3
where user_id=$4
returning *;

-- name: UpdateMerchant :one
update merchants
set merchant_first_name=$1,
merchant_last_name=$2,
country_id=$3
where merchant_id=$4
returning *;

-- name: UpdateProductStatus :one
update product_statuses
set status_name=$1
where product_status_id=$2
returning *;

-- name: UpdateProduct :one
update products
set product_name=$1,
product_price=$2,
product_status_id=$3,
merchant_id=$4
where product_id=$5
returning *;

-- name: UpdateOrderStatus :one
update order_statuses
set order_status_name=$1
where order_status_id=$2
returning *;

-- name: UpdateOrderHeader :one
update order_header 
set order_date=$1,
order_status_id=$2
where order_header_id=$3
returning *;

-- name: UpdateOrderLine :one
update order_lines
set quantity=$1,
product_id=$2
where line_id=$3
returning *;

-- name: DeleteCountries :exec
delete from countries
where country_id=$1;

-- name: DeleteUser :exec
delete from users
where user_id=$1;

-- name: DeleteMerchant :exec
delete from  merchants
where merchant_id=$1;

-- name: DeleteProductStatus :exec
delete from product_statuses
where product_status_id=$1;

-- name: DeleteProduct :exec
delete from products
where product_id=$1;

-- name: DeleteOrderStatus :exec
delete from order_statuses
where order_status_id=$1;

-- name: DeleteOrderHeader :exec
delete from  order_header
where order_header_id=$1;

-- name: DeleteOrderLines :exec
delete from order_lines
where line_id=$1;

-- name: DeleteOrderLinesUsingOrderHeaderID :exec
delete from order_lines
where order_header_id=$1;