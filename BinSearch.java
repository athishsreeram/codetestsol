public class BinSearch {

    public static void main(String[] args) {
        int a[]={1,3,5,10};


    }

    public int bs(int a[],int target){

        int i = 0 ;
        int j = a.length-1;

        int mid = i+j/2;

        while(i<j){

            if(a[mid]==target){
                return mid;
            }else if(a[mid]<target){
                i = mid+1;
            }else{
               j = mid -1;
            }

        }

        return -1;
    }


}
