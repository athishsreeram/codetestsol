class Plaendrom {
  public static void main(String[] args) {
    
    String ip = "abcdzba";
  
    boolean pal = true;
     
    int i = 0;
    int j = ip.length() -1;
    
    while(i < j){
      if(ip.charAt(i) != ip.charAt(j)){
        pal = false;
      }
      i++;
      j--;
    }
        
    System.out.println(pal);
  }
  
  
}
