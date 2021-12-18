build:
	go build -o bin/gopher-cli.exe main.go 

run: 
	make build
	bin/gopher-cli.exe $(arg)

clean:
	del *.png
	del bin\*.exe
