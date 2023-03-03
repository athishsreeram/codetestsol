public class MaxSumArray {

    public static void main(String[] args) {
        int[] a = {1,-1,2,-1,3};
        System.out.println(maxSumArray(a));
    }

    private static int maxSumArray(int[] a) {

        int maxSum = Integer.MIN_VALUE;
        int curSum = 0;

        for(int i=0;i<a.length;i++){
            curSum = curSum + a[i];
            if(curSum>maxSum){
                maxSum=curSum;
            }
        }
        return maxSum;
    }

}
