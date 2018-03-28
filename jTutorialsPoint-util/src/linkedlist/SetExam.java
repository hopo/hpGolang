package linkedlist;

import java.util.LinkedList;

public class SetExam {
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
		
		// set "Coffee" at index 1
		System.out.println("Object to be replaced: " + list.set(1, "Coffee"));
		
		// print the list
		System.out.println("LinkedList: " + list);
	}
}
