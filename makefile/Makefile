OBJ=$(wildcard *.go)
test:$(OBJ)
	go build -o $@ $^
	
x:=foo
y:=$(x)b	
x?=new
test1:
	@echo "y=>$(y)"
	@echo "x=>$(x)"

OBJ=$(patsubst %.go,%.o,1.go 2.go 3.go)
test2:
	@echo $(OBJ)	
OBJ=$(strip a     b    c)
test3:
	@echo $(OBJ)




