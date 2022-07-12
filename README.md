# Go-Cron-Job-Demo
This is a Demo project to implement Cron Job in Go


- go mod init github.com/gklathiya/Go-Cron-Job-Demo
- go mod tidy
- go get github.com/robfig/cron/v3@v3.0.0
- go get github.com/sirupsen/logrus


# ┌───────────── minute (0 - 59)
# │ ┌───────────── hour (0 - 23)
# │ │ ┌───────────── day of the month (1 - 31)
# │ │ │ ┌───────────── month (1 - 12)
# │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
# │ │ │ │ │                                   7 is also Sunday on some systems)
# │ │ │ │ │
# │ │ │ │ │
# * * * * * <command to execute>