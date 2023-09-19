# Makefile for go program
# alias antlr='java -jar $PWD/antlr-4.13.1-complete.jar'
# antlr -Dlanguage=Go -o parser -visitor  ./SOM.g4 

c1: c1.go service.go
	go build *.go

test: c1
	./c1 ../../Examples/Hello/Hello.som
