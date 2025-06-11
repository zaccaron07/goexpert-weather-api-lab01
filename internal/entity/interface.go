package entity

type ZipcodeRepositoryInterface interface {
	Get(string) (Zipcode, error)
}
