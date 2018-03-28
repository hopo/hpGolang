package linkedlist;

import java.util.LinkedList;

// confer pop(), poll() and remove()

public class PollExam {
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
		System.out.println("Head element of the list: " + list.poll());
		
		// print the list
		System.out.println("LinkedList: " + list);
	}

}
