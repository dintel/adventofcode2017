package main

import (
	"fmt"
	"log"
)

const DiskSize = 128

func BitsOn(ch byte) int {
	switch ch {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 1
	case '3':
		return 2
	case '4':
		return 1
	case '5':
		return 2
	case '6':
		return 2
	case '7':
		return 3
	case '8':
		return 1
	case '9':
		return 2
	case 'a':
		return 2
	case 'b':
		return 3
	case 'c':
		return 2
	case 'd':
		return 3
	case 'e':
		return 3
	case 'f':
		return 4
	}
	log.Fatalf("Unknown byte '%s' while counting bits", ch)
	return 0
}

func Parse(ch byte) int {
	switch ch {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'a':
		return 10
	case 'b':
		return 11
	case 'c':
		return 12
	case 'd':
		return 13
	case 'e':
		return 14
	case 'f':
		return 15
	default:
		log.Fatalf("Unknown byte '%s' while counting bits", ch)
	}
	return 0
}

type Disk struct {
	sectors [][]int
	size    int
}

func NewDisk(size int) *Disk {
	sectors := make([][]int, size)
	for row := range sectors {
		sectors[row] = make([]int, size)
	}
	return &Disk{
		size:    size,
		sectors: sectors,
	}
}

func (d *Disk) Set(x, y, value int) {
	d.sectors[x][y] = value
}

func (d *Disk) Get(x, y int) int {
	return d.sectors[x][y]
}

func (d *Disk) Find(value int) (int, int) {
	for x := 0; x < d.size; x++ {
		for y := 0; y < d.size; y++ {
			if d.sectors[x][y] == value {
				return x, y
			}
		}
	}
	return -1, -1
}

func (d *Disk) MarkGroup(x, y, group int) {
	queue := make([]struct {
		x int
		y int
	}, 1)
	queue[0].x = x
	queue[0].y = y
	var current struct {
		x int
		y int
	}
	for len(queue) != 0 {
		current, queue = queue[0], queue[1:]
		d.sectors[current.x][current.y] = group
		if current.x-1 >= 0 && d.sectors[current.x-1][current.y] == 1 {
			queue = append(queue, struct {
				x int
				y int
			}{x: current.x - 1, y: current.y})
		}
		if current.y-1 >= 0 && d.sectors[current.x][current.y-1] == 1 {
			queue = append(queue, struct {
				x int
				y int
			}{x: current.x, y: current.y - 1})
		}
		if current.x+1 < d.size && d.sectors[current.x+1][current.y] == 1 {
			queue = append(queue, struct {
				x int
				y int
			}{x: current.x + 1, y: current.y})
		}
		if current.y+1 < d.size && d.sectors[current.x][current.y+1] == 1 {
			queue = append(queue, struct {
				x int
				y int
			}{x: current.x, y: current.y + 1})
		}
	}
}

func (d *Disk) String() string {
	result := ""
	for x := 0; x < d.size; x++ {
		for y := 0; y < d.size; y++ {
			if d.sectors[x][y] == 0 {
				result += "."
			} else if d.sectors[x][y] == 1 {
				result += "#"
			} else {
				result += fmt.Sprintf("%d", d.sectors[x][y]%10)
			}
		}
		result += "\n"
	}
	return result
}

func Day14(part int, data []byte) {
	log.Printf("Loaded key %s", string(data))
	switch part {
	case 1:
		used := 0
		for i := 0; i < DiskSize; i++ {
			hasher := NewKnotHasher()
			rowInput := []byte(fmt.Sprintf("%s-%d", data, i))
			hasher.FullHash(rowInput)
			for _, ch := range hasher.String() {
				used += BitsOn(byte(ch))
			}
		}
		log.Printf("There are %d used blocks", used)
	case 2:
		disk := NewDisk(DiskSize)
		for i := 0; i < DiskSize; i++ {
			hasher := NewKnotHasher()
			rowInput := []byte(fmt.Sprintf("%s-%d", data, i))
			hasher.FullHash(rowInput)
			for j, ch := range []byte(hasher.String()) {
				digit := Parse(byte(ch))
				for k := 0; k < 4; k++ {
					bit := (digit >> byte(k)) & 1
					disk.Set(i, j*4+(3-k), bit)
				}
			}
		}
		regions := 1
		x, y := disk.Find(1)
		for x != -1 && y != -1 {
			disk.MarkGroup(x, y, regions+1)
			regions++
			x, y = disk.Find(1)
		}
		regions--
		log.Printf("There are %d regions", regions)
	}
}
