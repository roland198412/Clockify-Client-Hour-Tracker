package main

import (
	"clockify_client_hour_notifier/internal/client"
	"clockify_client_hour_notifier/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	debug := os.Getenv("DEBUG")

	debugMode, _ := strconv.ParseBool(debug)

	workspaceId := os.Getenv("WORKSPACE_ID")
	apiKey := os.Getenv("CLOCKIFY_API_KEY")
	clientName := os.Getenv("CLIENT_NAME")
	clientId := os.Getenv("CLIENT_ID")

	// Read environment variables outside of the client
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	emailRecipient := os.Getenv("EMAIL_RECIPIENT")

	// Parse integer environment variable
	monthlyHourLimitStr := os.Getenv("MONTHLY_HOUR_LIMIT")
	monthlyHourLimit, err := strconv.Atoi(monthlyHourLimitStr)
	if err != nil {
		panic("Invalid MONTHLY_HOUR_LIMIT value")
	}

	endPoint := fmt.Sprintf("https://reports.api.clockify.me/v1/workspaces/%s/reports/summary", workspaceId)

	now := time.Now()
	// First day of the next month
	firstOfNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
	beginningOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	// Subtract 1 second to get the last moment of the current month
	endOfMonth := firstOfNextMonth.Add(-time.Second)

	clockifyClient := client.NewClockifyClient(endPoint, apiKey, clientId)
	hoursWorkedInSeconds, err := clockifyClient.GetCurrentMonthHoursWorked(beginningOfMonth, endOfMonth)

	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	daysRemainingTillMonthEnd := utils.TotalWeekdaysBetweenTwoDates(now, endOfMonth)

	monthlyLimitInSeconds := utils.HoursToSeconds(monthlyHourLimit)

	remainingSeconds := monthlyLimitInSeconds - hoursWorkedInSeconds
	dailySecondAverage := remainingSeconds / int64(daysRemainingTillMonthEnd)

	if debugMode {
		fmt.Println("Now:          ", beginningOfMonth.Format(time.DateTime))
		fmt.Println("End of month: ", endOfMonth.Format(time.DateTime))
		fmt.Println(fmt.Sprintf("%-35s %s", "Period: %s - %s", beginningOfMonth, endOfMonth))
		fmt.Println(fmt.Sprintf("%-35s %s", "Client:", clientName))
		fmt.Println(fmt.Sprintf("%-35s %d hours", "Monthly Hour Limit:", monthlyHourLimit))
		fmt.Println(fmt.Sprintf("%-35s %d", "Days Remaining in Month:", daysRemainingTillMonthEnd))
		fmt.Println(fmt.Sprintf("%-35s %s", "Total Time Logged:", utils.FormatSecondsToHHMMSS(hoursWorkedInSeconds)))
		fmt.Println(fmt.Sprintf("%-35s %s", "Remaining Time Available:", utils.FormatSecondsToHHMMSS(remainingSeconds)))
		fmt.Println(fmt.Sprintf("%-35s %s", "Daily Target to Meet Goal:", utils.FormatSecondsToHHMMSS(dailySecondAverage)))
	}

	parsedSmtpPort, err := strconv.Atoi(smtpPort)
	smtpClient := client.NewSMTPClient(smtpHost, parsedSmtpPort, smtpUsername, smtpPassword)

	from := smtpUsername

	to := []string{emailRecipient}
	subject := fmt.Sprintf("Clockify Hour Report: %s", clientName)

	body := fmt.Sprintf(`
<html>
  <body>
    <h2>Monthly Hour Summary Report: %s %d - %s</h2>
    <table cellpadding="4" cellspacing="4" border="1">
      <tr><td><strong>Period:</strong></td><td>%s - %s</td></tr>
      <tr><td><strong>Client:</strong></td><td>%s</td></tr>
      <tr><td><strong>Monthly Hour Limit:</strong></td><td>%d hours</td></tr>
      <tr><td><strong>Days Remaining in Month:</strong></td><td>%d</td></tr>
      <tr><td><strong>Total Time Logged:</strong></td><td>%s</td></tr>
      <tr><td><strong>Remaining Time Available:</strong></td><td>%s</td></tr>
      <tr><td><strong>Daily Target to Meet Goal:</strong></td><td>%s</td></tr>
    </table>
  </body>
</html>`,
		time.Now().Month(),
		time.Now().Year(),
		clientName,
		beginningOfMonth.Format(time.DateTime),
		endOfMonth.Format(time.DateTime),
		clientName,
		monthlyHourLimit,
		daysRemainingTillMonthEnd,
		utils.FormatSecondsToHHMMSS(hoursWorkedInSeconds),
		utils.FormatSecondsToHHMMSS(remainingSeconds),
		utils.FormatSecondsToHHMMSS(dailySecondAverage),
	)

	if err := smtpClient.Send(from, to, subject, body, "text/html"); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	if debugMode {
		fmt.Println("Email sent successfully!")
	}
}
