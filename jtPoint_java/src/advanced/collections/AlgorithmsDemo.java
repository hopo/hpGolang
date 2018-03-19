package advanced.collections;

// import java.util.Collections;
// import java.util.Comparators;
// import java.util.Iterator;
// import java.util.LinkedList;
import java.util.*;

public class AlgorithmsDemo {
    public static void main(String args[]) {
        // Create and initialize linked list
        LinkedList ll = new LinkedList();
        ll.add(new Integer(-8));
        ll.add(new Integer(20));
        ll.add(new Integer(-20));
        ll.add(new Integer(8));

        // Create a reverse order comparator
        Comparator r = Collections.reverseOrder();

        // Sort list by using the comparato
        Collections.sort(ll, r);

        // Get iterator
        Iterator li = ll.iterator();
        System.out.print("List soreted in reverse: ");

        while(li.hasNext()) {
            System.out.print(li.next() + " ");
        }
        System.out.println();
        Collections.shuffle(ll);

        // display randomized list
        li = ll.iterator();
        System.out.print("List shffled: ");

        while(li.hasNext()) {
            System.out.print(li.next() + " ");
        }
        
        System.out.println();
        System.out.println("Minimum: " + Collections.min(ll));
        System.out.println("Maximum: " + Collections.max(ll));
    }
}
