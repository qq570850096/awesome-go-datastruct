package Sorts

func Select(data []int)  {
	llen := len(data)
	for i:=0;i<llen;i++ {
		tmp := data[i]
		flag := i
		for j:=i+1;j<llen;j++ {
			if data[j] < tmp {
				tmp = data[j]
				flag = j
			}
		}
		if flag != i {
			data[flag] = data[i]
			data[i] = tmp
		}
	}
}
