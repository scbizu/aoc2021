package main

var input = `
.##.##...###.####.#.##..#.#.#.#.##.#.###....#.###....#..##.#####.#.#..##..##.##.#.#.########...###..
#.##..#..#.###...#..####.##.####.#..#...#..##...####...#..#.#.##...#.#.#.#..#...#.####..#..#..#....#
##...##.#.....#...##.###.#..#.#...#....##.##.###....#..##.#...###.###..#..#.#.#######..#...###...###
.##.###....#..#.#...#.###..##.##.####.###.#....##.####.#.#..###..##....#.#####.#.#.#....##...#.##..#
.##.#.#..##..#..#...#.#..#..##.#.########..####.##...#..#..##..#..##..##.......##.##..#..##..#..#.#.
#.#.###...###.##.##.#.#...#####.#.#......#.###..#..#.##..#.#.##.###...###..##.#..##.###..###.####.##
#.####.######....#####.......##....#.#...#..#.#..#.#.#####.#...###.#..###.......#.###.#.##.#.#...#..
....##.##.###.####..#.###.###.....####.##....#..###.###..#####...##.##.#.#..#..#..#.#.########.#.#..
##..#..#..#..####..#..#.#..###.#.####..###.##.###.##.###.###.#.##.##.#...#..#.###.###......#.####..#
.####.#..######.#.#...#..###.#..#.##...#####.#####.#####.#.....###.#.##.......#....####..#..#.#.#...
......#....#.#..##.####.....#.##..###....###..##...###.#....#####.#..###..#.#....#..###..#.#...#####
#.#.#########.##.#..##.#.#.##.#.##...###.#.....####...#..#..#####...#..#..#...##.#####.#.#....#..#.#
#..####.#.#......###.#....#.#.#####.#..#..#.####..#.#...#.#.#.###..###.###..###..##.#.#.#.##....#.##
#..#.#...###....##.#.##....##..#.#.#....#.##..########.###....###.#.##.#.####.###..#.#...#####.#.##.
.####.###.#####..#.###..#.#..######.#.###.....#..####..#.#..###.##.##.####.......#..#.#.####..##.#..
##.##.#.......#.....####.#.##.###..###......#.##..#.#.#..##.###.#.###.###...#.##.####.######..##.###
##..#......#..#..#.#.###..##..##...#...###....#...#..#.###...#.##.###..###....##..##..###.###.#...##
..#...#.###....#....#.#######..#.#..#..#.####...#.###.#.#..#....#.#.##.#.###.#...####.#.#....#..#.#.
..##........#.....#.###..#.####.#..####.#..#.....#####.##..#.#.#..##......#.#..##...#...##.##...#.##
##.#..##....##...##.#####.##........#.#..#..#..###.####....#...####...#..###.#####.##....###.#.####.
.#..#.#.#####..#.#.##.#.#.####.#..#...#..#....#....#..###########...##..####..##...#.##..#.#.##.....
..#..#.#.#.#........#...##.#..#..##.#.##.####...###....###.......#....##....#.###.##...###..#.###.#.
..#..#.##.###.###.##...#.####.###.###.##.#.####.##.#.##.###..#.##.##..#.###...#.##.#.##.#####...#...
...##..###.#####..#.#.###...##.....#.##..##..##.##.##..#...#..#####.#.####.#....####...#.###.##..###
###.####..##.#..#####.....###.####..##.#....#.#.#...#.#...###..###.#...#....##..##.#.#...#.##.####..
..#.#######..#.##.#...####...###..###.###..#..#######.###.##.##..#..#.#.####.##.##...######..#..##.#
###..#.###...####.#..#.##.#....#...#.#.#..#.####..#..#.##.#####..#.#......###..#..#..##.#.....##.#.#
.#.#####.##.#.#..#..###.###.##.#.##..#.......#####.###..#.#.###.#####.....####.#....######.#.#.###.#
#####...#..#....##......#.#..##..##.##.......#.........#.#......#....#.#.##...#..###.##.######.###..
#...#..#.####.##....#.#....##.#.#...#..#.#.#.#..####.#......#####....##.##.####....##....#####...#.#
.#.###....#.##.#...##.#.#.....#.#..#.#..##.#...#.....#.##.....###...##.#.#....##...#..##.#..##.#.#..
..#####.....#...##.#.....#.###.#..#....##.#..###.####...#.#.....###.##.#.#..####...#.###.#.#....#.##
#..####.#...###...##.##.#.#..#..###.#..#.#.#.##..#...#.##.##..#.###.#..#.#######..#.######..#.###...
...#..#.#.###.##.##.###.###.#...###.#.#.#..##.##.##.###....#.##.#.#.#.#.#.####.#####...#.#.##.##..#.
..##.#.###.#####..##.#..####..#..##..#.#..#.#..###.#.##.##..#.#..###..#..###...##...###...#####.####
###.#...##.#.###.##...#.####....#####.###..###.#...#.####..#.#.#.#.#..#.###.###......###...#..#.####
.###.....##..###.#.#####.#..##.#.##.#.#.#.....#.####..###.###....#.##..#...##...###......#....##..#.
..###.#.#####..##..#.###...#.##...##.#...####...#..#####..#..#.##.#.#....##.##..#..##...#.#..##.#..#
.###...........#..###.#.#....#.##.#.##..##.#...#.##.#......#.#.###..##..#.####.###........#.##.#...#
##.....#.#####.##.##..#.#.######..##...######..#...####.###..#.##.#.###..##.#########.....#.#.....##
###....#..###..##..###.....#.#......##.#.##.#.#..##.###.#..#####...#.#######..........####...#....#.
.###....#.#.#.#..##.#.#.#.##...#..#.###..#.......#.####...###.#.#.#....#.##########....######.#####.
..##...######..###..#...#..#..#####..###...#..######.#####.###.##.#.#..##....#.#...###..#......##.#.
##..##.#...#.#.#..###..##.##.#.##..###.###.#.......###..##.#######.##.##.#.####..#.#.#####.#...#.#..
#.#..##..##.##...#...#.######.#.#.#.#.....##...#.##.##.#..###...#...##....#....#.#..#.#.##...#.###..
....#.#.....#.###...#.#...##....##..#.#.#....##.###....#....#..#...#.#.#####.###.##.####.####.#.#...
.##...####.##.#....#.##..#....#...#.##...#.#......###.##.####..#.#######....#.#####..##.#.....##.#.#
....#.#....#...####..##..#.#.##......###.#...##..##..###..#.####..############....#...#..#......####
....#######..##.#.#...#..#.#..####....##..#.##.#.#...###...##..####...#..#.##.#.#####..###.####...#.
......###.#..#..###.##.#.#.#.#.#.##...##.##...##.#.#....##.#..#........#..###.#####....#####.#.####.
##...#########...##....#.#....#..#.#......#.##.#.###..#.##...#..######.#....#.###.#..#########.##...
#..#..#.##..#####..#.###.#.##..##.##..#.......##.#...###..##.#.###..###..##..##..#.#.###.##....#####
#.#.##..#.####.#.#.#..#......#.###.###..##...#.#.#.#...#.#..###.##...###.#...##.##.#.##...##.####.##
#..##.###...##.#....#.#####.####.###.###.#..##..#....##..#..###.......####..##..#..###.#.#.#.#...#..
#..#####...#.###..####.###.#..#...#..#...#..#.#...##.#..#..###.#....#.#......#..###..###.#.##.####..
##.##.#.####.#..##..#.###...###...#.....##....##...#.#..#..##.##.#...#..#.###.##..##..##.#.#.####.#.
.#.#.######.##...###.####.#...###.##.##.#.#.#..##.##.###.#..#.###..####.######.##..###.#....##.#..##
##.##.###.#...#####.#.#.....#.#.#...#.##.###.##...#.###...###......###...##...##.##.#....#......##..
...##.#.#..#..#.#..###......#...###.#.####.##...##...##.##.##.##....#.##...#.##.##.##......###..#...
#.#.....#.#..###...#.##..#..#.###..#######....#.###.#...#.##.#......##...#.###.##.#...#..#...#.#.#..
.##.##........##.#.#...#####.##....#.#...###.#.##....##.##...##..#..#.#..####.#..##..#...##.##..##.#
#.##.#.##.#.####....###.#..#..#.###....######.#.....####.#.#####..###..##.#.##.......###...##.#.###.
.#..###..#..#.####...#..#..##.####.#..#.###.##.#.....###......####....##...###.#...####...#.#####.##
....#.#...#.##.....#...#...#.###...#...###...##.##.###.##.#.####..#.##.#....####..#####..#.#..#..#.#
.###..##.#....##.###..#...##.##..###.###.#.##.##.####.####.######..#.#.#.###.##...####.#...##.#.#.#.
##..##.#......#.#.#...#..#.##.#.#.##.#..#.#..#.###.#..#...#..#...#...#.#..#....#.##..#######......#.
.......##..#..#.###.#..###..###.#####.###.##...##.###.#..###.##....#.....#..#.#.###.###.#.####..#...
#.#....###.##..##..#.#...##.#.#####.....#.#####.#.####.#.#....####..#..#.##.##....###...##.#.#.###.#
.....#..##.#..##.#####.#.###.#.##.#..##.##...##..##..#.#...##.#.####.####..#..##.##.#.#..##..##.....
.####.#.#.##..#....##..##.#...#....##.#.##.#.#.#.#...######.#..#.###.#.#.##....#...#..#.#.####..#..#
..###.###..#.##.#...###..##.##...##..##.###...####.###..#..#..#..###..##.#...#...#..##.#.#..........
.....#.#..##.#.#..####..#..####.#....#..#..#.##.#..##..#...####.#####.##.#...##.####.#...#.#.##.##..
###....##....#.###....###.#.###..###..##..#..#...#####.##.#..#..####.#.#..###..#.#...###.##....##.##
###...###.##...##.#####.#..##.#....##.###.#######..##.####.#..#..###.###..#.#..#.##.#..##.#######.#.
.####..#.#.###.#######.#.#...#.##.#...###.#.#..##.#.#....#.####..##...###..######.###..#####.##.###.
.#...##.#.####...##.#...#.#...##.....#.#..#..##..#.#.####..#.##.##...##.#####.#...#..##.#.#.##.#.#.#
.....##..##.####..##.#######..#..###.##.##..####.##..#.##...##..#..#....#...##.##.#..#...##.......#.
#..#######..#####..###.##..###....######..#..#.#.....#....##..#.#..#.#..#.##.####.##.#.###.#####.###
.##...#.##..###.#...##.#.#.###...#.#....##.####...#.###.###.###.##.###.##.###.#.####.....##..#.....#
##.##.#..#.......#..##....##..#####........#.###..##.#..##..#####.##.....###.#...#######.#...#.###.#
.####..#.#####......#..###.##.##.##.###.#.#.......##.#####...#.#.####..#.#..###...#.....#.####.###.#
##..#.##.#....##.#######.#.#.#.####...#.#.#.#.#.......######.#..###..###.#..##..#.##.#......#.##..##
.#.#.#.#...##.....##.#.#.........#.#.#.##.#.#...#.#.#.#.##....#..###.....###..###........###..#.##..
.##.......#.#.####.....#....#.#....##.........#.......###.#..###..####...#...#.##.......###...#..###
###...#####.#.###.......#.##.#####.###..###..##..##.#..####..####..###..###.#..#.###.##.#.###...#.#.
##.#.#.##.##..#....#..##.#.#.#..#...#.##.###..##.###.#..##.......##...#.###.#.###.###....#...#.####.
.###..##.........##.#.....###.##..#.#...#.#....###.####.###..###...#.#.##..#.#..#..#...#...#.....###
##.##.#...#..#..#....#........###.##.####.#.#..####..#..##.#...#...###...##.###.##.##.#########..#..
.##.###.#.........####..####....##...###..##.##..#.#######.##.##...#..#....#.##..#..#..###.###.#.###
#....#...##....###.####.#####.#.....##.#.##.#..#.....#...#.###...##.###.#.###..#.......####.##...#..
.....#..#.##....##.#..##.#.#.#.##.#.#.####..#.##..##.#.#.#.#.#.#.######..#....##.##...#..#...##..#..
....#.#.##...#.##..#...#.##...#.#.###.#####.#.##..###...##.##.#.#..##.##.#...#..#.....##..######.###
#....###.#...##.####.#..##.####.#...#.#....#.##...###........#..#..####.#.#.#....##.#.#.#.##..#.#..#
##.......#...##.####.###...#...#..##.....#..##..#......#.#..#.##..........#.#..#..#..####..##....#..
#....####.....#...#.#.#####.#...##.###......#.##...##.##..#.#....##....#.....##..###.......#..##.#.#
....#..#...#..##.###.##.##..#.#.##..########..###.##.#..#.#.###.....#..###..###....#..##.#.....#.##.
#..#.#..#.###..#.####.....#...#...##...#....##..##...###.#.#########..##..##.#.#.#.##.#...#.#..#...#
..##.####.#..##...##.#.###.###..##.#.##..##....#.#.###.#.###.###..#.###..##.....##......####.##.....
.###.###..#....##.#####.##.#.#.####.#.##..#.###..#...#..#.###.####.#.#...##..###.##.#####...##..##.#
...####.##.###.#.###...####.....#.#.###..#.##.###..###.#.###..#.##....#.####.#..#.#..###...#..#####.
`

var checksum = `#..##..#.#.#.###.#..##.####..##..##.##.#.###.......##..#..#.#.#.#..#...##....#.#.##.###....###.#.##.##..##.#...##.##...#...##......##.#...#.......##.#.#..####.##..#.#.#.....##.....#....#.#.#.##..##..##.##.....###...#.#..###.#######.#.....###....#......#.###.#...#.#####.#.#.###..#...##.##.#..#...######.###.#.##...####..####.###...####........##...##.##.####...##.#.#...##.#####.#....#.....##..#......###..###.#.#..###..#.####......#.....#.#.#.###..#.#..#..#...##..##..#.##....#....#.##..###..#....##.##..#.###..`
