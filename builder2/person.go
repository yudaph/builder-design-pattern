package builder2

type person struct {
	name string
	age  int

	address    string
	city       string
	postalCode string

	collegeName string
	major       string
}

func newPerson(name string, age int) *person {
	return &person{name: name, age: age}
}

type PersonBuilder struct {
	person *person
}

func NewPersonBuilder(name string, age int) *PersonBuilder {
	return &PersonBuilder{person: newPerson(name, age)}
}

func (c *PersonBuilder) Address() *AddressBuilder {
	return &AddressBuilder{*c}
}

func (c *PersonBuilder) Collage() *CollegeBuilder {
	return &CollegeBuilder{*c}
}

func (c *PersonBuilder) Build() *person {
	return c.person
}

type AddressBuilder struct {
	PersonBuilder
}

func (c *AddressBuilder) At(address string) *AddressBuilder {
	c.person.address = address
	return c
}

func (c *AddressBuilder) City(city string) *AddressBuilder {
	c.person.city = city
	return c
}

func (c *AddressBuilder) PostalCode(postalCode string) *AddressBuilder {
	c.person.postalCode = postalCode
	return c
}

type CollegeBuilder struct {
	PersonBuilder
}

func (c *CollegeBuilder) At(collegeName string) *CollegeBuilder {
	c.person.collegeName = collegeName
	return c
}

func (c *CollegeBuilder) Major(major string) *CollegeBuilder {
	c.person.major = major
	return c
}
