package controller

import(
	. "../models"
	"../database"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

func PeopleRouting(router *gin.RouterGroup) {
	router.GET("/", GetPeople)
	router.GET("/:id", GetPerson)
	router.POST("/", CreatePerson)
	router.PUT("/:id", UpdatePerson)
	router.DELETE("/:id", DeletePerson)
}

// DeletePerson godoc
// @Summary Delete a person
// @Description Delete person by ID
// @ID int
// @Tags People
// @Accept  json
// @Produce  json
// @Param  id path int true "Person ID"
// @Success 204 {object} models.Person
// @Router /people/{id} [delete]
func DeletePerson(ctx *gin.Context) {
	db := database.Open()
	id := ctx.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	ctx.JSON(204, gin.H{"id #" + id: "deleted"})
	defer db.Close()
}

// UpdatePerson godoc
// @Summary Update a person
// @Description Update person by ID and by json
// @ID int
// @Tags People
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Param person body models.Person true "Update person"
// @Success 200 {object} models.Person
// @Router /people/{id} [put]
func UpdatePerson(ctx *gin.Context) {

	var person Person
	id := ctx.Params.ByName("id")
	aid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		ctx.JSON(500, person)
	}

	db := database.Open()
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {		
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	}
	ctx.BindJSON(&person)
	person.ID = aid

	db.Save(&person)
	ctx.JSON(200, person)
	defer db.Close()
}

// CreatePerson godoc
// @Summary Create a person
// @Description Create person by json
// @Tags People
// @Accept  json
// @Produce  json
// @Param  person body models.Person true "Create person"
// @Success 200 {object} models.Person
// @Router /people [post]
func CreatePerson(ctx *gin.Context) {
	db := database.Open()
	var person Person
	ctx.BindJSON(&person)

	db.Create(&person)
	ctx.JSON(200, person)
	defer db.Close()
}


// ShowPerson godoc
// @Summary Show a person
// @Description get person by ID
// @Tags People
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 {object} models.Person
// @Router /people/{id} [get]
func GetPerson(ctx *gin.Context) {
	db := database.Open()
	id := ctx.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, person)
	}
	defer db.Close()
}

// ListPeople godoc
// @Summary List people
// @Description get people
// @Tags People
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Person
// @Router /people/ [get]
func GetPeople(ctx *gin.Context) {
	db := database.Open()
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		ctx.JSON(200, people)
	}
	defer db.Close()
}