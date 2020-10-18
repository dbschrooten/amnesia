package service

type Implementation interface {
	Info() map[string]string
	Run() error
}
