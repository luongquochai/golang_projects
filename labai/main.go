package main

import "os"

func main() {
	nhieuLaBai := taoBaiMoi()
	nhieuLaBai.luuFile("Cobac")
	// bai, baiConLai := nhieuLaBai.chiaBai(4)
	// bai.print()
	// baiConLai.print()
	nhieuLaBai = append(nhieuLaBai, "9 co")

	nhieuLaBai.print()

	nhieuLaBai, err := taoBaiTuFile("Cobac")
	if err != nil {
		os.Exit(1)
	}

	nhieuLaBai.print()

	// nhieuLaBai.xaoBai()
	// nhieuLaBai.print()
}
