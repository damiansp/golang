package tempconv

// CToF converts temperature Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts temperature Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts temperature C to K
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToK converts temperature F to K
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

// KToC converts temperature K to C
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts temperature K to F
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
