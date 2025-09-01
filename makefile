args = cmd/tcpListener/main.go
out = output

build: 
	go build -o $(out) $(args)  

run: 
	go build  -o $(out) $(args) 
	./$(out)

clean: 
	rm $(out) 

debug: 
	dlv debug $(args) 

