# Migrations

- In order to use sqlc is good to use migrations (https://github.com/golang-migrate/migrate)
- To install migrations execute:

```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash

sudo apt-get update

sudo apt-get install migrate

migrate -version
```

- And at the root level of your project execute:

```
migrate create -ext=sql -dir=sql/migrations -seq init
```

- Execute the docker compose file at root level

```
docker-compose up -d
```
- To connect at the mysql instance use mysql client:
```
mysql -h 127.0.0.1 -P 3306 -u root -p
```

- To execute the migrations, run:
(up to create and down to clean up the resources)
```
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
```

- Install sqlc:
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```