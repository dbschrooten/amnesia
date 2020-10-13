package service

type ServiceImpl interface {
	Check() error
	Connect() error
}
