package main

import (
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	x       int // register X
	cycle   int // #cycles past
	counter int // set to simulate cpu fetch cycles
	buf     int // x register buffer
}

func (c *cpu) exec(args ...string) (step int) {
	if c.counter == 0 {
		c.x = c.buf
		c.buf = c.x
	}
	fmt.Println("executing: ", args)
	switch args[0] {
	case "noop":
		c.cycle++
		return 1
	case "addx":

		if c.counter == 0 {
			c.counter = 1
			c.cycle++
			return 0
		}

		if c.counter == 1 {
			val, _ := strconv.Atoi(args[1])
			c.buf = c.x + val
			c.cycle++
			c.counter = 0
			return 1
		}
	}

	panic("bad instruction")
}

func (c *cpu) ssi() int {
	fmt.Println(c.x, c.cycle)
	return c.x * c.cycle
}

type gpu struct {
	drawn map[int]int
}

// part 2
func (g *gpu) tick(c *cpu) {
	s2 := c.x + 1
	s1 := c.x
	s3 := c.x + 2
	fmt.Println("sp. pos:", s1, s2, s3, "cycle:", c.cycle)
	fmt.Println("X val:", c.x)

	if c.cycle%40 == s1 || c.cycle%40 == s2 || c.cycle%40 == s3 {
		fmt.Println("drawing pixel", c.cycle)
		g.drawn[c.cycle] = 1
	}
}

func (g *gpu) render() {
	fmt.Println("drawn:", g.drawn)
	w, h := 40, 6
	for i := 1; i <= h*w; i++ {
		if i%w == 1 {
			fmt.Println("")
		}

		if g.drawn[i] == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

	}

}

func main() {
	c := new(cpu)
	c.x = 1
	c.buf = 1
	sum := 0

	g := new(gpu)
	g.drawn = make(map[int]int, 0)

	cmds := strings.Split(rinput, "\n")
	for step := 0; step < len(cmds); {
		cmd := strings.Split(cmds[step], " ")
		step += c.exec(cmd...)

		if c.cycle == 20 || c.cycle == 60 ||
			c.cycle == 100 || c.cycle == 140 ||
			c.cycle == 180 || c.cycle == 220 {
			fmt.Println("cycle: ", c.cycle, "ssi:", c.ssi())
			sum += c.ssi()
		}
		g.tick(c)
	}
	fmt.Println(sum)
	g.render()
}

const input = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

const rinput = `noop
noop
noop
addx 6
addx -1
addx 5
noop
noop
noop
addx 5
addx 11
addx -10
addx 4
noop
addx 5
noop
noop
noop
addx 1
noop
addx 4
addx 5
noop
noop
noop
addx -35
addx -2
addx 5
addx 2
addx 3
addx -2
addx 2
addx 5
addx 2
addx 3
addx -2
addx 2
addx 5
addx 2
addx 3
addx -28
addx 28
addx 5
addx 2
addx -9
addx 10
addx -38
noop
addx 3
addx 2
addx 7
noop
noop
addx -9
addx 10
addx 4
addx 2
addx 3
noop
noop
addx -2
addx 7
noop
noop
noop
addx 3
addx 5
addx 2
noop
noop
noop
addx -35
noop
noop
noop
addx 5
addx 2
noop
addx 3
noop
noop
noop
addx 5
addx 3
addx -2
addx 2
addx 5
addx 2
addx -25
noop
addx 30
noop
addx 1
noop
addx 2
noop
addx 3
addx -38
noop
addx 7
addx -2
addx 5
addx 2
addx -8
addx 13
addx -2
noop
addx 3
addx 2
addx 5
addx 2
addx -15
noop
addx 20
addx 3
noop
addx 2
addx -4
addx 5
addx -38
addx 8
noop
noop
noop
noop
noop
noop
addx 2
addx 17
addx -10
addx 3
noop
addx 2
addx 1
addx -16
addx 19
addx 2
noop
addx 2
addx 5
addx 2
noop
noop
noop
noop
noop
noop`
