## Try it out

First init the module for your version of golang
```
go mod init <module_name>
```

#### Install the urfave package: 

```
go get github.com/urfave/cli/v2
```

#### build a executable for the app:

```
 go build cli-go.go
```

#### run on windows:
```
./cli-go.exe or cli-go.exe
```

#### for linux:
```
./cli-go
```

don't forget to specify the arguments and the flags

Also possible to add the binary to a folder, add the path of this foleder to your sytem enviroment vars of windows and type the name of the binary file without ".exe", then type start, 
it will run the node apps of the current directory, if they have a main index.js, app.js or server.js file, it could also be adjusted to suport other file names as well if needed.




