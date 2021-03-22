package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-gin-admin/dao"
	"strconv"
)

func GetSomeData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bookList, err := dao.GlobalDao.GetSomeData(c, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, bookList)
	}
}

func GetSomeOtherData(c *gin.Context) {

}
