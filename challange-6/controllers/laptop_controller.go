package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "golang.org/x/net/context/ctxhttp"
)

var LaptopDatas = []Laptop{}

type Laptop struct {
	LaptopID string `json:"laptopid"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Price    string `json:"price"`
}

func CreateLaptop(ctx *gin.Context) {
	var newLaptop Laptop

	if err := ctx.ShouldBindJSON(&newLaptop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newLaptop.LaptopID = fmt.Sprintf("c%d", len(LaptopDatas)+1)
	LaptopDatas = append(LaptopDatas, newLaptop)

	ctx.JSON(http.StatusCreated, gin.H{
		"laptop": newLaptop,
	})
}

func UpdateLaptop(ctx *gin.Context) {

	LaptopID := ctx.Param("LaptopID")
	condition := false
	var updateLaptop Laptop

	if err := ctx.ShouldBindJSON(&updateLaptop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i, laptop := range LaptopDatas {
		if LaptopID == laptop.LaptopID {
			condition = true
			LaptopDatas[i] = updateLaptop
			LaptopDatas[i].LaptopID = LaptopID

		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"eror_massage": fmt.Sprintf("laptop with data id %v not found", LaptopID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": fmt.Sprintf("laptop with data id %v has been succesfuly update", LaptopID),
	})

}

func GetLaptop(ctx *gin.Context) {

	LaptopID := ctx.Param("LaptopID")
	condition := false
	var LaptopData Laptop

	for i, laptop := range LaptopDatas {
		if LaptopID == laptop.LaptopID {
			condition = true
			LaptopData = LaptopDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"eror_massage": fmt.Sprintf("laptop with data id %v not found", LaptopID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"laptop": LaptopData,
	})

}

func DelateLaptop(ctx *gin.Context) {
	LaptopID := ctx.Param("LaptopID")
	condition := false
	var LaptopIndex int

	for i, laptop := range LaptopDatas {
		if LaptopID == laptop.LaptopID {
			condition = true
			LaptopIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"eror_massage": fmt.Sprintf("laptop with data id %v not found", LaptopID),
		})
		return
	}

	copy(LaptopDatas[LaptopIndex:], LaptopDatas[LaptopIndex+1:])
	LaptopDatas[len(LaptopDatas)-1] = Laptop{}
	LaptopDatas = LaptopDatas[:len(LaptopDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"massage": fmt.Sprintf("laptop with data id %v has been delate", LaptopID),
	})

}
