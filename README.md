# ⏱️ Clockify Client Hour Tracker

A simple tool to track and report progress toward monthly client hour goals using the [Clockify API](https://app.clockify.me).

## 📌 Overview

This CLI application helps freelancers and contractors monitor how many hours they’ve worked for a specific client within a given month and calculates how many hours per day are needed to meet their monthly time commitment. It sends a daily/ CRON specific summary email to keep you on track.

### ✅ Example Use Case

If a client allocates **50 hours per month**, and there are **21 workdays** in the month:

- If you've logged **10 hours** by the 1st, the app will notify you that you now need to average **1.2 hours/day** for the rest of the month to meet the target.

---

## ⚙️ Features

- Pulls time data from Clockify via API
- Calculates:
    - Total hours worked this month
    - Remaining hours to meet the goal
    - Daily hour target based on weekdays left
- Sends HTML-formatted email reports
- Supports debug mode for verbose output

---

## 🛠️ Requirements

- Go 1.18+
- A valid Clockify API key
- Email SMTP server for sending reports

---

## 📦 Installation

1. Clone the repository:

```bash
git clone https://github.com/roland198412/Clockify-Client-Hour-Tracker.git
cd clockify-client-hour-tracker
```

2. Build the application

```bash
go build -o hour-tracker
```

2. Build the application:

```bash
go build -o hour-tracker
```

---

## 📄 Environment Variables

| Variable             | Description                               |
|----------------------|-------------------------------------------|
| `WORKSPACE_ID`       | Clockify Workspace ID                     |
| `CLOCKIFY_API_KEY`   | Clockify API key                          |
| `CLIENT_NAME`        | Name of the client                        |
| `CLIENT_NID`         | ID of the client                          |
| `MONTHLY_HOUR_LIMIT` | Monthly limit in hours (e.g. `50`)        |
| `SMTP_HOST`          | SMTP host (e.g., `smtp.gmail.com`)        |
| `SMTP_PORT`          | SMTP port (e.g., `587`)                   |
| `SMTP_USERNAME`      | Email address used to send reports        |
| `SMTP_PASSWORD`      | SMTP password or app password             |
| `EMAIL_RECIPIENT`    | Recipient email address                   |
| `DEBUG`              | Set to `true` for debug output (optional) |

---

## 🧪 Usage

After configuring your environment variables, run the app:

```bash
./hour-tracker
```

If `DEBUG=true` is set, detailed output will be printed to the terminal.

---

## 📧 Email Report Sample

An HTML email will be sent containing the following summary:

- Client name
- Reporting period
- Total time logged
- Remaining time available
- Required daily hours to reach the monthly goal

---

## 🧰 Tech Stack

- Go
- Clockify API
- SMTP (email)

---

## 📌 Notes

- Only weekdays (Mon–Fri) are counted when calculating daily targets.
- Designed to run once daily (ideal as a cron job or scheduled task).

---

## 📝 License

MIT License. See [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgements

- [Clockify](https://clockify.me) for their API
- Inspiration from daily productivity tracking needs
