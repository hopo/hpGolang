package linkedlist;

import java.util.LinkedList;

public class IndexOfExam {
	public static void main(String[] args) {
		// create a LinkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add some elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// print the list
		System.out.println("LinkedList: "+ list);
		
		// get the index for "Chocolate"
		System.out.println("Index for Chocolate: " + list.indexOf("Chocolate"));
		
		// get the index for "Coffee"
		System.out.println("Index for Coffee: "+ list.indexOf("Coffee"));
	}
}