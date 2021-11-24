package routers

import (
	"fmt"
	"io"
	"os"
	"ottosfa-api-web/docs"
	v22merchantsNewRecController "ottosfa-api-web/v2.2/controllers/merchantNewRecruitment"
	v23countryontroller "ottosfa-api-web/v2.3/controllers/country"
	v23gendercontroller "ottosfa-api-web/v2.3/controllers/gender"
	v23jobmanagement "ottosfa-api-web/v2.3/controllers/jobManagement"
	v23jobcategories "ottosfa-api-web/v2.3/controllers/job_category"
	v23todolistController "ottosfa-api-web/v2.3/controllers/todolist"

	v23Controller "ottosfa-api-web/v2.3/controllers"
	v23attendanceController "ottosfa-api-web/v2.3/controllers/attendance"

	v2Controller "ottosfa-api-web/v2/controllers"
	v2merchantsNewRecController "ottosfa-api-web/v2/controllers/merchantNewRecruitment"
	v2paramConfigurationController "ottosfa-api-web/v2/controllers/paramConfiguration"
	v2todolistController "ottosfa-api-web/v2/controllers/todolist"

	v23ActivitasSalesmenController "ottosfa-api-web/v2.3/controllers/activitasSalesmen"
	v23AdminController "ottosfa-api-web/v2.3/controllers/admin"
	v23AdminSubAreaController "ottosfa-api-web/v2.3/controllers/adminSubArea"
	v23CompanyController "ottosfa-api-web/v2.3/controllers/company"

	// v23countryontroller "ottosfa-api-web/v2.3/controllers/country"
	// v23gendercontroller "ottosfa-api-web/v2.3/controllers/gender"
	v23provinceController "ottosfa-api-web/v2.3/controllers/province"
	v23roleController "ottosfa-api-web/v2.3/controllers/role"

	// v23todolistController "ottosfa-api-web/v2.3/controllers/todolist"
	v23villageController "ottosfa-api-web/v2.3/controllers/village"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/opentracing/opentracing-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"ottodigital.id/library/httpserver/ginserver"
	ottologer "ottodigital.id/library/logger"
	"ottodigital.id/library/ottotracing"
	"ottodigital.id/library/utils"
)

// ServerEnv ..
type ServerEnv struct {
	ServiceName     string `envconfig:"SERVICE_NAME" default:"OTTOSFA-API-WEB"`
	OpenTracingHost string `envconfig:"OPEN_TRACING_HOST" default:"13.250.21.165:5775"`
	DebugMode       string `envconfig:"DEBUG_MODE" default:"debug"`
	ReadTimeout     int    `envconfig:"READ_TIMEOUT" default:"120"`
	WriteTimeout    int    `envconfig:"WRITE_TIMEOUT" default:"120"`
}

var (
	server ServerEnv
)

func init() {
	if err := envconfig.Process("SERVER", &server); err != nil {
		fmt.Println("Failed to get SERVER env:", err)
	}
}

// Server ..
func Server(listenAddress string) error {
	sugarLogger := ottologer.GetLogger()

	ottoRouter := OttoRouter{}
	ottoRouter.InitTracing()
	ottoRouter.Routers()
	defer ottoRouter.Close()

	err := ginserver.GinServerUp(listenAddress, ottoRouter.Router)

	if err != nil {
		fmt.Println("Error:", err)
		sugarLogger.Error("Error ", zap.Error(err))
		return err
	}

	fmt.Println("Server UP")
	sugarLogger.Info("Server UP ", zap.String("listenAddress", listenAddress))

	return err
}

// OttoRouter ..
type OttoRouter struct {
	Tracer   opentracing.Tracer
	Reporter jaeger.Reporter
	Closer   io.Closer
	Err      error
	GinFunc  gin.HandlerFunc
	Router   *gin.Engine
}

// Routers ..
func (ottoRouter *OttoRouter) Routers() {
	gin.SetMode(server.DebugMode)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "ottosfa-api-web API"
	docs.SwaggerInfo.Description = "<ottosfa-api-web description>"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	switch utils.GetEnv("APPS_ENV", "local") {
	case "local":
		docs.SwaggerInfo.Host = utils.GetEnv("SWAGGER_HOST_LOCAL", "localhost:8045")
	case "dev":
		docs.SwaggerInfo.Host = utils.GetEnv("SWAGGER_HOST_DEV", "13.228.25.85:8045")
	}

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(ottoRouter.GinFunc)
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	v2 := router.Group("/ottosfa/v2")
	{
		merchantsNewRec := v2.Group("/merchants-new-recruitment")
		{
			merchantsNewRec.POST("/list", v2merchantsNewRecController.List)
			merchantsNewRec.POST("/export", v2merchantsNewRecController.Export)
			merchantsNewRec.POST("/create", v2merchantsNewRecController.Create)
			merchantsNewRec.GET("/detail/:merchant_new_rec_id", v2merchantsNewRecController.Detail)
			merchantsNewRec.GET("/download-template", v2merchantsNewRecController.DownloadTemplate)
			merchantsNewRec.POST("/bulk", v2merchantsNewRecController.Upload)
		}

		todolist := v2.Group("/todolist")
		{
			todolist.POST("/create", v2todolistController.Create)
			todolist.POST("/new-merchant-detail", v2todolistController.NewMerchantDetail)
			todolist.GET("/detail/:todolist_id", v2todolistController.Detail)
			todolist.POST("/new-merchant-list", v2todolistController.NewMerchantList)
			todolist.POST("/bulk", v2todolistController.Upload)
			todolist.GET("/download-template", v2todolistController.DownloadTemplate)
			todolist.POST("/update", v2todolistController.Update)
		}

		paramConfig := v2.Group("/param-configuration")
		{
			paramConfig.GET("min-max-phone", v2paramConfigurationController.MinMaxPhone)
		}

		v2.POST("/institution/list", v2Controller.Institution)
		v2.POST("/province/list", v2Controller.Province)
		v2.POST("/city/list", v2Controller.City)
		v2.POST("/district/list", v2Controller.District)
		v2.POST("/village/list", v2Controller.Village)
		v2.POST("/sub-area-channel/list", v2Controller.SubAreaChannel)

		v2.GET("/merchant-type/list", v2Controller.MerchantType)
		v2.GET("/merchant-category/list", v2Controller.MerchantCategory)
		v2.POST("/merchant-group/list", v2Controller.MerchantGroup)
		v2.GET("/merchant-business-type/list", v2Controller.MerchantBusinessType)
		v2.GET("/sales-retail/list", v2Controller.SalesRetail)

		v2.GET("/health-check", v2Controller.HealthCheck)
	}

	v22 := router.Group("/ottosfa/v2.2")
	{
		merchantsNewRec := v22.Group("/merchants-new-recruitment")
		{
			merchantsNewRec.POST("/bulk", v22merchantsNewRecController.Upload)
		}
	}

	v23 := router.Group("/ottosfa/v2.3")
	{
		merchantsNewRec := v23.Group("/todolist")
		{
			merchantsNewRec.POST("/create", v23todolistController.Create)
			merchantsNewRec.POST("/update", v23todolistController.Update)
			merchantsNewRec.POST("/merchant-detail", v23todolistController.MerchantDetail)
			merchantsNewRec.POST("/merchant-list", v23todolistController.MerchantList)
			merchantsNewRec.POST("/bulk", v23todolistController.Upload)
			merchantsNewRec.GET("/detail/:todolist_id", v23todolistController.Detail)
		}

		country := v23.Group("/country")
		{
			country.GET("/", v23countryontroller.List)
		}

		gender := v23.Group("/gender")
		{
			gender.GET("/", v23gendercontroller.List)
		}

		jobcategories := v23.Group("/jobcategories")
		{
			jobcategories.POST("/filter", v23jobcategories.Filter)
			jobcategories.POST("/save", v23jobcategories.Create)
			jobcategories.POST("/update", v23jobcategories.Update)
			jobcategories.DELETE("/delete/:id", v23jobcategories.Delete)
			jobcategories.GET("/detail/:id", v23jobcategories.Detail)
		}

		jobManagement := v23.Group("/job-management")
		{
			jobManagement.POST("/filter", v23jobmanagement.Filter)
			jobManagement.POST("/save", v23jobmanagement.Create)
			jobManagement.POST("/draft", v23jobmanagement.Draft)

			jobManagement.POST("/edit", v23jobmanagement.Edit)
			jobManagement.POST("/upload", v23jobmanagement.Upload)
			jobManagement.DELETE("/delete/:id", v23jobmanagement.Delete)
			jobManagement.GET("/detail/:id", v23jobmanagement.Detail)
			jobManagement.GET("/check-admin", v23jobmanagement.CheckAdmin)
			jobManagement.POST("/recipient-list", v23jobmanagement.RecipientList)
		}

		province := v23.Group("/province")
		{
			province.GET("/:countryId", v23provinceController.ListByCountry)
		}

		admin := v23.Group("/admin")
		{
			role := admin.Group("/role")
			{
				role.GET("/list", v23roleController.List)
			}

			admin.GET("/action-types", v23AdminController.ActionTypes)
		}

		village := v23.Group("/village")
		{
			village.GET("/:districtId", v23villageController.ListByDistrict)
		}

		company := v23.Group("/company")
		{
			company.GET("/list", v23CompanyController.List)
			company.GET("/company-codes", v23CompanyController.CompanyCodes)
		}

		adminsubarea := v23.Group("/admin-sub-area")
		{
			adminsubarea.POST("/delete-by-admin", v23AdminSubAreaController.DeleteByAdmin)
		}

		activitasSalesmen := v23.Group("/activitas-salesmen")
		{
			activitasSalesmen.POST("/list", v23ActivitasSalesmenController.ListActivitasSalesmen)
			activitasSalesmen.POST("/detail-sales", v23ActivitasSalesmenController.DetailActivitasSalesmen)
			activitasSalesmen.POST("/list-detail-todolist", v23ActivitasSalesmenController.ListDetailActivitasSalesmenTodoList)
			activitasSalesmen.POST("/list-detail-callplan", v23ActivitasSalesmenController.ListDetailActivitasSalesmenCallplan)
			activitasSalesmen.GET("/detail-callplan/:callplanMerchantID", v23ActivitasSalesmenController.DetailCallPlan)
			activitasSalesmen.GET("/detail-todolist/:todoListID", v23ActivitasSalesmenController.DetailTodolist)
		}
		
		attendance := v23.Group("/attendances")
		{
			attendance.POST("/list", v23attendanceController.List)
			attendance.GET("/detail/:attendance_id", v23attendanceController.Detail)
			attendance.POST("/validate", v23attendanceController.Validate)
			attendance.POST("/export", v23attendanceController.Export)
		}
		
		uploadImage := v23.Group("/upload-image")
		{
			UploadImageController := new(v23Controller.UploadImageController)
			uploadImage.POST("", UploadImageController.Upload)
		}
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ottoRouter.Router = router

}

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// InitTracing ..
func (ottoRouter *OttoRouter) InitTracing() {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "PROD"
	}

	tracer, reporter, closer, err := ottotracing.InitTracing(fmt.Sprintf("%s::%s", server.ServiceName, hostName), server.OpenTracingHost, ottotracing.WithEnableInfoLog(true))
	if err != nil {
		fmt.Println("Error :", err)
	}
	opentracing.SetGlobalTracer(tracer)

	ottoRouter.Closer = closer
	ottoRouter.Reporter = reporter
	ottoRouter.Tracer = tracer
	ottoRouter.Err = err
	ottoRouter.GinFunc = ottotracing.OpenTracer([]byte("api-request-"))
}

// Close ..
func (ottoRouter *OttoRouter) Close() {
	ottoRouter.Closer.Close()
	ottoRouter.Reporter.Close()
}
