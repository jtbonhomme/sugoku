# sugoku

Sodoku puzzle solver with batcktrack algorithm written in Golang.

## Backtracking 

### Solver

[Wikipedia article about backtracking algorithms](https://en.wikipedia.org/wiki/Backtracking)

![](https://github.com/jtbonhomme/sugoku/blob/master/Sudoku_solved_by_bactracking.gif)

![](https://github.com/jtbonhomme/sugoku/blob/master/Sudoku_solved_by_bactracking_with_candidates.gif)

```sh
go run cmd/solver/main.go -f _examples/not-unique-solution.json -s 3
```

### Generator

Same code than the solver, except than we start from an emprty grid and without graphical display.

```sh
go run cmd/headless/main.go
```

## Sources

* https://en.wikipedia.org/wiki/Backtracking
* https://openclassrooms.com/courses/le-backtracking-par-l-exemple-resoudre-un-sudoku

## Todo

* [ ] Interactive play mode