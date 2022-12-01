import std.core;

import day1;

std::vector<std::function<void()>> days = {
  Day1,
};

auto main(int argc, char *const argv[]) -> int {

  if (argc < 2) {
    for (auto func : days) {
      func();
    }
  } else {

    int run_day = 0;
    std::cout << "Choose a day to run between 1 and " << days.size() << ": ";
    std::cin >> run_day;

    if (run_day > 0 && run_day <= days.size()) {
      days[run_day - 1]();
    } else {
      std::cout << run_day << " is an invalid day.\n";
    }
  }

  return 0;
}
