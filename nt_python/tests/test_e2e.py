from networking_tabels.utils import init_people_scores, run


def test_integration():
    people_arr = [
        chr(i) for i in range(ord("a"), ord("z") + 1)
    ]  # this will be replaced by dynamically entered by the user

    tables_capacity = [
        4,
        4,
        4,
        4,
        4,
        3,
        3,
    ]  # this will be replaced by dynamically entered by the user
    combinations_length = 10000  # will dynamically entered by the user

    # Initializing people scores
    people_scores = init_people_scores(people_arr)

    # Running the algorithm
    run(
        combinations_length=combinations_length,
        people_arr=people_arr,
        people_scores=people_scores,
        tables_capacity=tables_capacity,
    )


if __name__ == "__main__":
    test_integration()
