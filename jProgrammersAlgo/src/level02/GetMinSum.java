package level02;

public class GetMinSum {
	public static void main(String[] args) {
		int[] arra1 = {1, 2}, arrb1 = {3, 4};
		int ex1 = getmm(arra1, arrb1); // 10
		System.out.println(ex1);

		int[] arra2 = {4, 2}, arrb2 = {3, 9};
		int ex2 = getmm(arra2, arrb2); // 30
		System.out.println(ex2);

		int[] arra3 = {2, 3}, arrb3 = {5, 7};
		int ex3 = getmm(arra3, arrb3); // 29
		System.out.println(ex3);
	}
	
	static int getmm(int[] arra, int[] arrb) {
		int[] smpl = new int[arra.length*arrb.length];
		int idx = 0;
		for(int a : arra) {
			for(int b : arrb) {
				smpl[idx++] = a*b;
			}
		}
		
		int min = smpl[0] + smpl[idx-1];
		for(int i = 0; i < idx/2; i++) {
			if(min > smpl[i] + smpl[idx-1-i]) {
				min = smpl[i] + smpl[idx-1-i];
			}
		}

		return min;
	}
}
