# LGO57: Optimal Warehouse Placement Example

This is an example of a non-linear optimization problem. Given C customers at
different locations with individual demands (expressed as weights), find an
optimal location to place a warehouse to minimize the weight carried during
transport.

## Basic program

A map with a configurable number of "customers" placed at random locations
(with some minimal distance, to be legible); with a fixed, but configurable
seed for reproducability.

To get a better visual impression, we render multiple files (internally, we may
use SVG or something else); svg or png per flag:

* a map with just the customers (image)
* a map with the customers and the optimal location (image)
* a gif that shows the optimization process and the locations found during the process (gif)

## Algorithm

We use a basic implementation of nelder-mead, downhill simplex.

## Misc

* A single Go file, for simplicity
* Simple, clear, appealing and comprehensible visuals
