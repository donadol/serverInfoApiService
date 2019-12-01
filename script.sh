go get -u github.com/lib/pq
go get -u github.com/go-chi/chi
go get github.com/PuerkitoBio/goquery

cockroach start-single-node \
--insecure \
--listen-addr=localhost:26257 \
--http-addr=localhost:8080 \
--background

cockroach init --insecure --host=localhost:26257
cockroach sql --insecure < statements.sql