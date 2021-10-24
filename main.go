package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type rule struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	IsFixed           bool   `json:"is_fixed"`
	MonthOfOccurrence int    `json:"month_of_occurrence"`
	DayOfMonth        int    `json:"day_of_month"`
	DayOfWeek         int    `json:"day_of_week"`
	WeekOfMonth       int    `json:"week_of_month"`
}

var rules = []rule{
	{ID: "1", Name: "New Year's Day", IsFixed: true, MonthOfOccurrence: 0, DayOfMonth: 1, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "2", Name: "Martin Luther King Day", IsFixed: false, MonthOfOccurrence: 0, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 3},
	{ID: "3", Name: "Presidents' Day", IsFixed: false, MonthOfOccurrence: 0, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 3},
	{ID: "4", Name: "Memorial Day", IsFixed: false, MonthOfOccurrence: 1, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 5},
	{ID: "5", Name: "Juneteenth", IsFixed: true, MonthOfOccurrence: 5, DayOfMonth: 19, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "6", Name: "Independence Day", IsFixed: true, MonthOfOccurrence: 6, DayOfMonth: 4, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "7", Name: "Labor Day", IsFixed: false, MonthOfOccurrence: 8, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 1},
	{ID: "8", Name: "Columbus/Indigenous Peoples Day", IsFixed: false, MonthOfOccurrence: 9, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 2},
	{ID: "9", Name: "Veterans Day", IsFixed: true, MonthOfOccurrence: 10, DayOfMonth: 11, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "10", Name: "Thanksgiving Day", IsFixed: false, MonthOfOccurrence: 10, DayOfMonth: 0, DayOfWeek: 4, WeekOfMonth: 4},
	{ID: "11", Name: "Christmas Day", IsFixed: true, MonthOfOccurrence: 11, DayOfMonth: 25, DayOfWeek: 4, WeekOfMonth: 4},
}

func getRules(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rules)
}

func main() {
	router := gin.Default()
	router.GET("/rules", getRules)
	router.Run("localhost:8080")
}
