package parser

import (
	"regexp"
	"strconv"
	"strings"
)

func (c *CronCommand) ParseMinute(s string) error {
	if s == "*" {
		for i := 0; i < 60; i++ {
			c.Minute = append(c.Minute, i)
		}
		return nil
	}

	c.Minute = make([]int, 0)

	var (
		withLow = false

		nums = []string{}

		i = 0
		j = 0
	)

	re := regexp.MustCompile("[0-9]+")
	nums = re.FindAllString(s, -1)
	if len(nums) == 0 {
		return ErrInvalidMinutes
	}

	if len(nums) == 1 && !strings.ContainsAny(s, ",/*-") {
		n, _ := strconv.Atoi(nums[0])
		c.Minute = append(c.Minute, n)
	}

	for ; i < len(s); i++ {
		if s[i] == '*' {
			withLow = true
		}

		if s[i] == '/' {
			if j >= len(nums) {
				return ErrInvalidMinutes
			}

			start := 1
			if withLow {
				start = 0
			}
			// error here can be ignored since we got nums
			// through regexp and we know for a fact that it
			// contains only numbers
			var n int
			if j == len(nums)-1 {
				n, _ = strconv.Atoi(nums[j])
			} else {
				n, _ = strconv.Atoi(nums[j+1])
			}
			if len(c.Minute) != 0 {
				start = c.Minute[len(c.Minute)-1] + n
			}

			for ; start < 60; start += n {
				c.Minute = append(c.Minute, start)
			}

			// never forget to increment j cause it show which
			// number to use in the current pattern
			j++
			if j == len(nums) {
				break
			}
			i += len(nums[j])
		}

		if s[i] == '-' {
			// '-' cannot be the first character since it needs two numbers
			// from both sides
			if i == 0 || j >= len(nums) {
				return ErrInvalidMinutes
			}

			first, _ := strconv.Atoi(nums[j])
			second, _ := strconv.Atoi(nums[j+1])
			if len(c.Minute) != 0 {
				first = c.Minute[len(c.Minute)-1] + 1
			}

			for ; first <= second; first++ {
				c.Minute = append(c.Minute, first)
			}
			i += len(nums[j])
		}

		if s[i] == ',' {
			if j >= len(nums) || i == 0 {
				return ErrInvalidMinutes
			}
			i--
			for i < len(s) {
				if s[i] == ',' {
					i++
					continue
				}
				if s[i] == '-' || s[i] == '/' || s[i] == '*' {
					i--
					j--
					break
				}

				n, _ := strconv.Atoi(nums[j])
				c.Minute = append(c.Minute, n)
				j++
				if j == len(nums) {
					break
				}
				i += len(nums[j])
			}
		}
	}
	return nil
}

func (c *CronCommand) ParseHour(s string) error {
	if s == "*" {
		for i := 0; i < 24; i++ {
			c.Hour = append(c.Hour, i)
		}
		return nil
	}

	c.Hour = make([]int, 0)

	var (
		withLow = false

		nums = []string{}

		i = 0
		j = 0
	)

	re := regexp.MustCompile("[0-9]+")
	nums = re.FindAllString(s, -1)
	if len(nums) == 0 {
		return ErrInvalidHours
	}

	if len(nums) == 1 && !strings.ContainsAny(s, ",/*-") {
		n, _ := strconv.Atoi(nums[0])
		c.Hour = append(c.Hour, n)
	}

	for ; i < len(s); i++ {
		if s[i] == '*' {
			withLow = true
		}

		if s[i] == '/' {
			if j >= len(nums) {
				return ErrInvalidHours
			}

			start := 1
			if withLow {
				start = 0
			}
			// error here can be ignored since we got nums
			// through regexp and we know for a fact that it
			// contains only numbers
			n, _ := strconv.Atoi(nums[j+1])

			if len(c.Hour) != 0 {
				start = c.Hour[len(c.Hour)-1] + n
			}

			for ; start < 12; start += n {
				c.Hour = append(c.Hour, start)
			}

			// never forget to increment j cause it show which
			// number to use in the current pattern
			j++
			if j == len(nums) {
				break
			}
			i += len(nums[j])
		}

		if s[i] == '-' {
			// '-' cannot be the first character since it needs two numbers
			// from both sides
			if i == 0 || j >= len(nums) {
				return ErrInvalidHours
			}

			first, _ := strconv.Atoi(nums[j])
			second, _ := strconv.Atoi(nums[j+1])
			if len(c.Hour) != 0 {
				first = c.Hour[len(c.Hour)-1] + 1
			}

			for ; first <= second; first++ {
				c.Hour = append(c.Hour, first)
			}
			i += len(nums[j])
		}

		if s[i] == ',' {
			if j >= len(nums) || i == 0 {
				return ErrInvalidHours
			}
			i--
			for i < len(s) {
				if s[i] == ',' {
					i++
					continue
				}
				if s[i] == '-' || s[i] == '/' || s[i] == '*' {
					i--
					j--
					break
				}

				n, _ := strconv.Atoi(nums[j])
				c.Hour = append(c.Hour, n)
				j++
				if j == len(nums) {
					break
				}
				i += len(nums[j])
			}
		}
	}
	return nil
}

func (c *CronCommand) ParseDayOfMonth(s string) error {
	if s == "*" {
		for i := 1; i < 32; i++ {
			c.DayOfMonth = append(c.DayOfMonth, i)
		}
		return nil
	}

	c.DayOfMonth = make([]int, 0)

	var (
		withLow = false

		nums = []string{}

		i = 0
		j = 0
	)

	re := regexp.MustCompile("[0-9]+")
	nums = re.FindAllString(s, -1)
	if len(nums) == 0 {
		return ErrInvalidDayOfMonth
	}

	if len(nums) == 1 && !strings.ContainsAny(s, ",/*-") {
		n, _ := strconv.Atoi(nums[0])
		c.DayOfMonth = append(c.DayOfMonth, n)
	}

	for ; i < len(s); i++ {
		if s[i] == '*' {
			withLow = true
		}

		if s[i] == '/' {
			if j >= len(nums) {
				return ErrInvalidDayOfMonth
			}

			start := 1
			if withLow {
				start = 0
			}
			// error here can be ignored since we got nums
			// through regexp and we know for a fact that it
			// contains only numbers
			n, _ := strconv.Atoi(nums[j+1])

			if len(c.DayOfMonth) != 0 {
				start = c.DayOfMonth[len(c.DayOfMonth)-1] + n
			}

			for ; start < 32; start += n {
				c.DayOfMonth = append(c.DayOfMonth, start)
			}

			// never forget to increment j cause it show which
			// number to use in the current pattern
			j++
			if j == len(nums) {
				break
			}
			i += len(nums[j])
		}

		if s[i] == '-' {
			// '-' cannot be the first character since it needs two numbers
			// from both sides
			if i == 0 || j >= len(nums) {
				return ErrInvalidDayOfMonth
			}

			first, _ := strconv.Atoi(nums[j])
			second, _ := strconv.Atoi(nums[j+1])
			if len(c.DayOfMonth) != 0 {
				first = c.DayOfMonth[len(c.DayOfMonth)-1] + 1
			}

			for ; first <= second; first++ {
				c.DayOfMonth = append(c.DayOfMonth, first)
			}
			i += len(nums[j])
		}

		if s[i] == ',' {
			if j >= len(nums) || i == 0 {
				return ErrInvalidDayOfMonth
			}
			i--
			for i < len(s) {
				if s[i] == ',' {
					i++
					continue
				}
				if s[i] == '-' || s[i] == '/' || s[i] == '*' {
					i--
					j--
					break
				}

				n, _ := strconv.Atoi(nums[j])
				c.DayOfMonth = append(c.DayOfMonth, n)
				j++
				if j == len(nums) {
					break
				}
				i += len(nums[j])
			}
		}
	}
	return nil
}

func (c *CronCommand) ParseMonth(s string) error {
	if s == "*" {
		for i := 1; i < 13; i++ {
			c.Month = append(c.Month, i)
		}
		return nil
	}

	c.Month = make([]int, 0)

	var (
		withLow = false

		nums = []string{}

		i = 0
		j = 0
	)

	re := regexp.MustCompile("[0-9]+")
	nums = re.FindAllString(s, -1)
	if len(nums) == 0 {
		return ErrInvalidMonth
	}

	if len(nums) == 1 && !strings.ContainsAny(s, ",/*-") {
		n, _ := strconv.Atoi(nums[0])
		c.DayOfMonth = append(c.DayOfMonth, n)
	}

	for ; i < len(s); i++ {

		if s[i] == '*' {
			withLow = true
		}

		if s[i] == '/' {
			if j >= len(nums) {
				return ErrInvalidMonth
			}

			start := 1
			if withLow {
				start = 0
			}
			// error here can be ignored since we got nums
			// through regexp and we know for a fact that it
			// contains only numbers
			n, _ := strconv.Atoi(nums[j+1])

			if len(c.Month) != 0 {
				start = c.Month[len(c.Month)-1] + n
			}

			for ; start < 13; start += n {
				c.Month = append(c.Month, start)
			}

			// never forget to increment j cause it show which
			// number to use in the current pattern
			j++
			if j == len(nums) {
				break
			}
			i += len(nums[j])
		}

		if s[i] == '-' {
			// '-' cannot be the first character since it needs two numbers
			// from both sides
			if i == 0 || j >= len(nums) {
				return ErrInvalidMonth
			}

			first, _ := strconv.Atoi(nums[j])
			second, _ := strconv.Atoi(nums[j+1])
			if len(c.Month) != 0 {
				first = c.Month[len(c.Month)-1] + 1
			}

			for ; first <= second; first++ {
				c.Month = append(c.Month, first)
			}
			i += len(nums[j])
		}

		if s[i] == ',' {
			if j >= len(nums) || i == 0 {
				return ErrInvalidMonth
			}
			i--
			for i < len(s) {
				if s[i] == ',' {
					i++
					continue
				}
				if s[i] == '-' || s[i] == '/' || s[i] == '*' {
					i--
					j--
					break
				}

				n, _ := strconv.Atoi(nums[j])
				c.Month = append(c.Month, n)
				j++
				if j == len(nums) {
					break
				}
				i += len(nums[j])
			}
		}
	}
	return nil
}

func (c *CronCommand) ParseDayOfWeek(s string) error {
	if s == "*" {
		for i := 0; i < 7; i++ {
			c.DayOfWeek = append(c.DayOfWeek, i)
		}
		return nil
	}

	c.DayOfWeek = make([]int, 0)

	var (
		withLow = false

		nums = []string{}

		i = 0
		j = 0
	)

	re := regexp.MustCompile("[0-9]+")
	nums = re.FindAllString(s, -1)
	if len(nums) == 0 {
		return ErrInvalidDayOfWeek
	}

	if len(nums) == 1 && !strings.ContainsAny(s, ",/*-") {
		n, _ := strconv.Atoi(nums[0])
		c.DayOfWeek = append(c.DayOfWeek, n)
	}

	for ; i < len(s); i++ {
		if s[i] == '*' {
			withLow = true
		}

		if s[i] == '/' {
			if j >= len(nums) {
				return ErrInvalidDayOfWeek
			}

			start := 1
			if withLow {
				start = 0
			}
			// error here can be ignored since we got nums
			// through regexp and we know for a fact that it
			// contains only numbers
			n, _ := strconv.Atoi(nums[j+1])

			if len(c.DayOfWeek) != 0 {
				start = c.DayOfWeek[len(c.DayOfWeek)-1] + n
			}

			for ; start < 7; start += n {
				c.DayOfWeek = append(c.DayOfWeek, start)
			}

			// never forget to increment j cause it show which
			// number to use in the current pattern
			j++
			i += len(nums[j])
		}

		if s[i] == '-' {
			// '-' cannot be the first character since it needs two numbers
			// from both sides
			if i == 0 || j >= len(nums) {
				return ErrInvalidDayOfWeek
			}

			first, _ := strconv.Atoi(nums[j])
			second, _ := strconv.Atoi(nums[j+1])
			if len(c.DayOfWeek) != 0 {
				first = c.DayOfWeek[len(c.DayOfWeek)-1] + 1
			}

			for ; first <= second; first++ {
				c.DayOfWeek = append(c.DayOfWeek, first)
			}
			i += len(nums[j])
		}

		if s[i] == ',' {
			if j >= len(nums) || i == 0 {
				return ErrInvalidDayOfWeek
			}
			i--
			for i < len(s) {
				if s[i] == ',' {
					i++
					continue
				}
				if s[i] == '-' || s[i] == '/' || s[i] == '*' {
					i--
					j--
					break
				}

				n, _ := strconv.Atoi(nums[j])
				c.DayOfWeek = append(c.DayOfWeek, n)
				j++
				if j == len(nums) {
					break
				}
				i += len(nums[j])
			}
		}
	}

	return nil
}
