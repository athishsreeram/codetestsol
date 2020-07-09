class Sumoftwo {
  public static void main(String[] args) {
    int arr[] = {1,2,4,5,6};
    
    int res[] = Sumoftwo.sumoftwo(arr,10);
    
    for(int i=0;i<res.length;i++)
      System.out.println(res[i]);

    
  }
  
  public static int[] sumoftwo(int arr[],int total){
 
    Map<Integer,Boolean> nolst = new HashMap<Integer,Boolean>(); 
    int res[]= new int[2];

    for(int i=0;i<arr.length;i++){

      int a = arr[i];
      int b = total -a;

      if(nolst.get(b) != null && nolst.get(b) == true){
        res = new int[]{a,b};
        return res;
      }else{
        nolst.put(a,true);
      }
    }


    return res;  
  
  }
}
