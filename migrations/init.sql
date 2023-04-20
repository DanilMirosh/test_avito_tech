CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists service, balance, reserve, report cascade;

create table if not exists service (
  id uuid unique not null default uuid_generate_v4(),
  service_name varchar(255) not null
);

insert into service(id, service_name) values ('a42753e9-d63e-451c-a9af-2ee62a3959ba', 'Append Balance'), ('103fb536-6f66-42b6-a49b-16c359e16635', 'Service transport posting payment'), ('deab7a77-95f9-44d6-bbdc-327d46d97a59', 'Service real estate posting payment'), ('828e355b-a721-42e3-b759-643411f07178', 'Receiving money from user'), ('eef5af83-a19b-493c-a3e1-d23dbcaec479', 'Sending money to user');

select * from service;

create table if not exists balance (
    user_uuid uuid not null default uuid_generate_v4(),
    balance bigint not null check ( balance >= 0 )
);

create table if not exists reserve (
    id uuid unique default uuid_generate_v4(),
    user_uuid uuid not null,
    service_uuid uuid references service(id),
    order_uuid uuid not null,
    reserve bigint not null check ( reserve >= 0 )
);

create table if not exists report (
    id uuid unique default uuid_generate_v4(),
    user_uuid uuid not null,
    service_uuid uuid references service(id),
    order_uuid uuid,
    money_amount bigint not null,
    operation_date timestamp not null
);

create or replace function update_balance(userUUID uuid, sum bigint) returns void as $$
UPDATE balance set balance = sum + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
insert into report(user_uuid, service_uuid, money_amount, operation_date) values (userUUID, 'a42753e9-d63e-451c-a9af-2ee62a3959ba', sum, CURRENT_TIMESTAMP);
$$ language sql;

create or replace function reserve_money(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
update balance set balance = (select balance from balance where user_uuid = userUUID) - amount
where user_uuid = userUUID;
insert into reserve(user_uuid, service_uuid, order_uuid, reserve) values (userUUID, serviceUUID, orderUUID, amount);
$$ language sql;

create or replace function check_money(IN balanceUUID uuid, IN amount bigint, OUT res bool) as
'select (select balance from balance where user_uuid = balanceUUID) >= amount'
language sql;

create or replace function accept_income(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
insert into report(user_uuid, service_uuid, order_uuid, money_amount, operation_date) values (userUUID, serviceUUID, orderUUID, amount, CURRENT_TIMESTAMP);
delete from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount
and id in (select id from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount LIMIT 1);
$$ language sql;

create or replace function check_required_reserve(IN userUUID uuid, IN serviceUUID uuid, IN orderUUID uuid, IN amount bigint, OUT res bool) as
'select EXISTS(SELECT * FROM reserve WHERE user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount)'
language sql;

create or replace function user_to_user_money_transfer(first_userUUID uuid, second_userUUID uuid, amount bigint) returns void as $$
update balance set balance = (select balance from balance where user_uuid = first_userUUID) - amount
where user_uuid = first_userUUID;
update balance set balance = amount + (select balance from balance where user_uuid = second_userUUID)
where user_uuid = second_userUUID;
insert into report(user_uuid, service_uuid, money_amount, operation_date) values (first_userUUID, 'eef5af83-a19b-493c-a3e1-d23dbcaec479', amount, CURRENT_TIMESTAMP),
                                                                                 (second_userUUID, '828e355b-a721-42e3-b759-643411f07178', amount, CURRENT_TIMESTAMP);
$$ language sql;

create or replace function unreserve_money(userUUID uuid, serviceUUID uuid, orderUUID uuid, amount bigint) returns void as $$
delete from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount
and id in (select id from reserve where user_uuid = userUUID and service_uuid = serviceUUID and order_uuid = orderUUID and reserve = amount LIMIT 1);
UPDATE balance set balance = amount + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
$$ language sql;

create or replace function create_new_balance(userUUID uuid, userBalance bigint) returns void as $$
INSERT INTO balance(user_uuid, balance) VALUES(userUUID, userBalance);
insert into report(user_uuid, service_uuid, money_amount, operation_date) values (userUUID, 'a42753e9-d63e-451c-a9af-2ee62a3959ba', userBalance, CURRENT_TIMESTAMP);
$$ language sql;

create or replace function check_transactions_by_date(IN timePeriod timestamp, OUT res bool) as
'select exists(select * from report where operation_date >= timePeriod
                                      and operation_date < date_trunc(''month'', timePeriod + interval ''1 month''))'
    language sql;

create or replace function get_all_transactions(timePeriod timestamp) returns table (service_name text, money_amount bigint) as $$
select distinct service.service_name, (select sum(money_amount) from report where service_uuid = service.id and operation_date >= timePeriod
                                                             and operation_date < date_trunc('month', timePeriod + interval '1 month')) from report
    join service on report.service_uuid = service.id
    $$ language sql;
