package arraytest;

import java.util.ArrayList;

public class GetExam {
	public static void main(String[] args) {
		// create an empty array lsit with an initial capacity
		ArrayList<Integer> arrlist = new ArrayList<Integer>(5);
		
		// use add() method to add elements in the list
		arrlist.add(15);
		arrlist.add(22);
		arrlist.add(30);
		arrlist.add(40);
		
		// let us print all the elements available in list
		for(Integer number: arrlist) {
			System.out.println("Number = " + number);
		}
		
		// retrives element at 4th position
		int retval = arrlist.get(3);
		System.out.println("Retrived element is = " + retval);
	}
}
