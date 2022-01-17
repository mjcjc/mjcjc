#include<iostream>


class Date {     //Date 클래스 선언
	int year_;     //int 형 year_변수 선언
	int month_;    //int 형 month_변수 선언
	int day_;      //int 형 day_변수 선언

public:          //public 멤버 선언.
	void SetDate(int year, int month, int date);    //void 형 함수 SetDate 선언후 매개변수 int형 year,month,date선언
	void AddDay(int inc);                           //void 형 함수 AddDay 선언후 매개변수 int형 inc선언
	void AddMonth(int inc);                         //void 형 함수 AddMonth 선언후 매개변수 int형 inc선언
	void AddYear(int inc);                          //void 형 함수 AddYear 선언후 매개변수 int형 inc선언

	int  GetCurrentMonthTotalDays(int year, int month);   //int형 함수 GetCurrentMonthTotalDays 선언후 매개변수 int형 year,month선언

	void ShowDate();    //void형 함수 ShowDate선언
};

void Date::SetDate(int year, int month, int day) { //Date클래스 안에 있는 SetDate함수 선언후 매개변수 int형 year,month,day 선언
	year_ = year;                                    //year_이 year값으로 초기화함.
	month_ = month;                                  //month_가 month값으로 초기화함.
	day_ = day;                                      //day_가 day값으로 초기화함
}

int Date::GetCurrentMonthTotalDays(int year, int month) {                   //Date클래스 안에 있는 GetCurrentMonthTotalDays함수 선언후 매개변수 int형 year,month 선언
	static int month_day[12] = { 31,28,31,30,31,30,31,31,30,31,30,31 };       //static 형식인 int형 변수 month_day를 1차원 배열로 선언후 배열값에 맞게 모두 초기화
		if (month != 2) {                                                       //조건문 if 선언후 month 가 2가 아닐시
			return month_day[month - 1];                                          // return 값인 배열 month_day에서의 값이 month에서 로 반환
		}
		else if (year % 4 == 0 && year % 100 != 0) {                            // 그렇지 않을시 year 값이 4로 나눴을때 나머지가 0 없거나 year 값이 100으로 나눴을때 나머지가 0이 아닐경우
			return 29;                                                            // return 값을 29로 반환
		}
		else {                                                                  // 만약 위 2가지의 경우가 아닐시
			return 28;                                                            // return 값을 28로 반환.
		}
}

void Date::AddDay(int inc) {                                                    // Date클래스에 있는 함수 AddDay와 int형 매개변수 inc를 선언 
	while (true) {                                                                // while반복문을 무한으로 돌림.
		int currnet_month_total_days = GetCurrentMonthTotalDays(year_, month_);     // int형 currnet_month_total_days값이 GetCurrentMonthTotalDays(year_, month_) 값으로 초기화

		if (day_ + inc <= currnet_month_total_days) {                               //조건문 if선언 day_ 
			day_ += inc;
			return;
		}
		else {
			inc -= (currnet_month_total_days - day_ + 1);
			day_ = 1;
			AddMonth(1);
		}
	}
}
void Date::AddYear(int inc) {
	year_ += inc; 
}

void Date::AddMonth(int inc) {
	AddYear((inc + month_ - 1) / 12);
	month_ = month_ + inc % 12;
	month_ = (month_ == 12 ? 12 : month_ % 12);
}

void Date::ShowDate() {
	std:: cout << "오늘은 " << year_ << " 년" << month_ << " 월" << day_ << " 일 입니다." << std::endl;
	std::printf("");
}

int main() {
	Date day;
	day.SetDate(2022, 1, 18);
	day.ShowDate();

	day.AddYear(10);
	day.ShowDate();

	day.AddDay(30);
	day.ShowDate();

	day.AddDay(2000);
	day.ShowDate();
	return 0;
}
/*==========================================================================================================================================
설명은 추후 업로드 예정*/
