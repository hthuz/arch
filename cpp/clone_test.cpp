
#include <iostream>
#include <vector>

using namespace std;




vector<int> get_odd(vector<int> arr) {
	vector<int> res;
	for (int i = 0; i < arr.size(); i++) {
		if (arr[i] % 2 == 1) {
			// Address is different, which means that it is a copy
			res.push_back(arr[i]);
		}
	}
	return res;
}

vector<int> get_odd2(vector<int> arr) {
	vector<int> res;
	for (int i = 0; i < arr.size(); i++) {
		if (arr[i] % 2 == 1) {
			// Same as above
			// Assign arr[i] new a new variable means creating a copy
			cout << "input " << &arr[i] << endl;
			int tmp = arr[i];
			cout << "tmp " << &tmp << endl;
			res.push_back(tmp);
			cout << "output " << &res.back() << endl;
		}
	}
	return res;
}

vector<int*> get_odd3(vector<int*> &arr) {
	vector<int*> res;
	for (int i = 0; i < arr.size(); i++) {
		if (*arr[i] % 2 == 1) {
			cout << "input val " << arr[i] << " addr " << &arr[i] << endl;
			res.push_back(arr[i]);
			cout << "output val " << res.back() << " addr " << &res.back() << endl;
		}
	}
	return res;
}



int main() {

	vector<int> arr{1,2,3,4,5,6,7};

	vector<int> res = get_odd2(arr);

	cout << endl;
	for (int i = 0; i < res.size(); i++) {
		cout << res[i] << " ";
	}

	cout << endl;
	vector<int*> arr2{new int (1), new int(2), new int(3)};
	vector<int*> res2 = get_odd3(arr2);
	// arr2.push_back(new int(1))
	*arr2[0] = 100;
	for (int i = 0; i < res2.size(); i++) {
		cout << *res2[i] << " ";
	}

}
