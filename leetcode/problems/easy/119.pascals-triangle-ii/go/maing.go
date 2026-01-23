// https://leetcode.com/problems/pascals-triangle-ii/
// @@@ 1d로 풀려다보니 생각보다 쉽지 않았음..
func getRow(rowIndex int) []int {
    row := []int{1}

    for i := 1; i <= rowIndex; i++ {
        row = append(row, 1)

        for j := i - 1; j > 0; j-- {
            row[j] = row[j] + row[j-1]
        }
    }

    return row
}
