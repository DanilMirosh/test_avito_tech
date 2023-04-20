# Тестовое задание от Avito.Tech

## Кейс 

Описание задачи

Разработать микросервис для работы с балансом пользователей (баланс, зачисление/списание/перевод средств).

Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON. Дополнительно реализовать методы конвертации баланса и получение списка транзакций. Полное описание в TASK (https://github.com/cowabungal/avitoTech/blob/master/TASK.md).

Реализация
- ️Следование дизайну REST API.
- ️Подход "Чистой Архитектуры" и техника внедрения зависимости.
- ️Работа с фреймворком gin-gonic/gin.
- Работа с СУБД Postgres с использованием библиотеки sqlx и написанием SQL запросов.
- ️Конфигурация приложения - библиотека viper.
- Запуск из Docker.
- ️Unit/Интеграционное - тестирование уровней обработчикоов, бизнес-логики и взаимодействия с БД с помощью моков - библиотеки testify, mock.

## Используемые технологии

- Web-framework: [`Gin Web Framework`](https://pkg.go.dev/github.com/gin-gonic/gin#section-readme)
- Базан данных: [`PostgreSQL 14`](https://www.postgresql.org/)
- Документация:
  - [`Swagger`](https://swagger.io/)
  - [`swaggo`](https://github.com/swaggo/swag#supported-web-frameworks)
  - [`Postman`](https://www.postman.com/)
    - [Postman файл](api/Avito.postman_collection.json) с эндпоинтами для тестирования приложения
- Контейнеризация:
  - [`Docker 20.10.21`](https://www.docker.com/) - ссылка на [Dockerfile](Dockerfile) приложения:
    - Использован multi-stage build для разделения процесса сборки
  - [`docker-compose 1.29.2`](https://docs.docker.com/compose/) - для автоматического поднятия окружения: базы данных и прилоежения:
    - сервис `psql` (порт 5432) - Контейнер базы данных на основе легковесного ообраза linux alpine
    - сервис `adminer` (порт 8080) - Для отладки работы с БД, просмотра изменений (веб-интерфейс)
    - сервис `app` (порт 9000) - Разработанное приложение
  - `GNU make 4.3` 

## Документация и запуск

Документация доступна для просмотра по адресу [`http://localhost:9000/swagger/index.html`](http://localhost:9000/swagger/index.html#/)

Для запуска выполнить сборку приложения

```bash
docker-compose up --build
```

## Выполненные требования

- [x] Сервис должен предоставлять HTTP API с форматом JSON как при отправке запроса, так и при получении результата.
- [x] Язык разработки: Golang.
- [x] Фреймворк gin-gonic/gin.
- [x] Реляционная СУБД: MySQL или PostgreSQL.
- [x] Использование docker и docker-compose для поднятия и развертывания dev-среды.
- [x] Весь код должен быть выложен на Github с Readme файлом с инструкцией по запуску и примерами запросов/ответов (можно просто описать в Readme методы, можно через Postman, можно в Readme curl запросы скопировать, и так далее).
- [x] Если есть потребность в асинхронных сценариях, то использование любых систем очередей - допускается.
- [x] При возникновении вопросов по ТЗ оставляем принятие решения за кандидатом (в таком случае Readme файле к проекту должен быть указан список вопросов с которыми кандидат столкнулся и каким образом он их решил).
- [x] Разработка интерфейса в браузере НЕ ТРЕБУЕТСЯ. Взаимодействие с API предполагается посредством запросов из кода другого сервиса. Для тестирования можно использовать любой удобный инструмент. Например: в терминале через curl или Postman.

## Выполненные дополнительные задания

- [x] Реализован метод для получения месячного отчета. На вход: год-месяц. На выходе ссылка на CSV файл.
- [x] Реализован метод получения списка транзакций с комментариями откуда и зачем были начислены/списаны средства с баланса. Предусмотрена пагинация и сортировка по сумме и дате.
- [x] Реализован сценарий разрезервирования денег, если услугу применить не удалось.