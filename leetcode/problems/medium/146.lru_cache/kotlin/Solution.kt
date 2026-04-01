// https://leetcode.com/problems/lru-cache/
class LRUCache(private val capacity: Int) {

    private val nodeMap = HashMap<Int, Node>(capacity) // 다른 사람들은 여기에 doubling을 막기위해 1.0f를 넣는듯
    private val head = Node(-1, -1)
    private val tail = Node(-1, -1)

    private class Node(
        val key: Int,
        var value: Int,
    ) {
        var prev: Node? = null
        var next: Node? = null
    }

    init {
        head.next = tail
        tail.prev = head
    }

    fun get(key: Int): Int {
        val node = nodeMap[key] ?: return -1
        unlink(node)
        linkTail(node)
        return node.value
    }

    fun put(key: Int, value: Int) {
        nodeMap[key]?.let { existing ->
            existing.value = value
            unlink(existing)
            linkTail(existing)
            return
        }

        if (nodeMap.size == capacity) evictLRU()

        val newNode = Node(key, value)
        nodeMap[key] = newNode
        linkTail(newNode)
    }

    private fun unlink(node: Node) {
        node.prev?.next = node.next
        node.next?.prev = node.prev
        node.prev = null
        node.next = null
    }

    private fun linkTail(node: Node) {
        val last = tail.prev ?: return
        node.prev = last
        node.next = tail
        last.next = node
        tail.prev = node
    }

    private fun  evictLRU() {
        val victim = head.next ?: return
        unlink(victim)
        nodeMap.remove(victim.key)
    }
}