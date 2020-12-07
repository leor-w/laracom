build_user_service:
	docker-compose build laracom-user-service
run_user_service:
	docker-compose up -d laracom-user-service
build_user_cli:
	docker-compose build laracom-user-cli --name="test" --email="test@leor.com" --password="123456"
run_user_cli:
	docker-compose run -d laracom-user-cli --name="test" --email="test@leor.com" --password="123456"
run_user_db:
	docker-compose up -d laracom-user-db