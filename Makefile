# development
dep :
	go get -u github.com/kardianos/govendor
	govendor sync

save_dep :
	govendor add +external

pretty:
	gofmt -w **/*.go
	gofmt -w *.go