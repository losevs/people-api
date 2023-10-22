#### Для запуска:

Добавить файл .env и заполнить его в соответствии:

```
DB_HOST= your_host
DB_USER= your_user
DB_PORT= your_port
DB_PASS= your_password
DB_NAME= db_name
```

Запуск:

```
$ make run
```

##### Get

`localhost:80/show/`

`localhost:80/show/:id`

`localhost:80/show/age/asc`

`localhost:80/show/pag/:page`

`localhost:80/show//pag/men/:page`

`localhost:80/show/pag/wmen/:page`

`localhost:80/show/filt/sex/:sex`

`localhost:80/show/filt/age/:age`

`localhost:80/show/filt/country/:country`

##### Post

`localhost:80/new`

##### Patch

`localhost:80/change/:id`

##### Delete

`localhost:80/del/:id`