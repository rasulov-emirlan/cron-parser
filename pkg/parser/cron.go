package parser

import (
	"fmt"
	"strings"
)

type CronCommand struct {
	Minute     []int
	Hour       []int
	DayOfMonth []int
	Month      []int
	DayOfWeek  []int
	Command    string
}

func NewParser() *CronCommand {
	return &CronCommand{
		Minute:     make([]int, 0),
		Hour:       make([]int, 0),
		DayOfMonth: make([]int, 0),
		Month:      make([]int, 0),
		DayOfWeek:  make([]int, 0),
		Command:    "",
	}
}

func (c *CronCommand) ParseAll(s string) error {
	cmd := strings.Fields(s)
	if len(cmd) < 6 {
		return ErrInvalidInput
	}

	if err := c.ParseMinute(cmd[0]); err != nil {
		return err
	}
	if err := c.ParseHour(cmd[1]); err != nil {
		return err
	}
	if err := c.ParseDayOfMonth(cmd[2]); err != nil {
		return err
	}
	if err := c.ParseMonth(cmd[3]); err != nil {
		return err
	}
	if err := c.ParseDayOfWeek(cmd[4]); err != nil {
		return err
	}
	c.Command = cmd[5]
	return nil
}

func (c CronCommand) String() string {
	minute := "minute"
	for _, i := range c.Minute {
		minute += fmt.Sprintf(" %d", i)
	}
	hour := "hour"
	for _, i := range c.Hour {
		hour += fmt.Sprintf(" %d", i)
	}
	dayOfMonth := "day of month"
	for _, i := range c.DayOfMonth {
		dayOfMonth += fmt.Sprintf(" %d", i)
	}
	month := "month"
	for _, i := range c.Month {
		month += fmt.Sprintf(" %d", i)
	}
	dayOfWeek := "day of week"
	for _, i := range c.DayOfWeek {
		dayOfWeek += fmt.Sprintf(" %d", i)
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\ncommand %s", minute, hour, dayOfMonth, month, dayOfWeek, c.Command)
}
