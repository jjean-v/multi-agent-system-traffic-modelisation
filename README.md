# multi-agent-system-traffic-modelisation

The goal of this project is to simulate the traffic near ou school and see if we can optimise it


# Structure of the repo

Here is a quick view of the overall structure.

```
project/
├── go.mod
├── cmd/
│   ├── gui/main.go        # launch simulation + graphical window
│   └── headless/main.go   # (optionnal) simu without graphic : tests, stats, batch
├── internal/
│   ├── agent/             # interface Agent + type of agent (one file per type)
│   ├── environment/       # state of the world, percept, application of act
│   ├── simulation/        # loop, clock, start/pause/stop, snapshots
│   ├── gui/               # everything that is linked to Ebiten : Game, rendu, input, sprites
│   └── config/            # parameter (nb of agents, size of the map, speed)
└── assets/                # images, font
```