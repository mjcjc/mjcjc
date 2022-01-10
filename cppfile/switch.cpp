//=========================================================================
#include<iostream>

using std::cout;
using std::endl;
using std::cin;

int main() {
	int user_input;
	cout << "저의 정보를 표시해줍니다." << endl;
	cout << "1.이름" << endl;
	cout << "2.나이" << endl;
	cout << "3.성별" << endl;
	cin >> user_input;

	switch (user_input)	{
	case 1:
		cout << "MJG !" << endl;
		break;

	case 2:
		cout << "XX살" << endl;
		break;

	case 3:
		cout << "남자" << endl;
		break;

	default:
		cout << "궁금한게 없군요~" << endl;
		break;
	}
	return 0;
}
/*=========================================================================
kor explain source code 
추후 업로드 시킬 예정. 
대충 스위치문에 1 2 3 번까지 정해두고 번호를 칠 시 그에 맞는 값을 출력한 후 종료 시키는 switch문임.
=========================================================================
eng explain source code

will be write it down
=========================================================================*/
