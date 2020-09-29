package Repository

const (
	ReadPerson = "SELECT name, age, address, workplace FROM public.persona WHERE id = $1;"
	ReadAllPersons = "SELECT id, name, age, address, workplace FROM public.persona;"
	UpdatePerson  = "UPDATE public.persona SET name=$1, age=$2, address=$3, workplace=$4 WHERE id = $5;"
	DeletePerson  = "DELETE FROM public.persona WHERE id = $1;"
	CreatePerson  = "INSERT INTO public.persona(name, age, address, workplace) VALUES ($1, $2, $3, $4) RETURNING id;"
)
