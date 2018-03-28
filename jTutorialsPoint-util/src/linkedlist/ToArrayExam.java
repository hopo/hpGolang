package linkedlist;

import java.util.LinkedList;

public class ToArrayExam {
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
		
		// create an array and copy  the list to it
		Object[] array =  list.toArray();
		
		// print the array
		for(int i = 0; i < list.size(); i++) {
			System.out.println("Array: " + array[i]);
		}
		
	}
}
