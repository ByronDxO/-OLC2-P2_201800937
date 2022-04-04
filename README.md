# -OLC2-P2_201800937
LABORATORIO ORGANIZACION DE LENGUAJES Y COMPILADORES 2 Sección D


```bash
go mod init OLC2
go get github.com/gofiber/fiber/v2
``` 


----------------------

# How to Install Go

```bash
sudo apt update
sudo apt upgrade
```

Search for Go:

```bash
sudo apt search golang-go
sudo apt search gccgo-go
```
Install Golang version 1.13 on Ubuntu Linux 20.04 LTS:
```bash
sudo apt install golang-go 

sudo snap install goland --classic

```
----------------------

# How to Install ANTLR

```bash
sudo apt update
sudo apt install antlr4
```
Comando para ejecutar ANTLR
```bash
antlr4 -Dlanguage=Go -o parser ChemsLexer.g4
antlr4 -Dlanguage=Go -o parser Chems.g4

```


----------------------

# Fiber

### Crear Proyecto
```bash
$ go mod init [nombre]
--------------------
-     Ejemplo      - 
- go mod init OLC2 -
--------------------
```

### Dependecias a instalar:
```bash
$ go get github.com/gofiber/fiber
$ go get -u github.com/gofiber/fiber/v2
$ go get -u github.com/gofiber/template
```

### Ejecutar Proyecto
```bash
$ go run app.go
```


