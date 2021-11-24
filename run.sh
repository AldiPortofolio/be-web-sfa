export LOGGER_FILENAME="./otto-logger.log"

#main.go
export OTTOSFA_API_WEB_PORT=0.0.0.0:8045
export MAXPROCS="1"

#routes/route.go
export OTTOSFA_API_APK="OTTOSFA_API_WEB"
export READ_TIMEOUT="120"
export WRITE_TIMEOUT="120"
export JAEGER_HOSTURL="13.250.21.165:5775"

#db/postgres.go
export DB_POSTGRES_USER=sfa_web_usr
export DB_POSTGRES_PASS="sfa32021"
export DB_POSTGRES_NAME="otto-sfa-admin-api"
export DB_POSTGRES_ADDRESS="34.101.208.156"
export DB_POSTGRES_PORT="6432"
export DB_POSTGRES_DEBUG="true"
export DB_POSTGRES_TYPE="postgres"
export DB_POSTGRES_SSL_MODE="disable"
export DB_LIMIT_DATA=" LIMIT 20 "

export ROSE_POSTGRES_USER="ottoagcfg"
#ROSE_POSTGRES_PASS="dTj*&56$es", ditambah backslash (\) untuk escape karakter '$'
export ROSE_POSTGRES_PASS="dTj*&56\$es"
export ROSE_POSTGRES_NAME="rosedb"
export ROSE_POSTGRES_HOST="13.228.23.160"
export ROSE_POSTGRES_PORT="8432"
export ROSE_POSTGRES_DEBUG="true"
export ROSE_TYPE="postgres"
export ROSE_POSTGRES_SSL_MODE="disable"
export ROSE_POSTGRES_TIMEOUT="10"

#router/routers.go
#swagger
export APPS_ENV="local"
export SWAGGER_HOST_LOCAL="localhost:8045"
export SWAGGER_HOST_DEV="13.228.25.85:8045"

#REDIS
export REDIS_HOST_CLUSTER1=13.228.23.160:8079
export REDIS_HOST_CLUSTER2=13.228.23.160:8078
export REDIS_HOST_CLUSTER3=13.228.23.160:8077

export OTTO_SFA_NOTIF_HOST="http://34.101.141.240:8046/ottosfa-api-notif/v1"
export OTTO_SFA_NOTIF_ENDPOINT_CREATE_NOTIF="/notification/create"

go run main.go
