package Controllers

import (
	"net/http"
	"url_shortener/Models"
	"github.com/gin-gonic/gin"
)


// curl localhost:8080/v1/urls
func GetUrls(c *gin.Context) {
	var urls []Models.Url
	err := Models.GetAllUrls(&urls)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, urls)
	}
}

// curl -X POST localhost:8080/v1/urls -d '{"url": "https://www.nintendo.com/switch/"}'
func CreateUrl(c *gin.Context) {
	var url Models.Url
	// what does it do
	c.BindJSON(&url)
	payloadUrl := url.Url

	if !IsValidUrl(payloadUrl) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid url"})
		return
	}

	err := Models.GetUrlByUrl(&url, payloadUrl)

	// Return if already exists
	if err == nil {
		c.JSON(http.StatusOK, url)
		return
	}

	err = Models.CreateUrl(&url)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, url)
	}
}

// curl localhost:8080/v1/urls/:id
func GetUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var url Models.Url
	err := Models.GetUrl(&url, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, url)
	}
}

// curl -X DELETE localhost:8080/v1/urls/:id
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

// http://localhost:8080/r/:id
func RedirectById(c *gin.Context) {
	var url Models.Url
	id := c.Params.ByName("id")
	err := Models.GetUrl(&url, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusMovedPermanently, url.Url)
	}
}