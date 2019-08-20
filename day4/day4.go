package day4

type base interface {

}

//通过typeSwitch 判断类型
func typeSwitch(f base) base {
	switch f.(type) {
	case int:
		return  f.(int)*2;
	case string:
		return f.(string)+f.(string)+f.(string)
	}
	return  f
}

func Map(arr []base,f func( base) base)[]base{
	for k,v :=range arr{
		arr[k] = f(v)
	}

	return  arr
}




