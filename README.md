# CGO Example with C++

You can build it using linux for linux or cross-compiled for windows.

```
make linux
make windows
```

And to run it on linux is logs a bunch of stuff from the C side which you can pipe away to look at the memory stats only.

```
./test 2> /dev/null
```
