//=========================================================
#include<iostream>

int main() {

	int sum = 0;
	
	for (int i = 1; i <= 10; i++) {
		sum += i;
	}

	std::cout << "합은 : " << sum << std::endl;
	return 0;
}
/*=========================================================
kor explain source code
2번 줄에서 iostream 이라는 헤더파일을 선언
4번 줄에서 메인함수 선언
6번 줄에서 int형 sum 변수를 0으로 초기화
8번 줄에서 반복문 for에 int i 변수를 1로 시작하여 10까지 i를 차례로 반복하여 출력함. 
9번 줄에서 sum변수에 i변수 값을 계속 더함.
12번 줄에서 "합은 : " 이라는 말을 먼저 출력후 sum 값을 출력한 후 한 줄을 줄 바꿈 함.
13번 줄에서  return 0 값을 부여.
=========================================================
eng explain source code
will be write it down
=========================================================*/
