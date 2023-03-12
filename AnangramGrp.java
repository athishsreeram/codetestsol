import java.util.*;
import java.util.concurrent.CopyOnWriteArrayList;

public class AnangramGrp {

    public static void main(String[] args) {

        String []src = {"ac","ba","ca","ab","a"};
        AnangramGrp g = new AnangramGrp();
        List<List<String>> lst=  g.groupAnagram(src);
        for(List<String> l:lst){
            System.out.println(l);
        }

    }

    public List<List<String>> groupAnagram(String []src){

        CopyOnWriteArrayList<List<String>> lst =
                new CopyOnWriteArrayList<List<String>>();
        List<Integer> completed = new CopyOnWriteArrayList<>();

        if(lst.size() == 0 && src.length>0){
            CopyOnWriteArrayList<String> l =
                    new CopyOnWriteArrayList<String>();
            l.add(src[0]);
            lst.add(l);
            completed.add(0);
        }



            for (int i = 1; i < src.length; i++) {

                for (Iterator<List<String>> it = lst.iterator(); it.hasNext(); ) {
                    List<String> angsrc = it.next();
                    if (isAnangram(angsrc.get(0).toCharArray(), src[i].toCharArray())) {
                        angsrc.add(src[i]);
                        completed.add(i);
                    }
                }
                if(!completed.contains(i)){
                    CopyOnWriteArrayList<String> l =
                            new CopyOnWriteArrayList<String>();
                    l.add(src[i]);
                    lst.add(l);
                    completed.add(i);
                }


            }

        return  lst;

    }


    public boolean isAnangram(char []ang,char []src){

        Map srcA = new HashMap<Character,Integer>();


        // create a map
        for(char a:ang){
            if(!srcA.containsKey(a))
             srcA.put(a,1);
            else {
                Integer o = (Integer) srcA.get(a);
                srcA.put(a, o + 1);
            }
        }

        for(char s:src){
            if(!srcA.containsKey(s)){
                return false;
            }else{
                Integer o = (Integer) srcA.get(s);
                srcA.put(s, o - 1);
            }
        }

        for (Object key : srcA.keySet()) {
            if((Integer)srcA.get(key) != 0)
            {
                return false;
            }
        }
        return true;
    }

}
