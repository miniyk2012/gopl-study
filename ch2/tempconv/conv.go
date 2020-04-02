// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC 把绝对零度转换为摄氏度
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

//CToK 把摄氏度转换为绝对零度
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }


//!-
