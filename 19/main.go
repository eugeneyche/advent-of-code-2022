package main

import (
	"fmt"
	"os"
)

type blueprint struct {
	oreRobotOreCost,
	clayRobotOreCost,
	obsidianRobotOreCost,
	obsidianRobotClayCost,
	geodeRobotOreCost,
	geodeRobotObsidianCost int
}

func intMax(l ...int) int {
	m := l[0]
	for _, i := range l {
		if m < i {
			m = i
		}
	}
	return m
}

func intMin(l ...int) int {
	m := l[0]
	for _, i := range l {
		if m > i {
			m = i
		}
	}
	return m
}

func intDivCeil(a, b int) int {
	if a <= 0 {
		return 0
	}
	return (a-1)/b + 1
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	blueprints := []blueprint{}

	var blueprintIndex,
		oreRobotOreCost,
		clayRobotOreCost,
		obsidianRobotOreCost,
		obsidianRobotClayCost,
		geodeRobotOreCost,
		geodeRobotObsidianCost int
	for true {
		if _, err := fmt.Fscanf(
			file,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n",
			&blueprintIndex,
			&oreRobotOreCost,
			&clayRobotOreCost,
			&obsidianRobotOreCost,
			&obsidianRobotClayCost,
			&geodeRobotOreCost,
			&geodeRobotObsidianCost,
		); err != nil {
			break
		}

		blueprints = append(blueprints, blueprint{
			oreRobotOreCost,
			clayRobotOreCost,
			obsidianRobotOreCost,
			obsidianRobotClayCost,
			geodeRobotOreCost,
			geodeRobotObsidianCost,
		})
	}

	getMaxGeodes := func(bl blueprint, t int) int {
		maxNeededOre := intMax(
			bl.oreRobotOreCost,
			bl.clayRobotOreCost,
			bl.obsidianRobotOreCost,
			bl.geodeRobotOreCost,
		)

		type memKey struct {
			s state
			t int
		}

		var getMaxGeodesImpl func(s state, t int) int
		getMaxGeodesImpl = func(s state, t int) int {
			maxGeodes := s.step(t).geodes
			needOre := maxNeededOre > s.oreRobot
			needObsidian := bl.geodeRobotObsidianCost > s.obsidianRobot
			needClay := bl.obsidianRobotClayCost > s.clayRobot && needObsidian
			if !needOre && !needObsidian {
				// Keep building geode robots until we run out of time
				return maxGeodes + t*(t-1)/2
			}
			if dt := s.timeToWaitForOreRobot(bl); needOre && 0 < dt && dt <= t {
				nextS := s.step(dt).buildOreRobot(bl)
				nextT := t - dt
				maxGeodes = intMax(getMaxGeodesImpl(nextS, nextT), maxGeodes)
			}
			if dt := s.timeToWaitForClayRobot(bl); needClay && 0 < dt && dt <= t {
				nextS := s.step(dt).buildClayRobot(bl)
				nextT := t - dt
				maxGeodes = intMax(getMaxGeodesImpl(nextS, nextT), maxGeodes)
			}
			if dt := s.timeToWaitForObsidianRobot(bl); needObsidian && 0 < dt && dt <= t {
				nextS := s.step(dt).buildObsidianRobot(bl)
				nextT := t - dt
				maxGeodes = intMax(getMaxGeodesImpl(nextS, nextT), maxGeodes)
			}
			if dt := s.timeToWaitForGeodeRobot(bl); 0 < dt && dt <= t {
				nextS := s.step(dt).buildGeodeRobot(bl)
				nextT := t - dt
				maxGeodes = intMax(getMaxGeodesImpl(nextS, nextT), maxGeodes)
			}
			return maxGeodes
		}
		return getMaxGeodesImpl(NewInitialState(), t)
	}

	p1 := func() {
		total := 0
		for idx, bl := range blueprints {
			geodes := getMaxGeodes(bl, 24)
			total += geodes * (idx + 1)
		}
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		total := 1
		for _, bl := range blueprints[:intMin(len(blueprints), 3)] {
			geodes := getMaxGeodes(bl, 32)
			total *= geodes
		}
		fmt.Printf("%d\n", total)
	}

	p1()
	p2()
}

type state struct {
	ore,
	clay,
	obsidian,
	geodes,
	oreRobot,
	clayRobot,
	obsidianRobot,
	geodeRobot int
}

func NewInitialState() state {
	return state{0, 0, 0, 0, 1, 0, 0, 0}
}

func (s state) timeToWaitForOreRobot(bl blueprint) int {
	if s.oreRobot == 0 {
		return -1
	}
	return intDivCeil(bl.oreRobotOreCost-s.ore, s.oreRobot) + 1
}

func (s state) timeToWaitForClayRobot(bl blueprint) int {
	if s.oreRobot == 0 {
		return -1
	}
	return intDivCeil(bl.clayRobotOreCost-s.ore, s.oreRobot) + 1
}

func (s state) timeToWaitForObsidianRobot(bl blueprint) int {
	if s.oreRobot == 0 || s.clayRobot == 0 {
		return -1
	}
	return intMax(
		intDivCeil(bl.obsidianRobotOreCost-s.ore, s.oreRobot),
		intDivCeil(bl.obsidianRobotClayCost-s.clay, s.clayRobot),
	) + 1
}

func (s state) timeToWaitForGeodeRobot(bl blueprint) int {
	if s.oreRobot == 0 || s.obsidianRobot == 0 {
		return -1
	}
	return intMax(
		intDivCeil(bl.geodeRobotOreCost-s.ore, s.oreRobot),
		intDivCeil(bl.geodeRobotObsidianCost-s.obsidian, s.obsidianRobot),
	) + 1
}

func (s state) buildOreRobot(bl blueprint) state {
	return state{
		s.ore - bl.oreRobotOreCost,
		s.clay,
		s.obsidian,
		s.geodes,
		s.oreRobot + 1,
		s.clayRobot,
		s.obsidianRobot,
		s.geodeRobot,
	}
}

func (s state) buildClayRobot(bl blueprint) state {
	return state{
		s.ore - bl.clayRobotOreCost,
		s.clay,
		s.obsidian,
		s.geodes,
		s.oreRobot,
		s.clayRobot + 1,
		s.obsidianRobot,
		s.geodeRobot,
	}
}

func (s state) buildObsidianRobot(bl blueprint) state {
	return state{
		s.ore - bl.obsidianRobotOreCost,
		s.clay - bl.obsidianRobotClayCost,
		s.obsidian,
		s.geodes,
		s.oreRobot,
		s.clayRobot,
		s.obsidianRobot + 1,
		s.geodeRobot,
	}
}

func (s state) buildGeodeRobot(bl blueprint) state {
	return state{
		s.ore - bl.geodeRobotOreCost,
		s.clay,
		s.obsidian - bl.geodeRobotObsidianCost,
		s.geodes,
		s.oreRobot,
		s.clayRobot,
		s.obsidianRobot,
		s.geodeRobot + 1,
	}
}

func (s state) step(n int) state {
	return state{
		s.ore + n*s.oreRobot,
		s.clay + n*s.clayRobot,
		s.obsidian + n*s.obsidianRobot,
		s.geodes + n*s.geodeRobot,
		s.oreRobot,
		s.clayRobot,
		s.obsidianRobot,
		s.geodeRobot,
	}
}
