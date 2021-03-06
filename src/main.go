package main

import (
	myService "BapOgrenciAPI/Services"
	_ "BapOgrenciAPI/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const (
	userNameSvc = "user"
	passwordSvc = "unp.123*"
)

// GetStudents example
// @Summary Get Students
// @Description Get All Students
// @Accept  json
// @Success 200 {object} Models.OgrenciSrvcResult
// @Router /student/getStudents [get]
func GetStudents(c *gin.Context) {
	studentResult := myService.GetStudents()
	c.JSON(http.StatusOK, studentResult)
}

// GetOrganizationFromStudent example
// @Summary GetOrganizationFromStudent
// @Description GetOrganizationFromStudent
// @Accept  json
// @Success 200 {object} Models.OrganizasyonSrvcResult
// @Router /student/getOrganizationFromStudent [get]
func GetOrganizationFromStudent(c *gin.Context) {
	studentResult := myService.GetOrganizationFromStudent()
	c.JSON(http.StatusOK, studentResult)
}

// GetPersonels example
// @Summary Get Personels
// @Description Get All Personels
// @Accept  json
// @Success 200 {object} Models.PersonelSrvcResult
// @Router /personel/getPersonels [get]
func GetPersonels(c *gin.Context) {
	personels := myService.GetPersonels()
	c.JSON(http.StatusOK, personels)
}

// GetOrganizationFromPersonel example
// @Summary GetOrganizationFromPersonel
// @Description GetOrganizationFromPersonel
// @Accept  json
// @Success 200 {object} Models.OrganizasyonSrvcResult
// @Router /personel/getOrganizationFromPersonel [get]
func GetOrganizationFromPersonel(c *gin.Context) {
	orgResult := myService.GetOrganizationFromPersonel()
	c.JSON(http.StatusOK, orgResult)
}

func basicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == userNameSvc && password == passwordSvc {
		c.Next()
	} else {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// @title Swagger Bap Integration API
// @version 1.0
// @description Integration i??in Data Sa??layan servis
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8011
// @BasePath /api/v1
// servicescore.ege.edu.tr/GoBapIntegrationAPI
func main() {
	r := gin.Default()

	//url := ginSwagger.URL("http://localhost:9000/swagger/doc.json") // The url pointing to API definition
	url := ginSwagger.URL("http://localhost:8011/swagger/doc.json")

	studentapi := r.Group("/api/v1/student")
	{
		studentapi.GET("/getStudents", basicAuth, GetStudents)
		studentapi.GET("/getOrganizationFromStudent", basicAuth, GetOrganizationFromStudent)

	}

	personelapi := r.Group("/api/v1/personel")
	{
		personelapi.GET("/getPersonels", basicAuth, GetPersonels)
		personelapi.GET("/getOrganizationFromPersonel", basicAuth, GetOrganizationFromPersonel)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://servicescore.ege.edu.tr"}
	//r.Use(cors.New(config))

	r.Use(cors.Default())

	r.Run(":9000")

}
