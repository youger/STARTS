import Foundation

//Definition for singly-linked list.
public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init() { self.val = 0; self.next = nil; }
    public init(_ val: Int) { self.val = val; self.next = nil; }
    public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
}

public func createLinklist(array: [Int]) -> ListNode? {
    var head: ListNode? = nil
    var pre = head
    for val in array {
        let node = ListNode(val)
        if head == nil {
            head = node
            pre = node
        }
        pre?.next = node
        pre = node
    }
    return head
}

public func printLinklist(_ head: ListNode?) {
    var node = head
    while node != nil {
        print(node!.val, terminator: "->")
        node = node?.next
    }
    print("nil")
}
