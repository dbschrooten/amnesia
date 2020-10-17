package service

type ServiceImpl interface {
	Info() map[string]string
	Check() error
	Connect() error
}
