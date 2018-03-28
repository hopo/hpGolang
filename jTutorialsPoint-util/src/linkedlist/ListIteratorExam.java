package linkedlist;

import java.util.Iterator;
import java.util.LinkedList;

public class ListIteratorExam {
	public static void main(String[] args) {
		// create a LinkedList
		LinkedList<Object> list = new LinkedList<>();
		
		// add sine elements
		list.add("Hello");
		list.add(2);
		list.add("Chocolate");
		list.add("10");
		
		// print the list
		System.out.println("LinkedList: " + list);
		
		// set Iterator at specified index
		Iterator<Object> x = list.listIterator(1); // from index 1
		
		// print list with the iterator
		while(x.hasNext()) {
			System.out.println(x.next());
		}
	}
}
