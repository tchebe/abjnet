module github.com/tchebe/abjnet/user_cli

go 1.14

require (
	github.com/lib/pq v1.5.2 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/miekg/dns v1.1.29 // indirect
	github.com/tchebe/abjnet/user_service v0.0.0-20201215103336-d071405eb2cc
)

replace github.com/zjjt/abjnet/user_service => github.com/tchebe/abjnet/user_service v0.0.0-20200508133603-c1790a700d4e
