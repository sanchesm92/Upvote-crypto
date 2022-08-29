package controllers

import (
	"log"
	"sort"
	"github.com/gin-gonic/gin"
	"strconv"
	"upvote-crypto/database/models"
	"upvote-crypto/database"
)

func ShowCurrencies(c *gin.Context) {
	db := database.GetDatabase()
	var currencies []models.Currency
	err := db.Find(&currencies).Error
	if currencies != nil {
		c.JSON(200, currencies)
	}
	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot list currencies: " + err.Error(),
		})
		return
	}
}

type Votes struct {
	Name string
	Vote int
}

func validateElement(s []Votes, e Votes) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

func ShowRanking(c *gin.Context) {
	db := database.GetDatabase()
	var currencies []models.Currency
	err := db.Find(&currencies).Error
	log.Print(currencies)

	var allCurrencies []string
	for _, s := range currencies {
		allCurrencies = append(allCurrencies, s.Name)
	}

	var result []Votes

	for _, s := range allCurrencies {
		acc := 0
		for _, b := range currencies {
			if b.Name == s  {
				acc += 1
			}
		}
		var curr Votes
		curr.Name = s
		curr.Vote = acc
		if !validateElement(result, curr) {
			result = append(result, curr)
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
    return result[i].Vote > result[j].Vote
})

	if currencies != nil {
		c.JSON(200, result)
	}
	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot list currencies: " + err.Error(),
		})
		return
	}
}

func CreateVote(c *gin.Context) {
	db := database.GetDatabase()

	var currency models.Currency

	err := c.ShouldBindJSON(&currency)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "cannot bind JSON: " + err.Error(),
			})
			return
	}
	err = db.Create(&currency).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create vote: " + err.Error(),
		})
		return
	}
	c.JSON(201, currency)
}

func EditCurrency(c *gin.Context) {

	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()

	var currency models.Currency

	err = c.ShouldBindJSON(&currency)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Model(&models.Currency{}).Where("id = ?", newid).Updates(models.Currency{Name: currency.Name, Vote: currency.Vote}).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create currency: " + err.Error(),
		})
		return
	}

	c.JSON(200, currency)
}


func DeleteCurrency(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Currency{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete currency: " + err.Error(),
		})
		return
	}

	c.Status(204)
}