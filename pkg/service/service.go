package service

type Authorization interface{

}

type Employee interface{

}

type Service struct{
	Authorization
	Employee
}

func NewService() *Service{
	return &Service{}
}