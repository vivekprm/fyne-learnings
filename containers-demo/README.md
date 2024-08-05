# Containers
If we want to put complex application together, you're probably going to look for various types of containers that help to position them on screen with potentially some interaction. 
E.g. Here we have horizontal split.
```go
container.NewHSplit(left,right)
```

A container can have any layout that we want, we can say:
```go
container.New(layout, object...)
```
We can have anything that implements Layer interface, which simply says how big something be at its minimum and what do you do to lay things out internally.