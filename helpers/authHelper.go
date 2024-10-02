package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized Access")
		return err
	}
	return err
}
func MatchUserTypeToUid(c *gin.Context, UserId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	if userType == "USER" && uid != UserId {
		err = errors.New("Not authorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
