// https://leetcode.com/problems/flood-fill/description/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    originColor := image[sr][sc]

    if originColor == color {
        return image
    }

    var recursion func(row, col int)
    recursion = func(row, col int) {
        if row < 0 || col < 0 || row >= len(image) || col >= len(image[0]){
	    return	
	}

    if image[row][col] != originColor {
        return
    }
	
	image[row][col] = color
	recursion(row + 1, col)
	recursion(row - 1, col)
	recursion(row, col + 1)
	recursion(row, col - 1)
    }

    recursion(sr, sc)	
    return image
}
