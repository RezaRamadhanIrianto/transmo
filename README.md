# Transmo


## Installation

To install Transmo package, you need to install Go and set your Go workspace first.

1. You can use the below Go command to install Transmo.

```sh
$ go get -u github.com/RezaRamadhanIrianto/transmo
```

## How to use it

````go
model := Model{
 Title: "Title",
}
fakeModel := FakeModel{}

transmo.Transform(model, &fakeModel)

fmt.Println(fakeModel) // { Title }
````
