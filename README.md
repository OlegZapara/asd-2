# 8 Queen problem

The eight-queens challenge consists of placing eight queens on the board in such a way that none of them are in a position to attack each other. That is, they should not stand in the same vertical, horizontal or diagonal.

# How to run

## LDFS implementation:

`go run main.go -mode=dfs`

## RBFS implementation:

`go run main.go -mode=bfs`

> Note: After running the program press ENTER to start execution

---

To reduce or increase algorithm execution time you can change delay(ms) between each iteration:

`go run main.go -mode=dfs -delay=10`

Default value for delay is 5 milliseconds
