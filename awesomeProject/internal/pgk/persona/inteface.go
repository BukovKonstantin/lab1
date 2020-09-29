package persona

type Iusecase interface {
	Get(p *Person) (uint, int)
	Read(p *Person) (int, uint)
	ReadAll() ([]*persona., int)
	Delete(id int) int
	Update(id uint, p *PersonRequest) int
}

type Irep interface {
	Get(p *Person) (int, int)
	Read(p *Person) (int, int)
	ReadAlls(p *Person) (int, int)
	Delete(id int) int
	Update(id uint, p *Person) int
}
