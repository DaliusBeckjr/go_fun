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



when having mutliple files in our go file we can add a go.work file more info inside the documentation.


**Side Note**: when in the terminal and we have multiple files we can run that file like so 
```sh
functions git:(main) -> go run ./functions/main.go
```
when having multiple files we would run that specific file from the go.work
to add into the work file we would use

```sh
functions git:(main) -> go use (new file)
```