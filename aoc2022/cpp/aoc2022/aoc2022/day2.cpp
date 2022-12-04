#include <fstream>
#include <vector>
#include <optional>
#include <algorithm>
#include <numeric>
#include <iostream>
#include <filesystem>
#include <unordered_map>
#include <functional>

using move_list = std::vector<std::pair<char, char>>;

auto read_input(const std::filesystem::path& path) -> std::optional<move_list> {
  if (std::filesystem::exists(path)) {
    std::ifstream in(path);
    move_list moves;

    if (in.is_open()) {
      while (in) {
        char opponent_move, my_move;

        if (in >> opponent_move >> my_move) {
          moves.push_back(std::pair<char, char>(opponent_move, my_move));
        }
      }
      return moves;
    }
  }

  return std::nullopt;
}

auto rps_play = [](int acc, std::pair<char, char> move) {
  std::unordered_map<char, int> my_value_map = {
    {'X', 1},
    {'Y', 2},
    {'Z', 3}
  };

  switch (move.second - move.first - 23) {
  case 1:
  case -2: acc += my_value_map[move.second] + 6; break;
  case 0: acc += my_value_map[move.second] + 3; break;
  default: acc += my_value_map[move.second]; break;
  }
  return acc;
};

auto elf_play = [](int acc, std::pair<char, char> move) {
  char my_move = 0;

  switch (move.second) {
  case 'X': my_move = move.first == 'A' ? move.first + 25 : move.first + 22; break;
  case 'Y': my_move = move.first + 23; break;
  case 'Z': my_move = move.first == 'C' ? move.first + 21 : move.first + 24; break;
  }

  acc += rps_play(0, std::pair<int, int>(move.first, my_move));

  return acc;
};

auto run_moves(const move_list& moves, std::function<int(int, std::pair<char, char>)> func) -> int {
  return std::accumulate(moves.cbegin(), moves.cend(), 0, func);
}

void Day2() {
  std::cout << "====== Day 2 Start ======\n";
  auto moves = read_input("input/day2");

  if (moves) {
    std::cout << run_moves(*moves, rps_play) << "\n";
    std::cout << run_moves(*moves, elf_play) << "\n";
  } else {
    std::cout << "[NO INPUT]\n";
  }

  std::cout << "-------------------------" << std::endl;
}
