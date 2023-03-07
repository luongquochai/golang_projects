package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type nhieuBai []string

func (nb nhieuBai) print() {
	fmt.Println(nb)
}

func taoBaiMoi() nhieuBai {
	boBai := nhieuBai{}

	nhieuNuoc := []string{"co", "ro", "chuon", "bich"}
	nhieuNut := []string{"1", "2", "3"}
	for _, nuoc := range nhieuNuoc {
		for _, nut := range nhieuNut {
			bai := nut + " " + nuoc
			boBai = append(boBai, bai)
		}
	}

	return boBai
}

func (n nhieuBai) chiaBai(sl int) (nhieuBai, nhieuBai) {
	return n[:sl], n[sl:]
}

func (n nhieuBai) chuyenThanhString() string {
	return strings.Join(n, ",")
}

func (n nhieuBai) luuFile(tenFile string) error {
	data := []byte(n.chuyenThanhString())
	return ioutil.WriteFile(tenFile, data, 0666)
}

func taoBaiTuFile(tenFile string) (nhieuBai, error) {
	data, err := ioutil.ReadFile(tenFile)

	if err != nil {
		log.Printf("doc file khong duoc (err=%v)", err)
		return nhieuBai{}, err
	}

	chuoiBai := string(data)
	mang := strings.Split(chuoiBai, ",")
	nhieuLaBai := nhieuBai(mang)

	return nhieuLaBai, nil
}

func (n nhieuBai) xaoBai() {
	rand.Seed(time.Now().UnixNano())
	for index := range n {
		random_pos := rand.Intn(len(n))
		// swap pos
		n[index], n[random_pos] = n[random_pos], n[index]

	}
}
