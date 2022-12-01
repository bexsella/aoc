export module day1;

import std.core;
import std.filesystem;

[[nodiscard]] auto read_input(const std::filesystem::path& path) -> std::optional<std::vector<int>> {
  std::fstream in(path);
  std::vector< int> output;

  in >> std::noskipws;

  if (in.is_open()) {
    while (in) {
      int cals;
      int cal_count = 0;

      while (in >> cals) {
        cal_count += cals;
        in.ignore(1, '\n');

        if (in.peek() == '\n') {
          in.ignore(1, '\n');
          break;
        }
      }

      output.push_back(cal_count);
    }
  } else {
    return std::nullopt;
  }

  return output;
}

auto sum_fat_elves(const std::vector<int>& elves, const size_t fat_elf_count = 1) -> int {
  return std::accumulate(elves.begin(), elves.begin() + fat_elf_count, 0);
}

export void Day1() {
  std::cout << "====== Day 1 Start ======\n";

  auto elves = read_input("input/day1");

  if (elves) {
    std::ranges::sort(*elves, std::greater{});

    std::cout << sum_fat_elves(*elves) << "\n";
    std::cout << sum_fat_elves(*elves, 3) << "\n";
  } else {
    std::cout << "[NO INPUT]\n";
  }

  std::cout << "-------------------------" << std::endl;
}
