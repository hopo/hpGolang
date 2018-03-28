package linkedlist;

import java.util.LinkedList;

public class PopExam {
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
		
		// retrieve and remove the head of the list
		System.out.println("Pop element of the list: " + list.pop());
		
		// print the list
		System.out.println("LinkedList: " + list);
	}

}
