import java.io.*;
import java.util.*;

/*
 * To execute Java, please define "static void main" on a class
 * named Solution.
 *
 * If you need more classes, simply define them inline.
 */

class CaserCipher {
  public static void main(String[] args) {
    
    String alpha = "abcdefghijklmnopqrstuvwxyz";
   String ip = "xyz";
    int index = 2;
    
    int newKey = index % 25;
    
     for(int i =0; i < ip.length() ; i++){
        System.out.println(CaserCipher.getCasercipher(ip.charAt(i),newKey,alpha));
     }
  }
  
  public static char getCasercipher(char ip,int index,String alpha){
  
     int newCharVal  = alpha.indexOf(ip) + index ;
      
    return   (newCharVal <=25)? alpha.charAt(newCharVal) 
      : alpha.charAt(-1 + newCharVal%25);       
  
  }
}

