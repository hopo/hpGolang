// vector::emplace
#include <iostream>
#include <vector>

using namespace std;

int main () {
  vector<int> myvector = {10, 20, 30};

  auto it = myvector.emplace(myvector.begin()+1, 100);
  myvector.emplace(it, 200);
  myvector.emplace(myvector.end(), 300);

  cout << "myvector contains:";
  for (auto& x : myvector) { cout << ' ' << x; }
  cout << '\n';

  return 0;
}
