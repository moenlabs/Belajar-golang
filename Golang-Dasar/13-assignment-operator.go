/* Assignment Operator adalah Operator penugasan yang digunakan untuk memberikan nilai pada variabel.
Pada contoh di bawah, kita menggunakan operator penugasan (=) untuk menetapkan nilai 10 ke variabel bernama x :
*/

package main

import "fmt"

func main() {
	var x = 10
	x += 5 // 15
	a := 10
	a -= 1 // 9
	b := 20
	b *= 2 // 40
	c := 30
	c /= 2 // 15
	d := 40
	d %= 3 // 1
	e := 50
	e &= 3 // 2
	f := 60
	f |= 7 // 63
	g := 70
	g ^= 2 // 68
	h := 80
	h >>= 7 // 0
	i := 90
	i <<= 9 // 46080
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

}

/* Daftar Operator Assignment
=
+=
-=
*=
/=
%=
&=
|=
^=
>>=
<<=
*/
