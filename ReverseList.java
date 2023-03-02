public class ReverseList {
    public static void main(String[] args) {
        Node n = new Node(1);
        n.next = new Node(2);

        printList(n);

        Node revlist = reverseList(n);
        printList(revlist);

    }

    public static Node reverseList(Node head) {
        Node prev = null;
        Node cur = head;
        while (cur != null) {
            Node nextNode = cur.next;
            cur.next = prev;
            prev = cur;
            cur = nextNode;
        }

        return prev;
    }

    public static void printList(Node cur) {
        while (cur.next != null) {
            System.out.println(cur.v);
            cur = cur.next;
        }
        System.out.println(cur.v);
    }

}
