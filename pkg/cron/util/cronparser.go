package util

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type parsedCronField struct {
	wildCard bool // if false, the value will be used as a number
	value    int  // if wildCard is true, this value is ignored if it's -1
	valid    bool // if false, the value is not valid
}

type cronRule struct {
	regex    string
	minValue int
	maxValue int
}

type CronField struct {
	Field    string // refer to the consts
	Value    string
	RunEvery int // -1 means "*"
	rule     cronRule
}

const (
	CRON_FIELD_MINUTE  = "minute"
	CRON_FIELD_HOUR    = "hour"
	CRON_FIELD_DAY     = "day"
	CRON_FIELD_MONTH   = "month"
	CRON_FIELD_WEEKDAY = "weekday"
)

// @TODO handle the case where the a number is used in a range with '/{number}'
const (
	cronFieldMinuteRegex  = "^((\\*)|(\\d{1,2})|(\\*\\/\\d{1,2}))$"
	cronFieldHourRegex    = "^((\\*)|(\\d{1,2})|(\\*\\/\\d{1,2}))$"
	cronFieldDayRegex     = "^((\\*)|(\\d{1,2})|(\\*\\/\\d{1,2}))$"
	cronFieldMonthRegex   = "^((\\*)|(\\d{1,2})|(\\*\\/\\d{1,2}))$"
	cronFieldWeekdayRegex = "^((\\*)|(\\d{1})|(\\*\\/\\d{1}))$"
)

var (
	cronrules_minute = cronRule{
		regex:    cronFieldMinuteRegex,
		minValue: 0, // 0 is treated as off, -1 is treated as "*"
		maxValue: 59,
	}
	cronrules_hour = cronRule{
		regex:    cronFieldHourRegex,
		minValue: 0, // 0 is treated as off, -1 is treated as "*"
		maxValue: 23,
	}
	cronrules_day = cronRule{
		regex:    cronFieldDayRegex,
		minValue: 1, // 1 is treated as off, -1 is treated as "*"
		maxValue: 31,
	}
	cronrules_month = cronRule{
		regex:    cronFieldMonthRegex,
		minValue: 1, // 1 is treated as off, -1 is treated as "*"
		maxValue: 12,
	}
	cronrules_weekday = cronRule{
		regex:    cronFieldWeekdayRegex,
		minValue: 0, // 0 is treated as off, -1 is treated as "*"
		maxValue: 6,
	}
)

func parseCronField(value string, rule cronRule) parsedCronField {
	field := parsedCronField{
		valid: false,
	}
	//check against regexp
	matched, err := regexp.MatchString(rule.regex, value)
	if err != nil {
		return field
	}
	if !matched {
		return field
	}

	numValue := 0
	//check against if value is strictly a '*' wildcard
	if value == "*" {
		field.wildCard = true
		field.value = -1
		return field
	} else if strings.Contains(value, "*/") {
		//split the value by '/' and check the second part
		split := strings.Split(value, "/")
		if len(split) != 2 {
			return field
		}
		//check if the second part is a number
		if value, err := strconv.Atoi(split[1]); err != nil {
			return field
		} else {
			field.wildCard = true
			numValue = value
		}
	} else {
		//check if the value is a number
		if value, err := strconv.Atoi(value); err != nil {
			return field
		} else {
			numValue = value
		}
	}

	//check if the value is in the range of min and max values
	if numValue < rule.minValue || numValue > rule.maxValue {
		return field
	}

	field.value = numValue
	return field
}

func ParseCronTime(cronTime string) ([]CronField, error) {
	cronObject := []CronField{}
	// Implement the logic to validate the cron time string
	// Cron strings must be in the format of "minute hour day month weekday"
	// where each field can be a number, a range, or a wildcard (*)

	// For example, "*/5 * * * *" means every 5 minutes
	// You can use regular expressions or string manipulation to validate the format
	cronFields := []string{"minute", "hour", "day", "month", "weekday"}
	inputSplit := strings.Split(cronTime, " ")
	if len(inputSplit) != len(cronFields) {
		return nil, errors.New("Invalid cron time format")
	}

	for i, field := range cronFields {
		var rule cronRule
		var fieldName string
		switch field {
		case "minute":
			rule = cronrules_minute
			fieldName = CRON_FIELD_MINUTE
		case "hour":
			rule = cronrules_hour
			fieldName = CRON_FIELD_HOUR
		case "day":
			rule = cronrules_day
			fieldName = CRON_FIELD_DAY
		case "month":
			rule = cronrules_month
			fieldName = CRON_FIELD_MONTH
		case "weekday":
			rule = cronrules_weekday
			fieldName = CRON_FIELD_WEEKDAY
		default:
			return nil, errors.New("Invalid cron field")
		}
		parsedCronField := parseCronField(inputSplit[i], rule)
		if !parsedCronField.valid {
			return nil, errors.New("Invalid cron time format")
		}

		cronObject = append(cronObject, CronField{
			Field:    fieldName,
			Value:    inputSplit[i],
			RunEvery: parsedCronField.value,
		})

	}

	return cronObject, nil
}
