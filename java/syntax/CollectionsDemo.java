package syntax;

import java.util.Arrays;
import java.util.Stack;

public class CollectionsDemo {
   public static void main(String[] args) {
       Stack<Integer> stack = new Stack<>();
       stack.push(1);
       if(!stack.isEmpty()){
           int top = stack.peek();// inspect
           int topEle = stack.pop();// pop
       }
       int[] arr = new int[10];
       Arrays.sort(arr);
   } 
}
