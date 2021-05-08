SRC=main.go new_project.go setting.go image-build.go html-build.go

build:
	go build -o acc.exe $(SRC)


