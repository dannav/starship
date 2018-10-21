get-test-data:
	rm -rf bin
	rm -rf files
	mkdir -p bin
	mkdir -p files
	rm -rf ./bin/*
	go build -o ./bin/aitestdata ./aitestdata.go
	for lang in "javascript" "c" "go" "python" "ruby" ; do export TOKEN=3724162d475b91ae749345e27cf80d24ea7844b0; ./bin/aitestdata -l $$lang -o ./files; done