package service

type RegisterCardParams struct {
	ID         string
	Vendor     string
	Language   string
	Finish     string
	Stock      int64
	Condition  string
	Visibility string
}
