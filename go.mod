module github.com/kitwork/kitwork

go 1.25.0

require github.com/kitwork/engine v0.1.1

require (
	github.com/lib/pq v1.10.9 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/kitwork/engine => ./engine
