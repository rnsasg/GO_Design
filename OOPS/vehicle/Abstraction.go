package vehicle

import (
	"github.com/rnsasg/GO_Design/OOPS/vehicle"
)

type Vehicle interface {
	GetNoOfWheel() int
}

// func (v *Vehicle) GetNoOfWheel() int {
// 	return v.noOfWheel
// }

type Bicycle struct {
	v vehicle.Vehicle
}

// func main() {
// 	vObj := vehicle.NewVechile(2, vehicle.RED)
// 	b := Bicycle{v: vObj}
// 	fmt.Println(b.v.GetNoOfWheel())
// }
