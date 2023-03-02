public class ReverseList {

    public static void main(String[] args) {
        Node node = new Node(1);
        node.next = new Node (2);
        node.next.next = new Node (3);

        printNode(node);

        printNode(reverseList(node));


    }

    private static Node reverseList(Node head) {

        Node prev = null;
        Node cur = head;

        while(cur != null){
            Node nextNode = cur.next;
            cur.next = prev;
            prev = cur;
            cur = nextNode;
        }

        return prev;
    }

    private static void printNode(Node cur) {
        while(cur !=null){
            System.out.println(cur.v);
            cur = cur.next;
        }
    }


}
