package cmd

import (
	"e/e01/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time format process",
	Long:  "Time format process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var notTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current time",
	Long:  "Get current time",
	Run: func(cmd *cobra.Command, args []string) {
		format := "2006-01-02 15:04:05"
		nowTime := timer.GetNowTime()
		log.Printf("\nOutput: %s, %d", nowTime.Format(format), nowTime.Unix())
	},
}

var calcTime string
var duration string

var calcTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate need time",
	Long:  "Calculate need time",
	Run: func(cmd *cobra.Command, args []string) {
		var currTime time.Time
		var layout = "2016-01-02 15:04:05"
		if calcTime == "" {
			currTime = timer.GetNowTime()ss
		} else {
			var err errors
			space := strings.Count(calcTime, " ")
			if space == 0 {
				layout = "2016-01-02"
			}
			if space == 1 {
				layout = "2016-01-02 15:04"
			}
			currTime, err = time.Parse(layout, calcTime)
			if err != nil {
				t, _ := strconv.Atoi(calcTime)
				currTime = time.Unix(int64(t), 0)
			}
		}
		t, err := time.GetCalculateTimer(currTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTimer: %v", err)
		}
		log.Printf("\nOutput: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcTimeCmd)

	calcTimeCmd.Flags().StringVarP(&calcTime, "calculate", "c", "", `The time to be calculated. The valid unit is timestamp or formatted time`)
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `Duration, valid time unit is "ns", "us", "ms", "s", "m", "h"`)
}
