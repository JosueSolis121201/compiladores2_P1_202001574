@echo off

REM Run ANTLR4 Parser Generator
java -classpath %ANTLR4_HOME%\lib\antlr-4.13.0-complete.jar;. org.antlr.v4.Tool -Xlog %*