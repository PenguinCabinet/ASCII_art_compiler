SRC=main.go new_project.go setting.go image-build.go html-build.go pdf-build.go real-preview.go

build:
	go build -o acc.exe $(SRC)
test:
	mkdir ./test/test1
	cd ./test/test1 && go test ../..


