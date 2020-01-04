### Конфигурация сервера

Серверная ОС - Ubuntu 18.04 LTS. 

Что должно быть установлено:

- [Git](https://www.digitalocean.com/community/tutorials/how-to-install-git-on-ubuntu-18-04-quickstart)
- [PostgreSQL](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-18-04) 
- [Go](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04)

Пользователь по умолчанию: `judge`. Используется и в ОС, и на сервере БД.

### Конфигурация базы данных

1. Создать роль `judge` с правами суперпользователя с помощью `createuser --interactive`
2. Создать базу данных `judge` c помощью `createdb judge`
3. Создать схему `judge` и установить `judge` её владельцем

```postgresql
CREATE SCHEMA judge;
ALTER SCHEMA judge OWNER TO judge;
```

Для создания таблиц воспользоваться файлом `judge_init.sql`.