package main

import (
	"fmt"
	"strings"
)

func uptime(total int) string {
	days := total / 86400
	total = total % 86400
	hours := total / 3600
	total = total % 3600
	minutes := total / 60
	total = total % 60
	seconds := total

	result := ""
	if days > 0 {
		result += fmt.Sprintf("%dд ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dч ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%dм ", minutes)
	}
	if seconds > 0 {
		result += fmt.Sprintf("%dс ", seconds)
	}
	if result == "" {
		return "0с"
	}
	return strings.TrimSuffix(result, " ")
}

func main() {
	fmt.Println(uptime(93784))
	fmt.Println(uptime(86405))
	fmt.Println(uptime(3600))
	fmt.Println(uptime(59))
	fmt.Println(uptime(0))
}
