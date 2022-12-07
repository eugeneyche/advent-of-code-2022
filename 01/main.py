with open("input.txt", "r") as f:
    input_txt = f.read().strip()

lines_per_elf = input_txt.split("\n\n")
calories_per_elf = [
    sum(map(int, lines.split("\n")))
    for lines in lines_per_elf
]
calories_per_elf.sort()

def p1():
    print(calories_per_elf[-1])

def p2():
    print(sum(calories_per_elf[-3:]))

p1()
p2()