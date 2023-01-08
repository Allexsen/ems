package routes // profiles

import (
	"net/http"

	session "github.com/Allexsen/ems/api/middlewares/sessions"
	prof_id "github.com/Allexsen/ems/pkg/controllers/pid"
	_ "github.com/Allexsen/ems/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ph struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func initProfile(r *gin.Engine) {
	profile := r.Group("/profile")
	profile.Use(session.CheckSession())
	{
		profile.GET("", func(c *gin.Context) {

			session := sessions.Default(c)
			userID := session.Get("user_id")
			firstName := session.Get("firstname")
			lastName := session.Get("lastname")

			user := ph{
				ID:        userID.(int),
				FirstName: firstName.(string),
				LastName:  lastName.(string),
			}

			c.JSON(http.StatusAccepted, user)
			// c.String(200, "Self Profile")
		})

		profile.GET("/:pid", func(c *gin.Context) {
			employee, err := prof_id.GetProfile(c.Param("pid"))
			if err != nil {
				c.Abort()
				panic(err)
			}
			c.JSON(http.StatusOK, employee)
		})

		profile.GET("/:pid/history", func(c *gin.Context) {
			c.String(200, "Employee History")
		})

		profile.GET("/:pid/reviews", func(c *gin.Context) {
			c.String(200, "Employee Reviews")
		})

		profile.POST("/:pid/reviews/add", func(c *gin.Context) {
			c.String(200, "Add Review")
		})
	}
}
