// https://leetcode.com/problems/time-based-key-value-store/description/
// timestamp값이 엄격하게 증가하므로 heap이나 다른 자료구조는 안써도 될듯함.
// @@@@@@@
class TimeMap {
    data class Record(val value: String, val timestamp: Int)

    private val elements = mutableMapOf<String, MutableList<Record>>()

    fun set(key: String, value: String, timestamp: Int) {
        elements.getOrPut(key) { mutableListOf() }.add(Record(value, timestamp))
    }

    fun get(key: String, timestamp: Int): String {
        val values = elements[key] ?: return ""

        var left = 0
        var right = values.lastIndex
        var result = ""

        while (left <= right) {
            val mid = (left + right) / 2
            if (values[mid].timestamp <= timestamp) {
                result = values[mid].value
                left = mid + 1
            } else {
                right = mid - 1
            }
        }

        return result
    }
}
