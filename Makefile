##### 用户服务 ##########
# 编译服务
build_user_service:
	docker-compose build laracom-user-service
# 启动服务
run_user_service:
	docker-compose up --build -d laracom-user-service
#######################

##### 用户服务客户端 #####
# 编译服务
build_user_cli:
	docker-compose build laracom-user-cli
# 启动服务
run_user_cli:
	docker-compose run -d laracom-user-cli --name="test" --email="test@leor.com" --password="123456"
#######################

##### 商品服务 #########
# 编译服务
build_product_service:
	docker-compose build laracom-product-service
# 启动服务
run_product_service:
	docker-compose up --build -d laracom-product-service
#######################

# 启动用户数据库服务
run_user_db:
	docker-compose up -d laracom-user-db

# 启动商品数据库服务
run_product_db:
	docker-compose up -d laracom-product-db

# 启动 micro 网关服务
run_micro_api:
	docker-compose up -d laracom-micro-api

# 启动 nats 服务
run_nats:
	docker-compose up -d laracom-nats
