package main

import (
	"fmt"
	"net/http"
	"time"

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
	{ID: "1", Name: "New Year's Day", IsFixed: true, MonthOfOccurrence: 1, DayOfMonth: 1, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "2", Name: "Martin Luther King Day", IsFixed: false, MonthOfOccurrence: 1, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 3},
	{ID: "3", Name: "Presidents' Day", IsFixed: false, MonthOfOccurrence: 2, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 3},
	{ID: "4", Name: "Memorial Day", IsFixed: false, MonthOfOccurrence: 5, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 5},
	{ID: "5", Name: "Juneteenth", IsFixed: true, MonthOfOccurrence: 6, DayOfMonth: 19, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "6", Name: "Independence Day", IsFixed: true, MonthOfOccurrence: 7, DayOfMonth: 4, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "7", Name: "Labor Day", IsFixed: false, MonthOfOccurrence: 9, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 1},
	{ID: "8", Name: "Columbus/Indigenous Peoples Day", IsFixed: false, MonthOfOccurrence: 10, DayOfMonth: 0, DayOfWeek: 1, WeekOfMonth: 2},
	{ID: "9", Name: "Veterans Day", IsFixed: true, MonthOfOccurrence: 11, DayOfMonth: 11, DayOfWeek: 0, WeekOfMonth: 0},
	{ID: "10", Name: "Thanksgiving Day", IsFixed: false, MonthOfOccurrence: 11, DayOfMonth: 0, DayOfWeek: 4, WeekOfMonth: 4},
	{ID: "11", Name: "Christmas Day", IsFixed: true, MonthOfOccurrence: 12, DayOfMonth: 25, DayOfWeek: 4, WeekOfMonth: 4},
}

func getRules(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rules)
}

func getRuleByID(c *gin.Context) {
	id := c.Param("id")

	for _, rule := range rules {
		if rule.ID == id {
			c.IndentedJSON(http.StatusOK, rule)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rule not found"})
}

func getNextHolidayByRuleID(id string) {

	var holidayRule rule

	for _, rule := range rules {
		if rule.ID == id {
			holidayRule = rule
		}
	}
	if holidayRule.IsFixed {
		d := time.Date(2020, time.Month(holidayRule.MonthOfOccurrence), holidayRule.DayOfMonth, 0, 0, 0, 0, time.UTC)
		year, month, day := d.Date()

		fmt.Printf("year = %v\n", year)
		fmt.Printf("month = %v\n", month)
		fmt.Printf("day = %v\n", day)
	} else {
		weekCount := 0
		d := time.Date(2020, time.Month(holidayRule.MonthOfOccurrence), 1, 0, 0, 0, 0, time.UTC)
		for {
			if int(d.Weekday()) == holidayRule.DayOfWeek {
				weekCount = weekCount + 1
				if weekCount == holidayRule.WeekOfMonth {
					break
				} else {
					d = d.AddDate(0, 0, 1)
					if int(d.Month()) != holidayRule.MonthOfOccurrence {
						d = d.AddDate(0, 0, -7)
						break
					}
				}
			} else {
				d = d.AddDate(0, 0, 1)
				if int(d.Month()) != holidayRule.MonthOfOccurrence {
					d = d.AddDate(0, 0, -7)
					break
				}
			}

		}
		year, month, day := d.Date()

		fmt.Printf("year = %v\n", year)
		fmt.Printf("month = %v\n", month)
		fmt.Printf("day = %v\n", day)
	}
}

func postRule(c *gin.Context) {
	var newRule rule

	if err := c.BindJSON(&newRule); err != nil {
		return
	}

	rules = append(rules, newRule)
	c.IndentedJSON(http.StatusCreated, newRule)
}

func main() {
	getNextHolidayByRuleID("1")
	getNextHolidayByRuleID("8")
	router := gin.Default()
	router.GET("/rules", getRules)
	router.GET("/rules/:id", getRuleByID)
	router.POST("/rules", postRule)
	router.Run("localhost:8080")
}
