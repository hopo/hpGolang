// ...ing
#include <iostream>
#include <string>
#include <vector>
#include <map>

using namespace std;

void sort_dictionary(map<string, int> ncard) {
	for(auto &x : ncard) { // &  what???
		cout << x << " - ";
		cout << x.first << " : " << x.second << endl;
	}
	
 	/*
	vector<string> 
	for k, _ := range nc {
		slis = append(slis, k)
	}
	sort.Strings(slis)
	for _, k := range slis {
		ret = append(ret, Ncard{k, nc[k]})
	}
	return
	*/
}

int main() {
	map<string, int> ncard = {{"김철수", 78}, {"이하나", 97}, {"정진원", 88}};
	sort_dictionary(ncard);
}
