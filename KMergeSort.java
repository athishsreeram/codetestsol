import java.util.HashMap;
import java.util.LinkedList;
import java.util.Map;

public class KMergeSort {


    public static void main(String[] args) {


        LinkedList []list = new LinkedList[3];

        list[0]= new LinkedList<Integer>();
        list[0].add(2);
        list[0].add(3);

        list[1]= new LinkedList<Integer>();
        list[1].add(1);
        list[1].add(3);

        list[2]= new LinkedList<Integer>();
        list[2].add(2);
        list[2].add(3);

        LinkedList kMerged = KMergeSorting(list);
        kMerged.forEach(a ->
                System.out.println(a));
    }

    private static LinkedList KMergeSorting(LinkedList[] list) {


        LinkedList out = new LinkedList<Integer>();

        Map mapList = new HashMap<Integer,LinkedList>();

        for (int i=0;i<list.length;i++){
            mapList.put(i,list[i]);
        }

        Integer[] minValue = {null,Integer.MAX_VALUE};

        Integer currentNode = Integer.MAX_VALUE;
        while(currentNode != null) {
            mapList.forEach((k, v) -> {
                Integer listIndex = (Integer) k;
                LinkedList listLink = (LinkedList) v;
                /*if(listLink !=null && listLink.size() !=0)
                    System.out.println(listIndex + " "+ listLink.getFirst() );*/
                if(listLink !=null && listLink.size() !=0  && minValue[1] > (int)listLink.getFirst()){
                    minValue[0] = listIndex;
                    minValue[1] = (int)listLink.getFirst();
                }
            });

            Integer currentKey = minValue[0];
            minValue[1] = Integer.MAX_VALUE;

            LinkedList minKeyList = (LinkedList) mapList.get(currentKey);
            currentNode = (Integer) minKeyList.poll();
            out.add(currentNode);
            //System.out.println(currentNode);

        }



        return out;
    }


}
