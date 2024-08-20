# GoDI - Dependency Injection Container
GoDI is a light weight yet easy-to-use IoC/DI container. It may replace the `uber/dig` and `samber/do` fantastic package in simple Go projects.

# Features
Service registration
Service invocation
Service shutdown

# Install
```go
go get -u github.com/DreamSkyLL/godi@v1
```

# Usage
```go
	c := godi.New()

	if err := c.Register(&CatService{}, NewCatService, "NewCatService says hi", "NewCatService says bye"); err != nil {
		t.Error(err)
	}

    c.Say
```
```go
type ICatService interface {
	SayHi()
	SayBye()
}

type CatService struct {
	hi string
	bye  string
}

func NewCatService(hi string, bye string) (ICatService, error) {
	return &CatService{hi: hi, bye: bye}, nil
}

func (b *CatService) SayHi() {
	fmt.Println(b.hi)
}

func (b *CatService) SayBye() {
	fmt.Println(b.bye)
}
```