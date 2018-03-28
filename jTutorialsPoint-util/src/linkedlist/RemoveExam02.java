package linkedlist;

import java.util.LinkedList;

public class RemoveExam02 {
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
		
		// remove "10"
		// and return boolean
		System.out.println("10 is in the list: " + list.remove("10"));
		
		// print the list
		System.out.println("LinkedList: " + list);
	}
}
