package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32  
 
	for i := 0; i < 15; i++ {
	   sum += input[i]
	}
 
	result = sum/float32(15)

	return result
}
