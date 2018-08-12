import random
import influunt

randint = influunt.add_operation("random/randint", random.randint)

with influunt.Graph() as graph:
    start = influunt.placeholder()
    end = influunt.placeholder()

    randomNumber = randint(start, end)

    (r,) = graph.executor().run({start: 1, end: 100}, [randomNumber])
 
    print(r)
    assert r >= 1 and r <= 100
