package main

import (
	"net/http"
	"os"
	"strconv"
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

type holiday struct {
	Name      string `json:"name"`
	Month     string `json:"month"`
	Year      int    `json:"year"`
	Day       int    `json:"day"`
	DayOfWeek string `json:"day_of_week"`
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

func generateFixedHoliday(rule rule, year int) holiday {
	var holiday holiday
	d := time.Date(year, time.Month(rule.MonthOfOccurrence), rule.DayOfMonth, 0, 0, 0, 0, time.UTC)
	year, month, day := d.Date()
	holiday.Name = rule.Name
	holiday.Year = year
	holiday.Month = month.String()
	holiday.Day = day
	holiday.DayOfWeek = d.Weekday().String()
	return holiday
}

func generateNonFixedHoliday(rule rule, year int) holiday {
	var holiday holiday
	weekCount := 0
	d := time.Date(year, time.Month(rule.MonthOfOccurrence), 1, 0, 0, 0, 0, time.UTC)
	//Starting from the first of the month the rule occurs
	for {
		// Check if the day of the week is equal to the day of the week the rule occurs
		if int(d.Weekday()) == rule.DayOfWeek {
			//Incrememt week count and check if it is equal to the week of the rule
			weekCount = weekCount + 1
			if weekCount == rule.WeekOfMonth {
				// If it is, then the day of the week and the week of the month match the rule and we have found our date
				break
			} else {
				d = d.AddDate(0, 0, 1)
				if int(d.Month()) != rule.MonthOfOccurrence {
					d = d.AddDate(0, 0, -7)
				}
			}
		} else {
			//If we are not on the correct day of the week, increment a day
			d = d.AddDate(0, 0, 1)
			//Check if we are outside the month it occurs, if so subtract a week
			if int(d.Month()) != rule.MonthOfOccurrence {
				d = d.AddDate(0, 0, -7)
			}
		}
	}
	year, month, day := d.Date()
	holiday.Name = rule.Name
	holiday.Year = year
	holiday.Month = month.String()
	holiday.Day = day
	holiday.DayOfWeek = d.Weekday().String()
	return holiday
}

func getAllHolidays(c *gin.Context) {
	var holidays []holiday
	start, startErr := strconv.Atoi(c.Query("start"))
	if startErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "no start provided")
		return
	}
	end, endErr := strconv.Atoi(c.Query("end"))
	if endErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "no end provided")
		return
	}
	for {
		if start <= end {
			for _, rule := range rules {
				var holiday holiday
				if rule.IsFixed {
					holiday = generateFixedHoliday(rule, start)

				} else {
					holiday = generateNonFixedHoliday(rule, start)

				}
				holidays = append(holidays, holiday)
			}
			start = start + 1
		} else {
			break
		}
	}
	c.IndentedJSON(http.StatusCreated, holidays)
}

func getNextHolidayByRuleID(c *gin.Context) {
	id := c.Param("id")
	start, startErr := strconv.Atoi(c.Query("start"))
	if startErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "no start provided")
		return
	}
	end, endErr := strconv.Atoi(c.Query("end"))
	if endErr != nil {
		c.IndentedJSON(http.StatusBadRequest, "no end provided")
		return
	}
	var holidayRule rule
	var holiday holiday
	for _, rule := range rules {
		if rule.ID == id {
			holidayRule = rule
		}
	}
	for {
		if start <= end {
			if holidayRule.IsFixed {
				holiday = generateFixedHoliday(holidayRule, start)
				c.IndentedJSON(http.StatusCreated, holiday)
			} else {
				holiday = generateNonFixedHoliday(holidayRule, start)
				c.IndentedJSON(http.StatusCreated, holiday)
			}
		} else {
			break
		}
		start = start + 1
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/rules", getRules)
	router.GET("/rules/:id", getRuleByID)
	router.POST("/rules", postRule)
	router.GET("/holiday/:id", getNextHolidayByRuleID)
	router.GET("/holidays", getAllHolidays)
	router.Run(":" + port)
}
