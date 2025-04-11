package main

import (
   "fmt"
   "math"
)

func main() {
	
    maxInt8 := math.MaxInt8
    minInt8 := math.MinInt8
    maxInt16 := math.MaxInt16
    minInt16 := math.MinInt16
    maxInt32 := math.MaxInt32
    minInt32 := math.MinInt32
        
	maxUint8 := math.MaxUint8
    maxUint16 := math.MaxUint16
    maxUint32 := math.MaxUint32
    maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64
	
	fmt.Println("Range of Int8 :: ", minInt8, " to ", maxInt8)
    fmt.Println("Range of Int16 :: ", minInt16, " to ", maxInt16)
    fmt.Println("Range of Int32 :: ", minInt32, " to ", maxInt32)

    fmt.Println("Max Uint8: ", maxUint8) 
    fmt.Println("Max Uint16: ", maxUint16)
    fmt.Println("Max Uint32: ", maxUint32) 
	
    fmt.Println("Max Float32: ", maxFloat32)
	fmt.Println("Max Float64: ", maxFloat64)
	
	fmt.Printf("maxUint8 - Type of data: %T\n", maxUint8)
	fmt.Printf("maxUint16 - Type of data: %T\n", maxUint16) 
	fmt.Printf("maxUint32 - Type of data: %T\n", maxUint32) 
	
	fmt.Printf("maxFloat32 - Type of data: %T\n", maxFloat32) 
	fmt.Printf("maxFloat64 - Type of data: %T\n", maxFloat64)	
}
	
