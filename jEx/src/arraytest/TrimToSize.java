package arraytest;

import java.util.ArrayList;

public class TrimToSize {
	public static void main(String[] args) {
		// create an empty array list with an initial capacity
		ArrayList<Integer> arrlist = new ArrayList<Integer>(5);
		
		// use add() method to add elements in the list
		arrlist.add(35);
		arrlist.add(20);
		arrlist.add(25);
		
		// Trim the arraylist
		arrlist.trimToSize();

		// let us print all the elements available in list
		for(Integer number: arrlist) {
			System.out.println("Number = " + number);
		}
	}
}
