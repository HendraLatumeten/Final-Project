package users

import (
	"encoding/json"

	"github.com/Lectron-Fazztrack/Final-Project/server/src/database/models"
	"github.com/Lectron-Fazztrack/Final-Project/server/src/interfaces"
	"github.com/Lectron-Fazztrack/Final-Project/server/src/libs"
	"github.com/gin-gonic/gin"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) Update(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}

	var data models.User
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}

	re.svc.Update(&data, claim_user.(string)).Send(c)
}
