from datetime import datetime
import random

random.seed()

max = 100000

loops = 100


def small_comp(loops: int) -> None:

    list_of_lists = [
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
    ]

    choices = list(range(0, 8))

    list1 = random.choice(choices)

    choices.remove(list1)

    list2 = random.choice(choices)
    start_time = datetime.now()
    matches = []
    for loop in range(0, loops):
        for i in range(0, max):
            if list_of_lists[list1][i] == list_of_lists[list2][i]:
                matches.append(i)
            else:
                continue

    spent_time = datetime.now() - start_time
    print(spent_time)
    print(len(matches))


def large_comp(loops: int) -> None:
    list_of_lists = [
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
        [random.randint(0, 10) for num in range(0, max)],
    ]

    start_time = datetime.now()
    matches = []
    for loop in range(0, loops):
        for i in range(0, max):
            if (
                (list_of_lists[0][i] == list_of_lists[1][i])
                and (list_of_lists[1][i] == list_of_lists[2][i])
                and (list_of_lists[2][i] == list_of_lists[3][i])
                and (list_of_lists[3][i] == list_of_lists[4][i])
                and (list_of_lists[4][i] == list_of_lists[5][i])
                and (list_of_lists[5][i] == list_of_lists[6][i])
                and (list_of_lists[6][i] == list_of_lists[7][i])
                and (list_of_lists[7][i] == list_of_lists[8][i])
            ):
                matches.append(i)
            else:
                continue

    spent_time = datetime.now() - start_time
    print(spent_time)
    print(len(matches))


if __name__ == "__main__":
    print("Small comp")
    small_comp(loops)
    print("Large comp")
    large_comp(loops)
