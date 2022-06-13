package matematika

func CekGanjilGenap(data ...int) string {
	var hasil string
	length := len(data) - 1
	for i, data := range data {
		if data%2 == 0 {
			if length == i {
				hasil += "Genap"
			} else {
				hasil += "Genap, "
			}
		} else {
			if length == i {
				hasil += "Ganjil"
			} else {
				hasil += "Ganjil, "
			}
		}
	}
	return hasil
}
