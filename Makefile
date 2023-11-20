test:
	go test -timeout=3s -race -count=10 -failfast -shuffle=on -short ./...
	go test -timeout=10s -race -count=1 -failfast  -shuffle=on ./...

bench:
	go test -bench ^Benchmark* -run XXX -benchtime 1s -count 1 -cpu 1,2 -benchmem | tee benchmark_output.txt
