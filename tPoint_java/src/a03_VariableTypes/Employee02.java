package a03_VariableTypes;
import java.io.*;

public class Employee02 {

	// salary variable is a private static variable
	private static double salary;

	// DEPARTMENT is a constant
	public static final String DEPARTMENT = "Developement ";

	public static void main(String args[]) {
		salary = 1000;
		System.out.println(DEPARTMENT + "average salary:" + salary); 
	}
}
