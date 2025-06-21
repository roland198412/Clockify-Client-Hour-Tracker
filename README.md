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
git clone https://github.com/roland198412/clockify-client-hour-tracker.git
cd clockify-client-hour-tracker
