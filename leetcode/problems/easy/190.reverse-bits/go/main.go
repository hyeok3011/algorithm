package main
// https://leetcode.com/problems/reverse-bits/description/
// 최적화 기법은 캐시를 적용하면 될듯함.
func reverseByte(b uint32) uint32 {
    var result uint32
    for i := 0; i < 8; i++ {
        result = (result << 1) | (b & 1)
        b >>= 1
    }
    return result
}
