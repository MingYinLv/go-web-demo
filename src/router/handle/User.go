package handle

import (
	"webtest/src/api/v1/db"
	"webtest/src/api/v1/db/Schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserList(c *gin.Context) {
	c.JSON(200, db.SearchUserList())
}

func GetUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, db.SearchUserById(id))
}

func DeleteUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	line, err := db.DeleteUserById(id)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{
		"delete": line,
	})
}

func AddUser(c *gin.Context) {
	user := Schema.User{Username: c.PostForm("username"), Password: c.PostForm("password")}
	id, err := db.AddUser(user)
	if err != nil {
		panic(err.Error())
	}
	user.User_id = id
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	user := Schema.User{User_id: id, Username: c.PostForm("username"), Password: c.PostForm("password")}
	_, err = db.UpdateUser(user)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, user)
}

func PatchUpdateUser(c *gin.Context) {
	db.PatchUpdateUser(c)
}
