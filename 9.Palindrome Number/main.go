func isPalindrome(x int) bool {
    var isPalindromeString func(string) bool
    isPalindromeString = func(str string) bool {
        if len(str) < 2 {return true}
        
        l, r := 0, len(str) - 1
        for l < r {
            if str[l] != str[r] {
                return false
            }
            l++
            r--
        }
        return true
    }
    
    str := strconv.Itoa(x)
    return isPalindromeString(str)
}
