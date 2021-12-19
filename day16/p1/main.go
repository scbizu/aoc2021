package main

import (
	"fmt"
	"os"
)

const (
	codes = "E058F79802FA00A4C1C496E5C738D860094BDF5F3ED004277DD87BB36C8EA800BDC3891D4AFA212012B64FE21801AB80021712E3CC771006A3E47B8811E4C01900043A1D41686E200DC4B8DB06C001098411C22B30085B2D6B743A6277CF719B28C9EA11AEABB6D200C9E6C6F801F493C7FE13278FFC26467C869BC802839E489C19934D935C984B88460085002F931F7D978740668A8C0139279C00D40401E8D1082318002111CE0F460500BE462F3350CD20AF339A7BB4599DA7B755B9E6B6007D25E87F3D2977543F00016A2DCB029009193D6842A754015CCAF652D6609D2F1EE27B28200C0A4B1DFCC9AC0109F82C4FC17880485E00D4C0010F8D110E118803F0DA1845A932B82E200D41E94AD7977699FED38C0169DD53B986BEE7E00A49A2CE554A73D5A6ED2F64B4804419508B00584019877142180803715224C613009E795E58FA45EA7C04C012D004E7E3FE64C27E3FE64C24FA5D331CFB024E0064DEEB49D0CC401A2004363AC6C8344008641B8351B08010882917E3D1801D2C7CA0124AE32DD3DDE86CF52BBFAAC2420099AC01496269FD65FA583A5A9ECD781A20094CE10A73F5F4EB450200D326D270021A9F8A349F7F897E85A4020CF802F238AEAA8D22D1397BF27A97FD220898600C4926CBAFCD1180087738FD353ECB7FDE94A6FBCAA0C3794875708032D8A1A0084AE378B994AE378B9A8007CD370A6F36C17C9BFCAEF18A73B2028C0A004CBC7D695773FAF1006E52539D2CFD800D24B577E1398C259802D3D23AB00540010A8611260D0002130D23645D3004A6791F22D802931FA4E46B31FA4E4686004A8014805AE0801AC050C38010600580109EC03CC200DD40031F100B166005200898A00690061860072801CE007B001573B5493004248EA553E462EC401A64EE2F6C7E23740094C952AFF031401A95A7192475CACF5E3F988E29627600E724DBA14CBE710C2C4E72302C91D12B0063F2BBFFC6A586A763B89C4DC9A0"
)

func main() {
	packets := hexToBinary([]byte(codes))
	fmt.Fprintf(os.Stdout, "Packets: %v\n", packets)
	versions, _ := decodePacket(packets)
	fmt.Fprintf(os.Stdout, "version: %d\n", versions)
}

func decodePacket(packets []byte) (int, []byte) {
	fmt.Fprintf(os.Stdout, "[I]packets: %v\n", packets)
	if len(packets) < 3 {
		return 0, packets
	}
	var versions int
	versions += binaryToDecimal(packets[:3])
	fmt.Fprintf(os.Stdout, "version: %d\n", versions)
	if len(packets) < 3 {
		return versions, packets
	}
	packets = packets[3:]
	switch binaryToDecimal(packets[:3]) {
	case 4:
		fmt.Fprintln(os.Stdout, "literals")
		if len(packets) < 3 {
			return versions, packets
		}
		// remove type id
		packets = packets[3:]
		fmt.Fprintf(os.Stdout, "packets: %v\n", packets)
		for {
			if len(packets) < 5 {
				return versions, packets
			}
			if packets[0] == 0x01 {
				fmt.Fprintln(os.Stdout, "normal part")
				packets = packets[5:]
				fmt.Fprintf(os.Stdout, "packets: %v\n", packets)
			} else {
				fmt.Fprintln(os.Stdout, "end part")
				packets = packets[5:]
				var v int
				v, packets = decodePacket(packets)
				versions += v
				fmt.Fprintf(os.Stdout, "packets: %v\n", packets)
				break
			}
		}
	default:
		fmt.Fprintln(os.Stdout, "operator")
		if len(packets) < 4 {
			return versions, packets
		}
		// remove type id
		packets = packets[3:]
		switch packets[0] {
		case 0x00:
			fmt.Fprintln(os.Stdout, "length type 0")
			packets = packets[1:]
			if len(packets) < 15 {
				return versions, packets
			}
			sub := binaryToDecimal(packets[:15])
			packets = packets[15:]
			fmt.Fprintf(os.Stdout, "sub packets len: %v\n", sub)
			var v, v2 int
			v, _ = decodePacket(packets[:sub])
			versions += v
			fmt.Fprintf(os.Stdout, "v1: %d\n", versions)
			v2, packets = decodePacket(packets[sub:])
			fmt.Fprintf(os.Stdout, "v2: %d\n", versions)
			versions += v2
		case 0x01:
			fmt.Fprintln(os.Stdout, "length type 1")
			packets = packets[1:]
			if len(packets) < 11 {
				return versions, packets
			}
			sub := binaryToDecimal(packets[:11])
			fmt.Fprintf(os.Stdout, "sub packets: %v\n", sub)
			packets = packets[11:]
			for i := 0; i < sub; i++ {
				var v int
				v, packets = decodePacket(packets)
				versions += v
				fmt.Fprintf(os.Stdout, "%d: %v\n", i, packets)
			}
		}
	}
	return versions, packets
}

func binaryToDecimal(bin []byte) int {
	var decimal int
	for _, b := range bin {
		decimal = decimal*2 + int(b)
	}
	return decimal
}

func hexToBinary(hex []byte) []byte {
	var binary []byte
	for _, h := range hex {
		switch h {
		case '0':
			binary = append(binary, 0x00, 0x00, 0x00, 0x00)
		case '1':
			binary = append(binary, 0x00, 0x00, 0x00, 0x01)
		case '2':
			binary = append(binary, 0x00, 0x00, 0x01, 0x00)
		case '3':
			binary = append(binary, 0x00, 0x00, 0x01, 0x01)
		case '4':
			binary = append(binary, 0x00, 0x01, 0x00, 0x00)
		case '5':
			binary = append(binary, 0x00, 0x01, 0x00, 0x01)
		case '6':
			binary = append(binary, 0x00, 0x01, 0x01, 0x00)
		case '7':
			binary = append(binary, 0x00, 0x01, 0x01, 0x01)
		case '8':
			binary = append(binary, 0x01, 0x00, 0x00, 0x00)
		case '9':
			binary = append(binary, 0x01, 0x00, 0x00, 0x01)
		case 'A':
			binary = append(binary, 0x01, 0x00, 0x01, 0x00)
		case 'B':
			binary = append(binary, 0x01, 0x00, 0x01, 0x01)
		case 'C':
			binary = append(binary, 0x01, 0x01, 0x00, 0x00)
		case 'D':
			binary = append(binary, 0x01, 0x01, 0x00, 0x01)
		case 'E':
			binary = append(binary, 0x01, 0x01, 0x01, 0x00)
		case 'F':
			binary = append(binary, 0x01, 0x01, 0x01, 0x01)
		}
	}
	return binary
}
