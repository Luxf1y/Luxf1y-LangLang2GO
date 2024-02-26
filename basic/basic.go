package main

func main() {

	//Part1
	/*	var attack = 40
		var defence = 20
		var damageRate float32 = 0.17
		var damage = float32(attack-defence) * damageRate
		fmt.Println(damage)//damage 变量的右值是一个复杂的表达式，整个过程既有 attack 和 defence 的运算还有强制类型转换。强制类型转换会在后面的章节中介绍。
	*/

	//Part2
	/*	// 声明 hp 变量
		var hp int
		// 再次声明并赋值
		hp := 10*/
	/*	编译报错如下：
		no new variables on left side of :=*/

	//Part3
	/*	conn1, err1 := net.Dial("tcp","127.0.0.1:8080")//此为短变量声明，短变量声明的形式在开发中的例子较多

		var conn net.Conn//此为正常声明
		var err error
		conn, err = net.Dial("tcp", "127.0.0.1:8080")

		fmt.Println(conn)//golang中Print和Println的小区基本一致，但是后者在输出后会进行换行
		fmt.Println(err)

		fmt.Println(conn1)
		fmt.Println(err1)*/

	//Part4
	/*	var a int = 100//使用 Go 的“多重赋值”特性，可以轻松完成变量交换的任务
		var b int = 200
		b, a = a, b
		fmt.Println(a, b)
	*/

}
