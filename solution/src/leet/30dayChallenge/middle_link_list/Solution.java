class Solution {
    public ListNode middleNode(ListNode head) {
         ListNode slow = head;
        ListNode fast = head;
        ListNode midNode = head;
        while (fast != null) {

            if(fast.next == null){
                midNode = slow;
                break;
            }else if(fast.next.next == null){
                midNode = slow.next;
                break;
            }else{
                fast = fast.next.next;
                slow = slow.next;
            }
        }
        return midNode;
    }
}