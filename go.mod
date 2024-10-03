module github.com/kjbreil/sil

go 1.23

require golang.org/x/text v0.18.0

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/kjbreil/crcloc v0.0.0-20240619200433-6432ff988581 // indirect
	github.com/kjbreil/crlf v0.0.0-20210116185654-a98352303dd9 // indirect
	github.com/kjbreil/glsp v0.2.2 // indirect
	github.com/kjbreil/loc-macro v0.0.0-20240917185356-d96238df39ad
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sourcegraph/jsonrpc2 v0.2.0 // indirect
)

replace (
	github.com/kjbreil/glsp => /Users/kjell/dev/glsp
	github.com/kjbreil/go-smb2 => /Users/kjell/dev/go-smb2
	github.com/kjbreil/go-sqlfmt => /Users/kjell/dev/go-sqlfmt
	github.com/kjbreil/loc-macro => /Users/kjell/dev/loc-macro
	github.com/kjbreil/ziploc => /Users/kjell/dev/ziploc
)
