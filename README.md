# Transmo
Transmo is a Go library for transform model into another model base on struct. This library detect your field name to copy that into another model. Basically this library just use reflection to copy model and read field and tag.

## Installation
 
To install Transmo package, you need to install Go and set your Go workspace first.

1. You can use the below Go command to install Transmo.

```sh
$ go get -u github.com/rezaramadhanirianto/transmo
```

## How to use it

````go
type Model struct {
	Id          string `transmo:"ignore"` // this will ignore transmo to copy this field
	Title       string
	Description string
}

type FakeModel struct {
	Id          string
	Title       string
	Description string
}


model := Model{
 Id: "Id",
 Title: "Title",
}
fakeModel := FakeModel{}

transmo.Transform(model, &fakeModel)

fmt.Println(fakeModel) // { Title }
````
