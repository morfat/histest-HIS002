package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountyIn struct {
	Name string `json:"name"`
}

func postCountyHandler(ctx *gin.Context) {

	var objIn CountyIn

	if err := ctx.ShouldBindJSON(&objIn); err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	log.Println(objIn)

}

type PatientIn struct {
	CountyId          int32  `json:"county_id"`
	FirstName         string `json:"first_name"`
	SurName           string `json:"surname"`
	OtherName         string `json:"other_name"`
	PatientNumber     string `json:"patient_number"`
	MobileNumber      string `json:"mobile_number"`
	EmailAddress      string `json:"email_address"`
	IsDisabled        bool   `json:"is_disabled"`
	Gender            string `json:"gender"`
	ContactPersonName string `json:"contact_person_name"`
	ContactPersonTel  string `json:"contact_person_tel"`
}

func postPatientHandler(ctx *gin.Context) {

	var objIn PatientIn

	if err := ctx.ShouldBindJSON(&objIn); err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	log.Println(objIn)

}
