import copy
import json
import random

from tqdm import tqdm


def get_user_input():
    people_arr: list[str] = input("Enter the names of the people: ").split()

    tables_capacity: list[int] = list(
        map(int, input("Enter the capacity of the each table (for example 4 5 5 6): "))
    )

    combinations_length = int(input("Enter the number of combinations to compute: "))

    return people_arr, tables_capacity, combinations_length


def init_people_scores(people_arr) -> dict:
    """Creates a dictionary of people
    with each person having a dictionary of other people
    with a value of 0"""

    people_scores = {}
    for person in people_arr:
        people_scores[person] = {}
        for other_person in people_arr:
            if person != other_person:
                people_scores[person][other_person] = 0

    return people_scores


def best_combination(
    combinations_length, people_arr, people_scores, tables_capacity
) -> dict:
    """Creates a list of combinations of tables"""

    _best_combination = {}

    for i in tqdm(range(combinations_length), desc="Computing...", unit=" combination"):
        tables = []
        people_arr_copy = copy.deepcopy(people_arr)
        people_scores_copy = copy.deepcopy(people_scores)

        for table_length in tables_capacity:
            table = []
            table_score = 0

            for i in range(table_length):
                table.append(
                    people_arr_copy.pop(random.randint(0, len(people_arr_copy) - 1))
                )

            for person in table:
                for other_person in table:
                    if person != other_person:
                        if people_scores_copy[person][other_person] > 0:
                            people_scores_copy[person][other_person] += 10
                        else:
                            people_scores_copy[person][other_person] += 1
                        table_score += people_scores_copy[person][other_person]

            tables.append(
                {
                    "table": table,
                    "score": table_score,
                }
            )
        combination_score = sum(table["score"] for table in tables)

        if combination_score < _best_combination.get("score", 9999999999999999999):
            _best_combination = {
                "tables": tables,
                "score": sum(table["score"] for table in tables),
                "people_scores": people_scores_copy,
            }

    return _best_combination


def run(combinations_length, people_arr, people_scores, tables_capacity):
    """Runs the algorithm"""

    iterations = len(tables_capacity)
    final_score = 0

    for i in range(iterations):
        print("Iteration", i + 1)
        combination = best_combination(
            combinations_length=combinations_length,
            people_arr=people_arr,
            people_scores=people_scores,
            tables_capacity=tables_capacity,
        )
        final_score += combination["score"]
        people_scores = combination["people_scores"]
        print(f"Score: {combination['score']}")
        print(f"Tables: {json.dumps(combination['tables'], indent=4, sort_keys=True)}")

    print(f"Final Score: {final_score}")
