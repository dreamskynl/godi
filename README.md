# GoDI - Dependency Injection Container
GoDI is a light weight yet easy-to-use IoC/DI container. It may replace the `uber/dig` and `samber/do` fantastic package in simple Go projects.

# Features
- Service registration
- Service invocation
- Service shutdown

# Install
```go
go get -u github.com/dreamskynl/godi@latest
```

# Usage
```go
// use container

c := godi.New()

if err := c.Register(&CatService{}, NewCatService, "Meow hi", "Meow bye~"); err != nil {
    panic(err)
}

catService := c.MustResolveAsInstance(&CatService{}).(CatService)

catService.SayHi()
```
```go
// service

type ICatService interface {
	SayHi()
	SayBye()
}

type CatService struct {
	hi  string
	bye string
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
# In real MVC project
### Use GoDI for prettier code
```go
/*
  Init GoDI Container 
*/
SERVICE_CONTAINER = godi.New()

if err := SERVICE_CONTAINER.Register(&services.AdminUserService{}, services.NewAdminUserService, databases.DB_CONN); err != nil {
	panic(err)
}

if err := SERVICE_CONTAINER.Register(&services.EmailService{}, services.NewEmailService, utils.EmailClient); err != nil {
	panic(err)
}

if err := SERVICE_CONTAINER.Register(&services.JWTService{}, services.NewJWTService); err != nil {
	panic(err)
}

/*
  Use controller 
*/
func Login(ctx *gin.Context) {
	adminUserController := http_controller.NewAdminUserController(ctx, godicontainer.SERVICE_CONTAINER)
	adminUserController.Login()
}

func AddAdmin(ctx *gin.Context) {
	adminUserController := http_controller.NewAdminUserController(ctx, godicontainer.SERVICE_CONTAINER)
	adminUserController.CreateUser()
}

func GetAdmins(ctx *gin.Context) {
	adminUserController := http_controller.NewAdminUserController(ctx, godicontainer.SERVICE_CONTAINER)
	adminUserController.GetUsers()
}

func DeleteAdmin(ctx *gin.Context) {
	adminUserController := http_controller.NewAdminUserController(ctx, godicontainer.SERVICE_CONTAINER)
	adminUserController.DeleteUser()
}

func ResetPassword(ctx *gin.Context) {
	adminUserController := http_controller.NewAdminUserController(ctx, godicontainer.SERVICE_CONTAINER)
	adminUserController.ResetPassword()
}

func ChangePassword(ctx *gin.Context) {
	adminUserController := http_controller.NewAdminUserController(ctx, godicontainer.SERVICE_CONTAINER)
	adminUserController.ChangePassword()
}

/*
  Define controller
*/
type IAdminUserController interface {
	Login()
	GetUsers()
	CreateUser()
	DeleteUser()
	ResetPassword()
	ChangePassword()
}

type AdminUserController struct {
	ctx          *gin.Context
	UserService  services.IAdminUserService
	JWTService   services.IJWTService
	emailService services.IEmailService
}

func NewAdminUserController(ctx *gin.Context, serviceContainer godi.IGoDI) IAdminUserController {
	return &AdminUserController{
		ctx:          ctx,
		UserService:  serviceContainer.MustResolve(&services.AdminUserService{}).(*services.AdminUserService),
		JWTService:   serviceContainer.MustResolve(&services.JWTService{}).(*services.JWTService),
		emailService: serviceContainer.MustResolve(&services.EmailService{}).(*services.EmailService),
	}
}
```
### Instead of
```go
/*
  Use controller
*/
func Login(ctx *gin.Context) {
	adminUserService := services.NewAdminUserService(databases.DB_CONN)
	jwtService := services.NewJWTService()
	adminUserController := http_controller.NewAdminUserController(ctx, adminUserService, &jwtService, nil)

	adminUserController.Login()
}

func AddAdmin(ctx *gin.Context) {
	adminUserService := services.NewAdminUserService(databases.DB_CONN)
	adminUserController := http_controller.NewAdminUserController(ctx, adminUserService, nil, nil)

	adminUserController.CreateUser()
}

func GetAdmins(ctx *gin.Context) {
	adminUserService := services.NewAdminUserService(databases.DB_CONN)
	adminUserController := http_controller.NewAdminUserController(ctx, adminUserService, nil, nil)

	adminUserController.GetUsers()
}

func DeleteAdmin(ctx *gin.Context) {
	adminUserService := services.NewAdminUserService(databases.DB_CONN)
	adminUserController := http_controller.NewAdminUserController(ctx, adminUserService, nil, nil)

	adminUserController.DeleteUser()
}

func ResetPassword(ctx *gin.Context) {
	adminUserService := services.NewAdminUserService(databases.DB_CONN)
	emailService := services.NewEmailService(utils.EmailClient)
	adminUserController := http_controller.NewAdminUserController(ctx, adminUserService, nil, &emailService)

	adminUserController.ResetPassword()
}

func ChangePassword(ctx *gin.Context) {
	adminUserService := services.NewAdminUserService(databases.DB_CONN)
	jwtService := services.NewJWTService()
	adminUserController := http_controller.NewAdminUserController(ctx, adminUserService, &jwtService, nil)

	adminUserController.ChangePassword()
}

/*
  Define controller
*/
type IAdminUserController interface {
	Login()
	GetUsers()
	CreateUser()
	DeleteUser()
	ResetPassword()
	ChangePassword()
}

type AdminUserController struct {
	ctx          *gin.Context
	UserService  services.IAdminUserService
	JWTService   *services.IJWTService
	emailService *services.IEmailService
}

func NewAdminUserController(ctx *gin.Context, userService services.IAdminUserService, jwtService *services.IJWTService, emailService *services.IEmailService) IAdminUserController {
	return &AdminUserController{
		ctx:          ctx,
		UserService:  userService,
		JWTService:   jwtService,
		emailService: emailService,
	}
}
```
