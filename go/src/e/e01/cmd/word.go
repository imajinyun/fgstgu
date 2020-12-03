package cmd

import (
	"log"
	"strings"

	"e/e01/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeToUpper = iota + 1
	ModeToLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeWordToSnakeCase
)

var str string
var mode int8
var desc = strings.Join([]string{
	"The word command supported modes are as follows:",
	"1. Word to upper case",
	"2. Word to lower case",
	"3. Underscore word to upper camel case",
	"4. Underscore word to lower camel case",
	"5. Convert a word to snake case",
}, "\n")
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Conversion of word case, hump, etc",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeToUpper:
			content = word.ToUpper(str)
		case ModeToLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeWordToSnakeCase:
			content = word.ToSnakeCase(str)
		default:
			log.Fatalf("\nNot supported mode, please see [help word] documentation")
		}
		log.Printf("\nOutput: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "Please input word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Please input mode")
}
