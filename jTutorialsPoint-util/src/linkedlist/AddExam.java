package linkedlist;

import java.util.LinkedList;

public class AddExam {
	public static void main(String[] args) {
		// create a LInkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add some elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// print the list
		System.out.println("LinkedList: " + list);
		
		// add a new element at the end fo the list
		list.add("Element");
		
		// print the updated list
		System.out.println("LinkedList: " + list);

	}
}
