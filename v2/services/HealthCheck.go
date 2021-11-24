package services

import (
	"fmt"
	"ottodigital.id/library/healthcheck"
	"ottodigital.id/library/utils"
	"ottosfa-api-web/models"
)

// HealthCheck ..
func (svc *Service) HealthCheck(res *models.Response) {
	fmt.Println(">>> HealthCheck - Service <<<")

	// === database ===
	databaseReq := healthcheck.HealthCheckDBReq{
		Host:     utils.GetEnv("DB_POSTGRES_ADDRESS", "13.229.6.53"),
		Port:     utils.GetEnv("DB_POSTGRES_PORT", "5432"),
		User:     utils.GetEnv("DB_POSTGRES_USER", "ubuntu"),
		Password: utils.GetEnv("DB_POSTGRES_PASS", "Ubuntu!23"),
		DBName:   utils.GetEnv("DB_POSTGRES_NAME", "otto-sfa-admin-api_copy_production"),
	}
	database := healthcheck.GenerateHealthCheckPostgres(databaseReq)

	// === redis ===
	redisReq := healthcheck.HealthCheckRedisReq{
		HostCluster: []string{
			utils.GetEnv("REDIS_HOST_CLUSTER1", "13.228.23.160:8079"),
			utils.GetEnv("REDIS_HOST_CLUSTER2", "13.228.23.160:8078"),
			utils.GetEnv("REDIS_HOST_CLUSTER3", "13.228.23.160:8077"),
		},
	}
	clusterRedis := healthcheck.GenerateHealthCheckRedisCluster(redisReq)

	healthCheckData := make([]healthcheck.DataHealthCheck, 0)
	healthCheckData = append(healthCheckData,
		<-database,
		<-clusterRedis,
	)

	resHealthCheck := healthcheck.GenerateResponseHealthCheck(healthCheckData...)

	res.Meta = models.Meta{
		Status:  true,
		Code:    resHealthCheck.ResponseCode,
		Message: resHealthCheck.Message,
	}
	res.Data = resHealthCheck.Data
	return
}
