# multi-agent-system-traffic-modelisation

The goal of this project is to simulate the traffic near ou school and see if we can optimise it


## Structure of the repo

Here is a quick view of the overall structure.
This is a starting point, don't hesitate to add new folder and add it in the Readme.

```
project/
├── go.mod
├── cmd/
│   ├── gui/main.go        # launch simulation + graphical window
│   └── headless/main.go   # (optionnal) simu without graphic : tests, stats, batch
├── internal/
│   ├── agent/             # interface Agent + type of agent (one file per type)
│   ├── env/               # state of the world, percept, application of act
│   ├── graph/             # Everything that is needed for the construction of a graph.
│   ├── simulation/        # loop, clock, start/pause/stop, snapshots
│   ├── gui/               # everything that is linked to Ebiten : Game, rendu, input, sprites
│   └── config/            # parameter (nb of agents, size of the map, speed)
└── assets/                # images, font
```

### Important Point

`internal/` : Go convention for code that will not be imported from outsite.

A `gui/` folder separated : `simu`, `agt` and `env` must never import Ebiten. Like this we can launch the simu without the graphical window, for test or even if we want to change the game engine.

Two main in `cmd/` : One version without graphical window, helps for debugging.