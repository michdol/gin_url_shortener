package Controllers

import (
	"net/http"
	"../Models"
	"github.com/gin-gonic/gin"
)

func GetUrls(c *gin.Context) {
	var urls []Models.Url
	err := Models.GetAllUrls(&urls)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, urls)
	}
}

func CreateUrl(c *gin.Context) {
	var url Models.Url
	// what does it do
	c.BindJSON(&url)
	err := Models.CreateUrl(&url)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, url)
	}
}

func GetUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var url Models.Url
	err := Models.GetUrl(&url, id)
	if err != nil {
		c.JSON(http.StatusNotFound, url)
	} else {
		c.JSON(http.StatusOK, url)
	}
}

func DeleteUrl(c *gin.Context) {
	var url Models.Url
	id := c.Params.ByName("id")
	err := Models.DeleteUrl(&url, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusNoContent, gin.H{"id": id})
	}
}
