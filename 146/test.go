package main

import "fmt"

func testExample()  {
	cache := Constructor( 2 /* 缓存容量 */ );
	cache.Put(1, 1);
	cache.Put(2, 2);
	fmt.Println(cache.Get(1));       // 返回  1
	cache.Put(3, 3);    // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2));       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1));       // 返回 -1 (未找到)
	fmt.Println(cache.Get(3));       // 返回  3
	fmt.Println(cache.Get(4));       // 返回  4
}

func testError1()  {
	cache := Constructor( 1 /* 缓存容量 */ );
	cache.Put(2, 1);
	fmt.Println(cache.Get(2));       // 返回  1
	cache.Put(3, 2);    // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2));       // 返回 -1 (未找到)
	fmt.Println(cache.Get(3));       // 返回  2
}

func testError2()  {
	cache := Constructor( 2 /* 缓存容量 */ );
	fmt.Println(cache.Get(2));       // 返回  -1
	cache.Put(2, 6);
	fmt.Println(cache.Get(1));       // 返回  1
	cache.Put(1, 5);
	cache.Put(1, 2);
	fmt.Println(cache.Get(1));       // 返回 2
	fmt.Println(cache.Get(2));       // 返回 -1
}

func testError3()  {
	cache := Constructor( 2 /* 缓存容量 */ );
	cache.Put(2, 1);
	cache.Put(1, 1);
	cache.Put(2, 3);
	cache.Put(4, 1);
	fmt.Println(cache.Get(1));       // 返回 -2
	fmt.Println(cache.Get(2));       // 返回 3
}

func main()  {
	//testExample()
	//testError1()
	//testError2()
	testError3()
}
