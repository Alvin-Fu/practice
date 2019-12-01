package builder


type Builder interface{
	Engine()
	Glass()
	SteeringWheel()
}

type Director struct{
	builder Builder
}
func NewDirector(builder Builder)*Director{
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct(){
	d.builder.Engine()
	d.builder.Glass()
	d.builder.SteeringWheel()
}

type Car1 struct {
	car string
}
func NewCar1()*Car1{
	return &Car1{}
}

func (c *Car1) Engine(){
	c.car += "car1 engine "
}
func (c *Car1) Glass(){
	c.car += " glass "
}
func (c *Car1) SteeringWheel(){
	c.car += " steering wheel "
}
func (c *Car1) GetCar()string{
	return c.car
}
type Car2 struct {
	car string
}
func NewCar2()*Car2{
	return &Car2{}
}

func (c *Car2) Engine(){
	c.car += "car2 engine "
}
func (c *Car2) Glass(){
	c.car += " glass "
}
func (c *Car2) SteeringWheel(){
	c.car += " steering wheel "
}
func (c *Car2) GetCar()string{
	return c.car
}
