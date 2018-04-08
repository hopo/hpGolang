package level02;

public class GetDayName {
	public static void main(String[] args) {
		int m1 = 5, d1 = 24;
		String ex1 = getdn(m1, d1); // "TUE"
		System.out.println(ex1);

		int m2 = 11, d2 = 5;
		String ex2 = getdn(m2, d2); // "SAT"
		System.out.println(ex2);
	}
	
	static String getdn(int mon, int day) {
		final int[] monthd = {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
		final String[] dnames = {"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"};
		int r;
		String ret;
		
		// default: January 1 is Friday.

		// how much days from Jan 1.
		for(int i = 0; i < mon-1; i++){
			day += monthd[i];
		}
		
		// find what day in week
		r = day%7;
		if(r < 3){
			ret = dnames[r+4];
		} else {
			ret = dnames[r-3];
		}

		return ret;
	}
}