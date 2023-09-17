
<p>
when multiple arguments are the same type and are next to each other in the function signature. the variable type only needs to be declared after the last argument
</p>


```go
func addToDatabase(hp, power, mana int) {
// ...
}
```

```go
func addToDatabase(hp, power, mana int, name string) {
  // ...
}
```

```go
func addToDatabase(hp, power, mana int, name string, level int) {
  // ...
}
```

key reminder: if they aren't in order for some reason you rebel. they have to be defined seperately

<p> 
as you can see in the examples there are some relations to the C language from the authors becasue C is reverse <mark> addToDatabase(string hp, string power) </mark>
</p>
