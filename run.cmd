@ECHO OFF

SET GOTPATH=~/projetos/cadSolidario/fontes/cadSolidario/
SET PATH=%PATH%,$GOPATH

SET PORT=5000
SET DATABASE_URL=postgres://ulkhcekdsivynv:c17c6b25a4a6fd4db3100e8863e50fefe851464e34e49105ad7860e069e4a272@ec2-52-0-114-209.compute-1.amazonaws.com:5432/d62d5f9ijcofoh

go build -o bin/cadSolidario -v .

rem heroku local web
go run main.go