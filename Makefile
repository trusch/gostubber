test: test_stage1 test_stage2

test_stage1:
	go run test.go -stage 1

test_stage2:
	go run test.go -stage 2

clean:
	-rm ./stubber/*_stub.go 2>/dev/null