# mars-rover-simulator
  
## Overview

Mars Rover Simulator is a Go application that simulates the movement and control of a rover on a grid representing the surface of Mars. The simulator processes commands to navigate the rover, avoid obstacles, and report its final position and orientation.

## Features

- Command-line interface for inputting rover instructions
- Configurable grid size and initial rover position
- Detailed output of rover movements and final state

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Celso-Alves/mars-rover-simulator.git
   ```
2. Navigate to the project directory:
   ```bash
   cd mars-rover-simulator
   ```
3. Build the project:
   ```bash
   make build
   ```

## Usage

Run the simulator with a configuration file or interactively:

```bash
make run
```

Example input:
```
Grid: 5 5
Rover 1: 1 2 N
Commands: LMLMLMLMM
Rover 2: 3 3 E
Commands: MMRMMRMRRM
```

## Example Output

```
Rover 1 Final Position: 1 3 N
Rover 2 Final Position: 5 1 E
```

## Configuration

- **Grid:** Specify the grid size (e.g., `5 5`)
- **Rover:** Set initial position and orientation (`x y direction`)
- **Commands:** Sequence of instructions (`L`, `R`, `M`)

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
