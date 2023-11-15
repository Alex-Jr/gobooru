t="/tmp/go-cover.$$.tmp"
go test -v ./... -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
