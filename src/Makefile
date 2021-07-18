TEST_FILES = $(shell find -name '*_test.go')
export TESTMODE=test
build:
	go version
test:
	for file in $(TEST_FILES); do\
	    echo $$file;\
		go test ./$$file;\
	done
	
	
