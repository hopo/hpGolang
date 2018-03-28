package linkedlist;

import java.util.LinkedList;

public class PeekExam {
	public static void main(String[] args) {
		// create a LinkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add some elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// printl the list
		System.out.println("LinkedList: " + list);
		
		// peek at the head of the list
		System.out.println("Head of the list: " + list.peek());
	}
}
