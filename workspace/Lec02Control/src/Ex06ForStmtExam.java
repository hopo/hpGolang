
public class Ex06ForStmtExam {
	public static void main(String[] args) {

		final String ST = "*";
		final String LN = "\n";
		final String WS = " ";
		int count = 10;
		String ret = "";

		// // Ex2
		// System.out.println("========== Ex2 ==========");
		// ret = "";
		// for(int i = count; i > 0; i--) {
		// for(int j = 0; j < i ; j++) {
		// ret += ST;
		// }
		// ret += "\n";
		// }
		// System.out.print(ret);

		// // Ex2-02
		// System.out.println("========== Ex2-02 ==========");
		// ret = "";
		// for(int i = 0; i < count; i++) {
		// for(int j = 0; j < count-i; j++) {
		// ret += ST;
		// }
		// ret += "\n";
		// }
		// System.out.print(ret);

		// // Ex3
		// System.out.println("========== Ex3 ==========");
		// ret = "";
		// for(int i = 0; i < count; i++) {
		// for(int j = 0; j < i+1 ; j++) {
		// ret += ST;
		// }
		// ret += "\n";
		// }
		// System.out.print(ret);

		// // Ex3-02
		// // ..ing
		// System.out.println("========== Ex3-02 ==========");
		// ret = "";
		// int j = 0;
		// for(int i = 0; i < 30; i++) {
		// if(i != j) {
		// ret += ST;
		// j++;
		// continue;
		// } else {
		// j = 0;
		// ret += "\n";
		// }
		//
		// }
		// System.out.print(ret);

		// // Ex4
		// System.out.println("========== Ex4 ==========");
		// ret = "";
		// for(int i = 0; i < count; i++) {
		// for(int j = 0; j < count ; j++) {
		// if(count-i > j+1) {
		// ret += WS;
		// } else {
		// ret += ST;
		// }
		// }
		// ret += "\n";
		// }
		// System.out.print(ret);

		// // Ex4-02
		// System.out.println("========== Ex4-02 ==========");
		// for(int i = 0; i < count; i++) {
		// for(int j = 0; j < count-i-1; j++) {
		// System.out.print(WS);
		// }
		// for(int k = 0; k <= i; k++) {
		// System.out.print(ST);
		// }
		// System.out.println();
		// }

		// // Ex5
		// System.out.println("========== Ex5 ==========");
		// ret = "";
		// for(int i = 0; i < count; i++) {
		// for(int j = 0; j < count ; j++) {
		// if(i < j+1) {
		// ret += ST;
		// } else {
		// ret += WS;
		// }
		// }
		// ret += "\n";
		// }
		// System.out.print(ret);

		// // Ex5-02
		// System.out.println("========== Ex5-02 ==========");
		// for(int i = 0; i < count; i++) {
		// for(int j = 0; j < i; j++) {
		// System.out.print(WS);
		// }
		// for(int k = 0; k < count-i; k++) {
		// System.out.print(ST);
		// }
		// System.out.println();
		// }

//		// Ex6
//		System.out.println("========== Ex6 ==========");
//		ret = "";
//		for (int i = 0; i < count; i++) {
//			for (int j = 0; j < i; j++) {
//				ret += WS;
//			}
//			for (int k = 0; k < count - i; k++) {
//				ret += ST;
//			}
//			for (int l = 0; l < count - i - 1; l++) {
//				ret += ST;
//			}
//			ret += "\n";
//		}
//		System.out.print(ret);

//		// Ex6-02
//		System.out.println("========== Ex6-02 ==========");
//		ret = "";
//		for (int i = 0; i < count; i++) {
//			for (int j = 0; j < i; j++) {
//				System.out.print(WS);
//			}
//			for(int k = 0; k < (count-i)*2-1; k++) {
//				System.out.print(ST);
//			}
//			System.out.println();
//		}

//		// Ex7
//		System.out.println("========== Ex7 ==========");
//		// uppper
//		for(int i = 0; i < count; i++) {
//			for(int j = 0; j < count-i-1; j++) {
//				System.out.print(WS);
//			}
//			// left char
//			int k = 0;
//			for(k = 0; k < i+1; k++) {
//				System.out.print((char)('A'+k));
//			}
//			// right char, depend on k value
//			for(int n = k - 1; n >= 0; n--) {
//				System.out.print((char)('A'+n));
//			}
//			System.out.println();
//		}
//		// lower
//		for(int i = 0; i < count; i++) {
//			for(int j = 0; j < i; j++) {
//				System.out.print(WS);
//			}
//			int k = 0;
//			for(k = 0; k < count-i ; k++) {
//				System.out.print((char)('A'+k));
//			}
//			for(int n = k - 1; n >= 0; n--) {
//				System.out.print((char)('A'+n));
//				
//			}
//			System.out.println();
//		}

		// Ex8
		// ..ing
		System.out.println("========== Ex8 ==========");
		// uppper
		for(int i = 0; i < count; i++) {
			for(int j = 0; j < count-i-1; j++) {
				System.out.print(WS);
			}
			// left char
			int k = 0;
			for(k = 0; k < i+1; k++) {
				System.out.print((char)('A'+k));
			}
			// right char, depend on k value
			for(int n = k - 1; n >= 0; n--) {
				System.out.print((char)('A'+n));
			}
			System.out.println();
		}
		// lower
		for(int i = 0; i < count; i++) {
			for(int j = 0; j < i; j++) {
				System.out.print(WS);
			}
			int k = 0;
			for(k = 0; k < count-i ; k++) {
				System.out.print((char)('A'+k));
			}
			for(int n = k - 1; n >= 0; n--) {
				System.out.print((char)('A'+n));
				
			}
			System.out.println();
		}
	}
}
