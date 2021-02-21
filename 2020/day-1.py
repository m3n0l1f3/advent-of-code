#!/usr/bin/python3

FILE_PATH = 'day-1.txt'

def two_sum(data, target):
  mapper = set()

  for num in data:
    other_num = target - num
    if other_num in mapper:
      return other_num * num
    mapper.add(num)
  return 0

def three_sum(data, target):
  for count, _ in enumerate(data):
    data_iterator = iter(data)
    for _ in range(count - 1):
      next(data_iterator)
    num = next(data_iterator)
    new_target = target - num
    result = two_sum(data_iterator, new_target)
    if result != 0:
      return result * num
  return 0

with open(FILE_PATH, 'r') as file:
  lines = file.readlines()
  data = list(map(lambda line: int(line.strip()), lines))
  part_1 = two_sum(data, 2020)
  print('Part 1: ', part_1)
  part_2 = three_sum(data, 2020)
  print('Part 2: ', part_2)
