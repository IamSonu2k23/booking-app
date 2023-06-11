# booking-app
A Booking Service which allows you to choose seats and book

# Requirements
1. Postgres
2. Golang

# Running Steps on Windows
1. Install Postgres database if not installed
2. git clone https://github.com/IamSonu2k23/booking-app.git
3. cd booking-app
4. go get
5. verify env parameters set in ".env" file in it
6. go run main.go
7. Access the app e.g localhost:9080 or at set port number

# Running Steps on Docker
1. From Dockerfile make docker image
2. First run postgres container 
3. Then run your booking-app container
4. Make sure environment parameters are set
5. Access the app e.g localhost:9080 or hostaddress:9080 (or at set port number)
