TinyViz
=======
TinyViz is CLI graph visualization tool that was intended to help me with assignments for algorithms course.

## What can it do?
* Draw graphs described in JSON configuration files. It supports directed/undirected, weighted/unweighted (both vertice and edge weights), colored/uncolored (both vertice and edge coloring) graphs.
* Save output images in JPG, PNG formats.
* Hot reload output image on single config file.

## Usage
By default `tinyviz` walks on current directory and visualize all graphs described in files with .descr extension. If you pass file name, then graph will be visualized from specific configuration file. Tool has several command-line flags, you can find out more by typing `tinyviz --help` in your terminal.

## Example
![Hot reloading](demo.gif?raw=true)