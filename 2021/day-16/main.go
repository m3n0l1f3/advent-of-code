package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var HexDecodeMap = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type Packet struct {
	version      int
	typeID       int
	lengthTypeID *int
	literal      *[]string // typeID == 4
	subpackets   *[]Packet // typeID != 4
}

func min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func max(left int, right int) int {
	if left < right {
		return right
	}
	return left
}

func loadData(scanner bufio.Scanner) string {

	if !scanner.Scan() {
		panic("Scanner.Scan() returns false")
	}

	lineText := scanner.Text()
	fullText := ""
	for _, r := range lineText {
		fullText += HexDecodeMap[r]
	}

	return fullText
}

func readBits(text string, bits int) (int, string) {
	bitText, remainText := text[:bits], text[bits:]
	bitInt, _ := strconv.ParseInt(bitText, 2, 0)
	return int(bitInt), remainText
}

func getHeader(text string) (version int, typeID int, remainText string) {
	version, newText := readBits(text, 3)
	typeID, remainText = readBits(newText, 3)
	return
}

func parseLiteralPacket(text string) ([]string, string) {

	pointer := 0
	literals := make([]string, 0)
	for pointer < len(text) {
		literals = append(literals, text[pointer:pointer+5])

		pointer += 5
		if text[pointer-5] == '0' {
			break
		}
	}

	return literals, text[pointer:]
}

func parseOperationPacket(text string) ([]Packet, string) {
	bit, contentText := readBits(text, 1)
	packetList := make([]Packet, 0)
	subpacketText := ""
	var packetLength int
	var packetCount int

	// Length type ID
	// - If is `0`, next 15 bits represents total length in bits
	// - If ii `1`, next 11 bits represents number of sub-packets
	if bit == 0 {
		packetLength, subpacketText = readBits(contentText, 15)
		for i := 0; i < packetLength; {
			packet, newText := parse(subpacketText)
			packetList = append(packetList, packet)
			i += len(subpacketText) - len(newText)
			subpacketText = newText
			if len(subpacketText) == 0 {
				break
			}
		}
	} else {
		packetCount, subpacketText = readBits(contentText, 11)
		for i := 0; i < packetCount; i++ {
			packet, newText := parse(subpacketText)
			packetList = append(packetList, packet)
			subpacketText = newText
		}
	}
	return packetList, subpacketText
}

func parse(text string) (Packet, string) {
	versionNumber, typeID, remainText := getHeader(text)
	packet := Packet{versionNumber, typeID, nil, nil, nil}
	var literals []string
	var subPackets []Packet
	if packet.typeID == 4 {
		literals, remainText = parseLiteralPacket(remainText)
		packet.literal = &literals
	} else {
		subPackets, remainText = parseOperationPacket(remainText)
		packet.subpackets = &subPackets
	}

	return packet, remainText
}

func countVersionNumber(packet Packet) int {
	result := packet.version
	if packet.subpackets != nil {
		for _, p := range *packet.subpackets {
			result += countVersionNumber(p)
		}
	}
	return result
}

func partOne(text string) int {
	packet, _ := parse(text)
	result := countVersionNumber(packet)

	return result
}

func computePacket(packet Packet) int {
	result := 0
	switch packet.typeID {
	case 0:
		{
			for _, p := range *packet.subpackets {
				result += computePacket(p)
			}
		}
	case 1:
		{
			result = 1
			for _, p := range *packet.subpackets {
				result *= computePacket(p)
			}
		}
	case 2:
		{
			currentMin := math.MaxInt
			for _, p := range *packet.subpackets {
				currentMin = min(currentMin, computePacket(p))
			}
			result = currentMin
		}
	case 3:
		{
			currentMax := math.MinInt
			for _, p := range *packet.subpackets {
				currentMax = max(currentMax, computePacket(p))
			}
			result = currentMax
		}
	case 4:
		{

			binaryString := ""
			for _, l := range *packet.literal {
				binaryString += l[1:]
			}
			result64, _ := strconv.ParseInt(binaryString, 2, 64)
			result = int(result64)
		}
	case 5:
		{
			packetOne := computePacket((*packet.subpackets)[0])
			packetTwo := computePacket((*packet.subpackets)[1])
			if packetOne > packetTwo {
				result = 1
			}
		}
	case 6:
		{
			packetOne := computePacket((*packet.subpackets)[0])
			packetTwo := computePacket((*packet.subpackets)[1])
			if packetOne < packetTwo {
				result = 1
			}
		}
	case 7:
		packetOne := computePacket((*packet.subpackets)[0])
		packetTwo := computePacket((*packet.subpackets)[1])
		if packetOne == packetTwo {
			result = 1
		}
	}
	return result
}

func partTwo(text string) int {
	packet, _ := parse(text)
	result := computePacket(packet)

	return result
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Missing Input File Path")
		os.Exit(1)
	}

	// Open File
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fullText := loadData(*scanner)
	result := partTwo(fullText)

	fmt.Println(result)
}
