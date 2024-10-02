up:
	docker-compose -p mysql-project up -d mysql

down:
	docker-compose -p mysql-project down mysql --remove-orphans

logs:
	docker-compose -p mysql-project logs -f mysql

mysql:
	docker-compose -p mysql-project exec mysql mysql -uuser -p recordings

clean: down
	docker volume rm mysql-project_mysql-data

set-env:
	export DBUSER=user
	export DBPASS=password
