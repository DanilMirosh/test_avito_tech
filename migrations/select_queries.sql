create or replace function update_balance(userUUID uuid, sum int) returns void as $$
UPDATE balance set balance = sum + (select balance from balance where user_uuid = userUUID)
where user_uuid = userUUID;
$$ language sql;

select update_balance('7e931dee-aa59-4687-a18a-44ea256fc733', 100);

create or replace function reserve_money(balanceUUID uuid, reserveUUID uuid, amount bigint) returns void as $$
    update balance set balance = (select balance from balance where user_uuid = balanceUUID) - amount
    where user_uuid = balanceUUID;
    update reserve set reserve = amount + (select reserve from reserve where user_uuid = reserveUUID)
    where user_uuid = reserveUUID;
    $$ language sql;

insert into balance(balance) values
                                 (10),
                                 (200);

insert into reserve(reserve) values
                                 (0),
                                 (0);

select reserve_money('bcaed786-44a4-45f6-a547-ab8642be3a0f', 'd1238908-6003-4c88-bca9-d09dae89e44a', 5);