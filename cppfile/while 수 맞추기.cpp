//=================================================================================
#include<iostream>

int main() {
	int lucky_number = 3;
	std::cout << "내 비밀 수를 맞추어 보세요~" << std::endl;

	int user_input;

	while (1) {
		std::cout << "입력 : ";
		std::cin >> user_input;
		if (lucky_number == user_input) {
			std::cout << "맞추셨습니다~~" << std::endl;
			break;
		}
		else {
			std::cout << "다시 생각해보세요" << std::endl;
		}
	}
	return 0;
}
/*=================================================================================
kor explain source code
2번 줄에서 헤더파일 iostream 선언.
4번 줄에서 메인함수 선언
5번 줄에서 int형 lucky_number을 선언한후 3으로 초기화
6번 줄에서 "내 비밀 수를 맞추어 보세요~~"를 출력후 줄 바꿈.
8번 줄에서 int형 user_input을 선언
10번 줄에서 while 반복문 선언
11번 줄에서 "입력 : " 출력을 함.
12번 줄에서 cin이라는 입력을 user_input 자료형에 입력 받음.
13번 줄에서 조건문 if를 선언 lucky_number와 user_input이 같을경우라 작성.
14~15번 줄에서 맞추셨습니다를 출력한 후 줄 바꾼뒤 프로그램을 멈춤.
17~18번 줄에서 그게 아닐경우 다시 생각해보세요를 말 한뒤 값을 반복시킴.
=================================================================================
eng explain source code.

will be write it down.
=================================================================================*/
