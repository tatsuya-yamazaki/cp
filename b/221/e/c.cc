#include<bits/stdc++.h>
using namespace std;

int main() {
	string N = "1010";
	sort(N.begin(),N.end());
	int ans = 0;
	do {
		for(int i=1; i<N.size(); i++) {
			int l = 0, r = 0;
			for(int j=0; j<i; j++) {
				cout << "beforel:" << l << endl;
				cout << "N[j]:" << N[j] << endl;
				cout << "code \'0\':" << +'0' << endl;
				cout << "code N[j]:" << +N[j] << endl;
			       	l = l*10+N[j] - '0';
			       	//l = l*10+N[j];
				cout << "afterl:" << l << endl;
			}
			for(int j=i; j<N.size(); j++) r = r*10+N[j] - '0';
			ans = max(ans,l*r);
		}

	} while(next_permutation(N.begin(), N.end()));
	cout << ans << endl;
}
