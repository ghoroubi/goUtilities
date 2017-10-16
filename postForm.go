package main

import (
	"github.com/gin-gonic/gin"

)

func  GetParamsBuJSON(c *gin.Context) {

	var err error
	var length,width int
	var area,square,radius float64
	var shapeName,shapeType,subject,body string

	toBePosted :=struct {
		width, length *int
		shapeName,shapeType,subject,body *string
		area,square,radius *float64
	}{}
	err=c.BindJSON(&toBePosted)
	token := c.Request.Header.Get("Authorization")
	if x:= toBePosted.length;x!=nil{
		length=*x
	}
	if x:= toBePosted.width;x!=nil{
		width=*x
	}
	if x:= toBePosted.area;x!=nil{
		area=*x
	}
	if x:= toBePosted.square;x!=nil{
		square=*x
	}
	if x:= toBePosted.radius;x!=nil{
		radius=*x
	}
	if x:= toBePosted.shapeName;x!=nil{
		shapeName=*x
	}
	if x:= toBePosted.shapeType;x!=nil{
		shapeType=*x
	}

	if x:= toBePosted.subject;x!=nil{
		subject=*x
	}
	if x:= toBePosted.body;x!=nil{
		body=*x
	}

	success, fail, err := Post(token, length,width,radius,area,square,subject, body)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK,"Status")
}