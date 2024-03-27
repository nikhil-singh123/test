package routes

import (
	"test/auth"
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {

	admin := r.Group("/admin")
	admin.Use(auth.AdminAuthMiddleware())
	{
		admin.GET("/search-book/:email/:isbn", controllers.UpdateISBN)
		admin.POST("/add-user/:email", controllers.UserCreate)
		admin.POST("/add-library/:email", controllers.RegLibrary)
		admin.POST("/add-book/:email", controllers.AddBook)
		admin.DELETE("/remove-book/:email", controllers.RemoveBook)
		admin.PUT("/update-book/:email", controllers.UpdateBook)
		admin.GET("/list-issue-request/:email", controllers.ListIssueRequest)
		admin.POST("/approve-reject-issue-request/:email", controllers.ApproveRejectIssueRequest)
	}

	reader := r.Group("/reader")
	admin.Use(auth.ReaderAuthMiddleware())
	{
		reader.GET("/search-book/:email/:query", controllers.SearchBook)
		reader.POST("/raise-issue-request/:email", controllers.RaiseIssueRequest)
	}

}
