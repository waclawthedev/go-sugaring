# Deprecation notice
This project transformed to [https://github.com/waclawthedev/gopie](https://github.com/waclawthedev/gopie). Please update your projects to that new package. It has even more useful functions :)
# What go-sugaring is?
This is the package with functions from Python/other language, that you are searching for in Go (Golang)

## Disclaimer
Most of the functions from this lib became possible to be written because of generics. So, it is make sense to use the library only if your project is ready to use go >=1.18

## Installation and usage
1. go get github.com/waclawthedev/go-sugaring
2. Call the function. For example:
```
print(sugar.Contains([]string{"a", "b", "c"}, "a"))
//output: true
```

## Already implemented functions:
1. sugar.Contains - check if slice contains element
2. sugar.IndexesOf - get the slice with occurrences (indexes) of element in slice
3. sugar.Map - Map function. Python analog
4. sugar.Filter - Filter function. Python analog
5. sugar.Reduce - Reduce function. Python analog
6. sugar.Flatten - Transforms matrix to slice
7. sugar.Min - min element in slice
8. sugar.Max - max element in slice
9. sugar.Keys - get keys of map as slice
10. sugar.Values - get values of map as slice
