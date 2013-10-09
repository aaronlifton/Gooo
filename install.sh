#!/bin/bash
go get github.com/bmizerany/pq
go get github.com/aaronlifton/gooo/introspection
go get github.com/aaronlifton/gooo/memory
go get github.com/aaronlifton/gooo/model
go get github.com/aaronlifton/gooo/router
go get github.com/aaronlifton/gooo/session
go get github.com/aaronlifton/gooo/util
go get github.com/aaronlifton/gooo/view
go get github.com/aaronlifton/gooo
go build
relocate () {
  cd $GOPATH/src/github.com/aaronlifton
}
relocate
go install