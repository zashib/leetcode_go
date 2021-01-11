package leetcode_go

type ParkingSystem struct {
	max      []int
	occupied []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		max:      []int{big, medium, small},
		occupied: make([]int, 3),
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.occupied[carType-1] >= this.max[carType-1] {
		return false
	} else {
		this.occupied[carType-1]++
		return true
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
