#include <iostream>
#include <vector>
#include <functional>

extern void Day1();
extern void Day2();
extern void Day3();

std::vector<std::function<void()>> days = {
  Day1,
  Day2,
  Day3
};

auto main(int argc, char *const argv[]) -> int {
  if (argc < 2) {
    int run_day = 0;
    std::cout << "Choose a day to run between 1 and " << days.size() << ", or 0 to run all: ";
    std::cin >> run_day;

    if (run_day > 0 && run_day <= days.size()) {
      days[run_day - 1]();
    } else if (run_day == 0) {
      for (auto func : days) {
        func();
      }
    } else {
      std::cout << run_day << " is an invalid day.\n";
    }
  } else {
    int run_day = std::atoi(argv[1]);
    if (run_day > 0 && run_day <= days.size()) {
      days[run_day - 1]();
    } else {
      std::cout << run_day << " is an invalid number.\n";
    }
  }

  return 0;
}
