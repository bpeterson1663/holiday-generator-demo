# Holiday Generator

The purpose of this app is to be able to calculate the date of any holiday (given the specific rule it follows; 4th Thursday of November, 1st Monday of September, 1st day of January, etc)
Application is hosted on Heroku https://federal-holiday-generator.herokuapp.com/rules (Free version is used so it may take a few minutes to start if application has not been visited in a while)

## GET /api/holidays
### Query Params
 Required
 - start: start year for range of holidays
 - end: end year for range of holidays
### Response Body
```
[
    {
        "name": string,
        "year": int,
        "month": string,
        "day": int,
        "day_of_week": string
        
    }
]
```