#!/bin/bash
export HEREIS=$PWD
alias antlr='java -jar $PWD/antlr-4.13.1-complete.jar'
# antlr -Dlanguage=Go -o parser -visitor  ./SOM.g4

alias som='java -jar $PWD/som.jar'
alias somc='java -cp $PWD/som.jar som.compiler.Main'
alias somi='java -cp $PWD/som.jar som.vm.Universe'
alias somd='java -cp $PWD/som.jar som.interpreter.Main'
alias somt='java -cp $PWD/som.jar som.tests.AllSuites'
 
# /usr/bin/env java  -cp som/bin vm.Universe -cp Smalltalk -d Hello.som >& Hello.txt