
all:
	go test -coverprofile=cov.out mytesting
	go tool cover -func=cov.out
	go tool cover -html=cov.out