package prime

// This uses the best simple algorithm as shown in http://en.wikipedia.org/wiki/Primality_test
func Is(p int) bool {
    if p <= 3 {
        return p >= 2
    }
    if p % 2 == 0 || p % 3 == 0 {
        return false
    }
    for i := 5; i*i <= p; i += 6 {
        if p % i == 0 || p % (i + 2) == 0 {
            return false
        }
    }
    return true
}