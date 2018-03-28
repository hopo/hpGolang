package linkedlist;

import java.util.*;

public class DescendingIterator {
	public static void main(String[] args) {
		// create a LinkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add some elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// print the list
		System.out.println("LinkedList:" + list);
		
		// set Iterator as descending
		Iterator<Object> x =  list.descendingIterator();
		
		// print list with desceding order
		while(x.hasNext()) {
			System.out.println(x.next());
		}
	}

}
