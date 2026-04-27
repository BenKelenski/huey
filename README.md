# huey

A terminal UI for controlling Philips Hue lights, built with [Bubbletea](https://charm.land/bubbletea).

![demo](demo.gif)

## Features

- Browse all rooms on your Hue bridge
- Toggle lights on or off per room
- Set room color from a palette of presets (Red, Orange, Yellow, Green, Cyan, Blue, Purple, Pink, Warm White, Cool White)

## Prerequisites

- Go 1.21+
- A Philips Hue bridge on your local network
- A Hue API application key — follow the [Hue developer getting started guide](https://developers.meethue.com/develop/get-started-2/) to create one by pressing the bridge link button and making a POST request to `http://<bridge-ip>/api`

## Configuration

Set the following environment variables before running:

| Variable          | Description                                      |
|-------------------|--------------------------------------------------|
| `HUE_IP_ADDRESS`  | Local IP address of your Hue bridge              |
| `HUE_USERNAME`    | Your Hue application key (also called username)  |

```sh
export HUE_IP_ADDRESS=192.168.1.x
export HUE_USERNAME=your-api-key-here
```

## Installation

```sh
git clone https://github.com/BenKelenski/huey
cd huey
go build -o huey .
```

## Usage

```sh
./huey
```

### Keybindings

| Key              | Action                  |
|------------------|-------------------------|
| `↑` / `k`        | Move cursor up          |
| `↓` / `j`        | Move cursor down        |
| `enter` / `space`| Select                  |
| `esc` / `backspace` | Go back              |
| `q`              | Quit / go back          |
| `ctrl+c`         | Quit                    |
