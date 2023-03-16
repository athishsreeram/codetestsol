import java.util.Stack;

public class RPN {

    public static void main(String[] args) {
       String a = "13+2*";
        RPN rpn =  new RPN();
       System.out.println(rpn.RPN(a));
    }

    public int RPN(String rpnStr){

        char []rpnChr = rpnStr.toCharArray();

        Stack<Character> s = new Stack<>();

        int a;
        int b;
        Integer c;

        for(char r:rpnChr){
            Character c1 = new Character('*');
            Character c2 = new Character('+');
            if(c1.equals(r)){
               a= Character.getNumericValue(s.pop());b=Character.getNumericValue(s.pop());
               c = a * b;
               s.push( c.toString().toCharArray()[0]);
            }else if( c2.equals(r)){
                a= Character.getNumericValue(s.pop());b=Character.getNumericValue(s.pop());
                c = a + b;
                s.push(c.toString().toCharArray()[0]);
            }else{
                s.push((char)r);
            }
        }

        return Character.getNumericValue(s.pop());

    }



}
