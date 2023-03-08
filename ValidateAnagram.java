import java.util.HashMap;
import java.util.Map;

public class ValidateAnagram {

    public static void main(String[] args) {


        String s = "anagram";
        String t = "nagaram";

        System.out.println(checkAnagram(s,t));
    }

    public static boolean checkAnagram(String a,String b){

        char []ca=a.toCharArray();
        char []cb=b.toCharArray();

        Map<Character, Integer> map = new HashMap<>();
        boolean chk = false;

        for(int i = 0; i<ca.length; i++){
            if(map.containsKey(ca[i])){
                map.put(ca[i], map.get(ca[i])+1);
            }else{
                map.put(ca[i], 1);
            }
        }

        for(int i = 0; i<cb.length; i++){
            if(map.containsKey(cb[i]) && map.get(cb[i])>0)
                map.put(cb[i],map.get(cb[i])-1);
        }

        int[] count = {0};
       map.forEach((k,v)->{
           count[0] = count[0] + v;
       });

if(count[0]>0){
    return chk;
}else{
    return  !chk;
}

    }
}
