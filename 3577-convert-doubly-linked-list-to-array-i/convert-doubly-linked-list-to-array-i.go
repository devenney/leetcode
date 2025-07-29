/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Prev *Node
 * }
 */

func toArray(head *Node) []int {
    var array []int

    current := head
    
    for true {
        array = append(array, current.Val)
        if current.Next == nil {
            break
        }
        current = current.Next
    }

    return array
}