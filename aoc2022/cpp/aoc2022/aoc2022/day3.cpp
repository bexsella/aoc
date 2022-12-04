#include <filesystem>
#include <string>
#include <optional>
#include <fstream>
#include <vector>
#include <iostream>
#include <algorithm>
#include <numeric>
#include <set>
#include <ranges>
#include <functional>
#include <tuple>

[[nodiscard]] auto read_input(const std::filesystem::path& path) -> std::optional<std::vector<std::string>> {
  if (std::filesystem::exists(path)) {
    std::ifstream in(path);

    if (in.is_open()) {
      std::vector<std::string> input;
      std::string rucksack;

      while (in >> rucksack) {
        input.push_back(rucksack);
      }
      return input;
    }
  }
  return std::nullopt;
}

int sum_intersection(int acc, char c) {
  acc += c - (std::islower(c) ? 96 : 38);
  return acc;
}

auto sum_group(const std::vector<std::string>& rucksacks) -> int {
  using group_sets = std::tuple<std::set<char>, std::set<char>, std::set<char>>;
  std::vector<group_sets> new_list;
  int acc = 0;

  // Surely this cannot be the way, 
  for (int i =0; i < rucksacks.size(); i+=3) {
    new_list.push_back(group_sets(
      std::set<char>(rucksacks[i].cbegin(), rucksacks[i].cend()), 
      std::set<char>(rucksacks[i + 1].cbegin(), rucksacks[i + 1].cend()),
      std::set<char>(rucksacks[i + 2].cbegin(), rucksacks[i + 2].cend()))
    );
  }

  for (auto t : new_list) {
    std::list<char> temp_set;
    std::list<char> result;

    std::set_intersection(std::get<0>(t).cbegin(), std::get<0>(t).cend(),
      std::get<1>(t).cbegin(), std::get<1>(t).cend(),
      std::back_inserter(temp_set));

    std::set_intersection(std::get<2>(t).cbegin(), std::get<2>(t).cend(),
      temp_set.cbegin(), temp_set.cend(),
      std::back_inserter(result));

    acc += std::accumulate(result.cbegin(), result.cend(), 0, sum_intersection);
  }

  return acc;
}

auto sum_rucksack(const std::vector<std::string>& rucksack) -> int {
  int acc = 0;

  for (auto r : rucksack) {
    auto compartment_a = std::set<char>(r.cbegin(), std::next(r.cbegin(), r.length() / 2));
    auto compartment_b = std::set(std::next(r.cbegin(), r.length() / 2), r.cend());
    std::list<char> intersection;

    std::set_intersection(compartment_a.cbegin(), compartment_a.cend(), compartment_b.cbegin(), compartment_b.cend(), std::back_inserter(intersection));

    acc += std::accumulate(intersection.cbegin(), intersection.cend(), 0, sum_intersection);
  }

  return acc;
}


void Day3() {
  std::cout << "====== Day 3 Start ======\n";
  auto input = read_input("input/day3");

  if (input) {
    std::cout << sum_rucksack(*input) << "\n";
    std::cout << sum_group(*input) << "\n";
  } else {
    std::cout << "[NO INPUT]\n";
  }

  std::cout << "-------------------------" << std::endl;
}
