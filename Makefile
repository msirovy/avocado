build:
	env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-extldflags '-static'" -o avocado