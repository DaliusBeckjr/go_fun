## set up

download go 
cd into the folder 
create a .go file
in the terminal of the directory 
```sh
(example) => go mod init (folder name)
go: creating new go.mod: module functions
go: to add module requirements and sums:
        go mod tidy
```

## to run a file
Open a terminal and change the current directory to the directory which contains the above source file, then run
```sh
go run .
```
or...

```sh
go run main.go
```

**Note**: 
* The <code> go run </code> command isn't a recommended way when running large projects in Go. It's used out of conveniency to run simple small Go programs