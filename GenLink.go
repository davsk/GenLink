// GenLink.go

// by David Skinner

/* 
For instance, the UPC-A barcode for a box of tissues is "036000241457".
The last digit is the check digit "7", 
and if the other numbers are correct 
then the check digit calculation must produce 7.

Add the odd number digits: 0+6+0+2+1+5 = 14.
Multiply the result by 3: 14 × 3 = 42.
Add the even number digits: 3+0+0+4+4 = 11.
Add the two results together: 42 + 11 = 53.

To calculate the check digit, 
take the remainder of (53 / 10), which is also known as (53 modulo 10), 
and if not 0, subtract from 10. 
Therefore, the check digit value is 7. i.e. (53 / 10) = 5 remainder 3; 10 - 3 = 7.
*/

package main

import (
    "fmt"
    "strconv"
)

func main() {
    const base = 85000302400
    const even = 5 + 0 + 3 + 2
    const odd = 8 + 0 + 0 + 0 + 4
    
    for i := 0; i < 100; i++ {
        od := i % 10;
        ev := (i -od)/10
        sum := ((3 * (odd + od)) + even + ev) % 10 
        if sum != 0 {
            sum = 10 - sum;
        }
        fmt.Println(base + i, sum)        
        fmt.Println()
        fmt.Println("https://barcode.tec-it.com/barcode.ashx?data="+ strconv.Itoa(base + i)+strconv.Itoa(sum)+ "&code=UPCA&multiplebarcodes=false&translate-esc=false&unit=Fit&dpi=96&imagetype=Gif&rotation=0&color=%23000000&bgcolor=%23ffffff&qunit=Mm&quiet=0");
        fmt.Println()
        
        fmt.Println("![](https://barcode.tec-it.com/barcode.ashx?data="+ strconv.Itoa(base + i)+strconv.Itoa(sum)+ "&code=UPCA&multiplebarcodes=false&translate-esc=false&unit=Fit&dpi=96&imagetype=Gif&rotation=0&color=%23000000&bgcolor=%23ffffff&qunit=Mm&quiet=0)");
        fmt.Println();
        fmt.Println("https://qrcode.tec-it.com/API/QRCode?data=https%3a%2f%2fHuny-B.CBD-Oil.App%2fUPC%2f"+strconv.Itoa(i)+"%2f&errorcorrection=M&backcolor=%23ffffff&quietzone=0.125&quietunit=In&size=Small");
        fmt.Println();
        fmt.Println("![](https://qrcode.tec-it.com/API/QRCode?data=https%3a%2f%2fHuny-B.CBD-Oil.App%2fUPC%2f"+strconv.Itoa(i)+"%2f&errorcorrection=M&backcolor=%23ffffff&quietzone=0.125&quietunit=In&size=Small)");
        }
}