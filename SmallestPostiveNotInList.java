 import java.util.*;
 import java.util.stream.Collectors;
 import java.util.stream.IntStream;


/**

Write a function:

class SmallestPostiveNotInList { public int solution(int[] A); }

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].

**/

class SmallestPostiveNotInList {
    public int solution(int[] A) {
        IntStream intStream = Arrays.stream(A);
        OptionalInt max = intStream.max();

        List<Integer> intList = new ArrayList<Integer>();
        for (int i : A)
        {
            intList.add(i);
        }



        List<Integer> listWithoutDuplicates =  intList.stream().distinct().sorted().collect(Collectors.toList());

            for(int j=0;j<=listWithoutDuplicates.size()-1;j++){
             Integer element = listWithoutDuplicates.get(j);
                
                if(j+1 != element.intValue()){
                    
                    return j+1;
                }
            
        }



    return max.getAsInt()+1;
    }
}

