insert into "user"(id, email, passwd) values
    ('12f1bd84-2892-497d-951c-054b3eb7f818', 'maksim@gmail.com', '$2a$16$sMrNQOCu.xi8gjgyTfjvZ.stPuxpXGP5klx02CyGRdYwODS6drvLa'),
    ('aec85bb9-8d1f-446a-9495-8e4828da8a26', 'kutsenko@mail.ru', '$2a$16$BK1p5.FgSr9zbQB4eZTWMunMlfH5MMV.aTmQbTBcAMZCbXpqp5sQS'),
    ('30e98cfe-ab7c-4171-8dd6-187f0d29e75e', 'valevin_grigory@internet.ru', '$2a$16$GrAD5HQpIRtwgKS3dDsEhe30.e4lIt8ZVOBX1T/LvtyaGmrV2kxdi');

    -- 16 rounds encryption
    -- maksim@gmail.com:1234qwerty
    -- kutsenko@mail.ru:parol_ogurec123
    -- valevin_grigory@internet.ru:xxx-xaker-999

insert into profile(user_uuid, firstname, lastname) values
    ('12f1bd84-2892-497d-951c-054b3eb7f818', 'Максим', 'Лагус'),
    ('aec85bb9-8d1f-446a-9495-8e4828da8a26', 'Алексей', 'Куценко'),
    ('30e98cfe-ab7c-4171-8dd6-187f0d29e75e', 'Григорий', 'Валевин');

insert into service(id, "name", description, price) values
    ('1da52566-5acd-4cbb-a61e-fd372d9bbac0', 'Разовое размещение объявления: Недвижимость', 'Предоставление возможности пользователю осуществить однократное размещение объявления на Авито в категории Недвижимость', 1000),
    ('83543ffe-d3eb-4e60-9325-44014772de33', 'Разовое размещение объявления: Транспорт', 'Предоставление возможности пользователю осуществить однократное размещение объявления на Авито в категории Транспорт', 500),
    ('f60348bf-9d2d-4a49-aede-80222f4c007d', 'Разовое размещение объявления: Услуги', 'Предоставление возможности пользователю осуществить однократное размещение объявления на Авито в категории Услуги', 250),
    ('96a2eda3-a0c4-4252-a31b-be46099612e8', 'Разовое размещение объявления: Работа', 'Предоставление возможности пользователю осуществить однократное размещение объявления на Авито в категории Работа', 750);

insert into balance(user_uuid, balance) values
    ('12f1bd84-2892-497d-951c-054b3eb7f818', 120500),
    ('aec85bb9-8d1f-446a-9495-8e4828da8a26', 5000),
    ('30e98cfe-ab7c-4171-8dd6-187f0d29e75e', 1000);

insert into reserve(user_uuid, reserve) values
    ('12f1bd84-2892-497d-951c-054b3eb7f818', 0),
    ('aec85bb9-8d1f-446a-9495-8e4828da8a26', 0),
    ('30e98cfe-ab7c-4171-8dd6-187f0d29e75e', 0);