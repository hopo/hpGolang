package linkedlist;

import java.util.*;

public class ElementExam {
	public static void main(String[] args) {
		// creaete a LinkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add some elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// print the list
		System.out.println("LinkedList:" + list);
	}
}
