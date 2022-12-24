
clean:
	rm *prof || true
	rm *.test || true
	rm main || true
	rm *.main || ture

gen-old:
	go test -cpuprofile old.cpu.prof -memprofile old.mem.prof -bench=. -benchmem old/*.go
	go build -gcflags=-m -o old.main old/*.go

gen-new:
	go test -cpuprofile new.cpu.prof -memprofile new.mem.prof -bench=. -benchmem new/*.go
	go build -gcflags=-m -o new.main new/*.go

prof-old: gen-old
	go tool pprof -http=:8000 old.main old.cpu.prof

prof-new: gen-new
	go tool pprof -http=:8001 new.main new.cpu.prof