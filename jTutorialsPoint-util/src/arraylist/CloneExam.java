package arraylist;

import java.util.ArrayList;

public class CloneExam {
	public static void main(String[] args) {
		// create an empty array list
		ArrayList<StringBuilder> arrlist1 = new ArrayList<StringBuilder>();
		
		// use add for new value
		arrlist1.add(new StringBuilder("Learning-"));
		
		// using clone to affect the objects pointed to by the reference.
		ArrayList arrlist2 = (ArrayList)arrlist1.clone();
		
		// appending the string
		StringBuilder strbuilder = arrlist1.get(0);
		strbuilder.append("list1, list2-both pointing to the same StringBuilder");
		
		System.out.println("The 1st list prints: ");
		
		// both lists will print the same value, printing list1
		for(int i = 0; i < arrlist1.size(); i++) {
			System.out.print(arrlist1.get(i) + " ");
		}
		
		System.out.println("The 2nd list prints ths same i.e:");
		
		// both lists will print the same value, printing list2
		for(int i = 0; i < arrlist2.size(); i++) {
			System.out.println(arrlist2.get(i));
		}
	}
}
