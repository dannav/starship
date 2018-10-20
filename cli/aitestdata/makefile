get-test-data:
	mkdir -p bin
	rm -rf ./bin/*
	go build -o ./bin/aitestdata ./aitestdata
	for lang in "javascript" "c" "go" "python" "ruby" ; do ./bin/aitestdata -l $$lang -o ./files; done