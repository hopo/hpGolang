package linkedlist;

import java.util.LinkedList;

// what is different? offer() and add()

public class OfferExam {
	public static void main(String[] args) {
		// create a LinkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add some elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// print the list
		System.out.println("LinkedList: " + list);
		
		// offer a new element
		list.offer("Element");
		
		// print the new list
		System.out.println("LinkedList: " + list);
	}
}
