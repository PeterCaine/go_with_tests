# Ch2 - integers
## Takeaways
- package name for test and entrypoint - should be identical
- we can add examples to test:
```
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

```

if the commented out expected output is there and the cmd: "go test -v"
is run, the example will also run

if we run "godoc -http=localhost:8080" we will see under the packages name (in this case 'integers') the signature and the example. 
