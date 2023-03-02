public class BinarySearch {

    public static void main(String[] args) {
        int a[]= {4,5,10,30,70};

        System.out.println(binarySearch(a,0,a.length-1,0));
    }

    private static int binarySearch(int a[],int left, int right, int x) {

        while(left<right){
            int mid = left + (right-left)/2;

            if(x==a[mid]){
                return mid;
            } else if (x < a[mid]) {
                right = mid-1;
            }else if (x > a[mid]) {
                left = mid+1;
            }
            binarySearch(a,left,right,x);
        }

        return -1;
    }


}
