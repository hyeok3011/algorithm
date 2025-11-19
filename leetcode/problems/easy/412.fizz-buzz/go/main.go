// https://leetcode.com/problems/fizz-buzz/
func fizzBuzz(n int) []string {
    answer := make([]string, n)

    for i := 0; i < n; i++ {
        number := i + 1
        if number % 3 == 0 {
            answer[i] = "Fizz"
        }

        if number % 5 == 0 {
            answer[i] += "Buzz"
        }

        if answer[i] == "" {
            answer[i] = strconv.Itoa(number)
        }
    }

    return answer
}
