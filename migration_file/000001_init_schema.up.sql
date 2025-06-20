create table if not exists countries(
	country_id bigserial,
	country_code varchar(5),
	country_name varchar(255),
	continent_name varchar(255),
	constraint country_id_primary_key primary key(country_id),
	constraint country_code unique(country_code),
	constraint country_name_continent_name unique(country_name,continent_name)
);

create sequence if not exists countries_country_id_sequence
as bigint
start with 1
increment by 1 
no maxvalue 
no minvalue
no cycle
owned by countries.country_id;

create table if not exists users(
	user_id bigserial,
	first_name varchar(255),
	last_name varchar(255),
	email varchar(255),
	password varchar(255),
	salt varchar(255),
	country_id bigint,
	constraint user_id_primary_key primary key(user_id),
	constraint country_id_foreign_key foreign key(country_id) references countries(country_id) on delete set null
);

create sequence if not exists users_user_id_sequence
as bigint
start with 1
increment by 1 
no maxvalue 
no minvalue
no cycle
owned by users.user_id;



create table if not exists merchants(
	merchant_id bigserial,
	merchant_first_name varchar(255),
	merchant_last_name varchar(255),
	country_id bigint,
	user_id bigint,
	constraint merchant_id_primary_key primary key(merchant_id),
	constraint merchant_first_name_last_name_unique unique(merchant_first_name,merchant_last_name),
	constraint country_id_foreign_key foreign key(country_id) references countries(country_id) on delete set null,
	constraint user_id_foreign_key foreign key(user_id) references users(user_id) on delete set null
);

create sequence if not exists merchants_merchant_id_sequence
as bigint
increment by 1
start with 1
no maxvalue
no minvalue
no cycle
owned by merchants.merchant_id;



create table if not exists product_statuses(
	product_status_id bigserial, 
	status_name varchar(255),
	constraint product_status_id_primary_key primary key(product_status_id),
	constraint status_name_unique unique(status_name)
);

create sequence if not exists product_statuses_product_status_id_sequence
start with 1
increment by 1
no maxvalue
no minvalue
no cycle
owned by product_statuses.product_status_id;


create table if not exists products(
	product_id bigserial,
	product_name varchar(255),
	product_price int,
	product_status_id bigint,
	merchant_id bigint,
	constraint product_id_primary_key primary key(product_id),
	constraint product_status_id_foreign_key foreign key(product_status_id) references product_statuses(product_status_id) on delete set null,
	constraint merchant_id_foreign_key foreign key(merchant_id) references merchants(merchant_id) on delete set null
);

create sequence if not exists products_product_id_sequence
as bigint
increment by 1
start with 1
no maxvalue
no minvalue
no cycle
owned by products.product_id;

create table if not exists order_statuses(
	order_status_id bigserial,
	order_status_name varchar(255),
	constraint order_status_id_primary_key primary key(order_status_id),
	constraint order_status_name unique(order_status_name)
);

create sequence if not exists order_statuses_order_status_id
as bigint
increment by 1
start with 1
no maxvalue
no minvalue
no cycle
owned by order_statuses.order_status_id;

create table if not exists order_header(
	order_header_id bigserial,
	order_date timestamp,
	order_status_id bigint,
	user_id bigint,
	merchant_id bigint ,
	constraint order_header_id_primary_key primary key(order_header_id),
	constraint user_id_foreign_key foreign key(user_id) references users(user_id) on delete set null,
	constraint merchant_id_foreign_key foreign key(merchant_id) references merchants(merchant_id) on delete set null,
	constraint order_status_id_foreign_key foreign key(order_status_id) references order_statuses(order_status_id) on delete set null
);

create sequence if not exists order_header_order_header_id_sequence
as bigint
increment by 1
start with 1
no maxvalue
no minvalue
no cycle
owned by order_header.order_header_id;

create table if not exists order_lines(
	line_id bigserial,
	product_id bigint,
	order_header_id bigint,
	quantity int,
	constraint line_id_primary_key primary key(line_id),
	constraint product_id_foreign_key foreign key(product_id) references products(product_id) on delete set null,
	constraint order_header_id_foreign_key foreign key(order_header_id) references order_header(order_header_id) on delete set null,
	constraint product_id_order_header_id_unique unique(product_id,order_header_id)
);

create sequence if not exists order_lines_line_id_sequence
as bigint
increment by 1
start with 1
no maxvalue
no minvalue
no cycle
owned by order_lines.line_id;