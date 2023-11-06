from networking_tabels.utils import get_user_input, init_people_scores, run

# Getting user input
people_arr, tables_capacity, combinations_length = get_user_input()

# Initializing people scores
people_scores = init_people_scores(people_arr)

# Running the algorithm
run(
    combinations_length=combinations_length,
    people_arr=people_arr,
    people_scores=people_scores,
    tables_capacity=tables_capacity,
)
