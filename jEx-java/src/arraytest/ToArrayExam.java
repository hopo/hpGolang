package arraytest;

import java.util.ArrayList;

public class ToArrayExam {
	public static void main(String[] args) {
		// create an empty array list with an initial capacity
		ArrayList<Integer> arrlist = new ArrayList<Integer>();
		
		// use add() method to add values in the list
		arrlist.add(10);
		arrlist.add(12);
		arrlist.add(31);
		arrlist.add(49);
		
		System.out.println("Printing elements of arrlist");
		
		// let us print all the elements availabel in list
		for(Integer number: arrlist) {
			System.out.println("Number = " + number);
		}
		
		// toArray copies content into othe array
		Integer list2[] = new Integer[arrlist.size()];
		list2 = arrlist.toArray(list2);
		
		System.out.println("Printing elements of list2");
		
		// let us print all the elements available in list
		for(Integer number: list2) {
			System.out.println("Numbher = " + number);
		}
	}

}
