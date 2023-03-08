import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class ContainsDuplicate {

    public static void main(String[] args) {
        ContainsDuplicate cd = new ContainsDuplicate();

        int []a = {1,2,3,10};
        System.out.println(cd.containDuplicate(a));
    }

    public boolean containDuplicate(int[] args)
    {
       /* Map<Integer, Integer> map = Arrays.stream(args)
                .boxed()
                .collect(Collectors.toMap(
                        Function.identity(),
                        i -> args[i]
                ));**/
        Map<Integer, Integer> map = new HashMap<>();
        boolean chk = false;

        for(int i = 0; i<args.length; i++){
            if(map.containsKey(args[i])) {
                chk = true;
            }
          map.put(args[i],i);
      }


        return chk;
    }


}
