package main

import (
	"fmt"
	"time"
)

type item struct {
	name   string
	price  int
	amount int
}

type buyer struct {
	point          int
	shoppingbucket map[string]int
}

type delivery struct {
	status      string
	onedelivery map[string]int
}

func newBuyer() *buyer {
	d := buyer{}
	d.point = 1000000
	d.shoppingbucket = map[string]int{}
	return &d
}

func newDelivery() delivery {
	d := delivery{}
	d.onedelivery = map[string]int{}
	return d
}

func main() {
	numbuy := 0 // 주문한 개수

	tempdelivery := make(map[string]int) // 배달 물품 임시 저장

	items := make([]item, 5) // 물품 목록
	buyer := newBuyer()      // 구매자 정보(장바구니, 마일리지)

	deliverystart := make(chan bool) // 주문 시작 신호 송/수신 채널

	deliverylist := make([]delivery, 5) // 배송 중인 상품 목록

	for i := 0; i < 5; i++ { // 배송 상품 객체 5개 생성
		deliverylist[i] = newDelivery()
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond) //고루틴 순서대로 실행되도록 약간 딜레이
		go deliveryStatus(deliverystart, i, deliverylist, &numbuy, &tempdelivery)
	}

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"롱패딩", 500000, 20}
	items[2] = item{"투미 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}

	for {
		menu := 0 // 첫 메뉴
		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하시오 :")

		fmt.Scanln(&menu)
		fmt.Println()

		if menu == 1 { // 물건 구매

			for {
				itemchoice := 0

				for i := 0; i < 5; i++ {
					fmt.Printf("물품%d: %s,  가격: %d,  잔여 수량: %d\n", i+1, items[i].name, items[i].price, items[i].amount)
				}
				fmt.Print("구매할 물품을 선택하세요 :")
				fmt.Scanln(&itemchoice)
				fmt.Println()
				if itemchoice == 1 {
					buying(items, buyer, itemchoice, deliverylist, deliverystart, &numbuy, tempdelivery)
					break
				} else if itemchoice == 2 {
					buying(items, buyer, itemchoice, deliverylist, deliverystart, &numbuy, tempdelivery)
					break
				} else if itemchoice == 3 {
					buying(items, buyer, itemchoice, deliverylist, deliverystart, &numbuy, tempdelivery)
					break
				} else if itemchoice == 4 {
					buying(items, buyer, itemchoice, deliverylist, deliverystart, &numbuy, tempdelivery)
					break
				} else if itemchoice == 5 {
					buying(items, buyer, itemchoice, deliverylist, deliverystart, &numbuy, tempdelivery)
					break
				} else {
					fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
				}
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 2 { // 남은 수량 확인
			for i := 0; i < 5; i++ {
				fmt.Printf("%s, 잔여 수량: %d\n", items[i].name, items[i].amount)
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 3 {
			fmt.Printf("현재 잔여 마일리지는 %d점입니다.\n", buyer.point)
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 4 { // 배송 상태 확인
			total := 0
			for i := 0; i < 5; i++ {
				total += len(deliverylist[i].onedelivery)
			}
			if total == 0 {
				fmt.Println("배송중인 상품이 없습니다.")
			} else {
				for i := 0; i < len(deliverylist); i++ {
					if len(deliverylist[i].onedelivery) != 0 { // 배송중인 항목만 출력
						for index, val := range deliverylist[i].onedelivery {
							fmt.Printf("%s %d개/ ", index, val)
						}
						fmt.Printf("배송상황: %s\n", deliverylist[i].status)
					}
				}
			}
			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 5 { // 장바구니 확인
			bucketmenu := 0
			for {
				emptyBucket(buyer) // 장바구니 비었는지 확인, 안비었으면 물품 출력

				canbuy := requiredPoint(items, buyer)
				canbuy = excessAmount(items, buyer)

				fmt.Println("1. 장바구니 상품 주문")
				fmt.Println("2. 장바구니 초기화")
				fmt.Println("3. 메뉴로 돌아가기")
				fmt.Print("실행할 기능을 입력하시오 :")
				fmt.Scanln(&bucketmenu)
				fmt.Println()

				if bucketmenu == 1 {
					if canbuy {
						bucketBuying(items, buyer, &numbuy, tempdelivery, deliverystart)
						break
					} else {
						fmt.Println("구매할 수 없습니다.")
						break
					}
				} else if bucketmenu == 2 {
					buyer.shoppingbucket = map[string]int{} // 장바구니 초기화
					fmt.Println("장바구니를 초기화했습니다.")
					break
				} else if bucketmenu == 3 {
					fmt.Println()
					break
				} else {
					fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
				}
			}

			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 6 { // 프로그램 종료
			fmt.Println("프로그램을 종료합니다.")
			return
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
		}
	}

}

func buying(itm []item, byr *buyer, itmchoice int, dlt []delivery, d chan bool, num *int, temp map[string]int) {
	inputamount := 0 // 구매 수량

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "\n")
		}
	}()

	fmt.Print("수량을 입력하시오 :")
	fmt.Scanln(&inputamount)
	fmt.Println()
	if inputamount <= 0 {
		panic("올바른 수량을 입력하세요.")
	}

	if byr.point < itm[itmchoice-1].price*inputamount || itm[itmchoice-1].amount < inputamount { // 수량, 포인트로 구매 가능 여부
		panic("주문이 불가능합니다.")
	} else { // 살 수 있다면
		for {
			buy := 0 // 살지 장바구니 담을지
			fmt.Println("1. 바로 주문\n2. 장바구니에 담기")
			fmt.Print("실행할 기능을 입력하시오 :")
			fmt.Scanln(&buy)
			fmt.Println()

			if buy == 1 { // 바로 구매

				if *num < 5 {
					itm[itmchoice-1].amount -= inputamount
					byr.point -= itm[itmchoice-1].price * inputamount
					temp[itm[itmchoice-1].name] = inputamount

					d <- true

					*num++

					fmt.Println("상품이 주문 접수 되었습니다.")
					break
				} else {
					fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
					break
				}

			} else if buy == 2 { // 장바구니에 담기
				checkbucket := false

				for itms := range byr.shoppingbucket { // 물품 체크
					if itms == itm[itmchoice-1].name {
						checkbucket = true
					}
				}

				if checkbucket == true { // 장바구니에 중복되는 물품이 있을 때
					temp := byr.shoppingbucket[itm[itmchoice-1].name] + inputamount
					if temp > itm[itmchoice-1].amount {
						fmt.Println("물품의 잔여 수량을 초과했습니다.")
						break
					} else {
						byr.shoppingbucket[itm[itmchoice-1].name] += inputamount // 수량만 더함
					}
				} else { // 장바구니에 중복되는 물품이 없을 때
					byr.shoppingbucket[itm[itmchoice-1].name] = inputamount // 새로 품목 추가
				}

				fmt.Println("상품이 장바구니에 추가되었습니다.")
				break
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
			}
		}

	}
}

func emptyBucket(byr *buyer) {
	if len(byr.shoppingbucket) == 0 {
		fmt.Println("장바구니가 비었습니다.")
	} else {
		for index, val := range byr.shoppingbucket {
			fmt.Printf("%s, 수량: %d\n", index, val)
		}
	}
	fmt.Println()
}

func requiredPoint(itm []item, byr *buyer) (canbuy bool) {
	bucketpoint := 0
	for index, val := range byr.shoppingbucket { // 총 필요 마일리지 계산
		for i := 0; i < len(itm); i++ {
			if itm[i].name == index {
				bucketpoint += itm[i].price * val
			}
		}
	}
	fmt.Println("필요 마일리지 :", bucketpoint)
	fmt.Println("보유 마일리지 :", byr.point)
	fmt.Println()
	if byr.point < bucketpoint {
		fmt.Println("마일리지가 %d점 부족합니다.", bucketpoint-byr.point)
		return false
	}
	return true
}

func excessAmount(itm []item, byr *buyer) (canbuy bool) {
	for index, val := range byr.shoppingbucket {
		for i := 0; i < len(itm); i++ {
			if itm[i].name == index {
				if itm[i].amount < val { // 장바구니의 물품 총 개수가 판매하는 물품 개수보다 클 때
					fmt.Println("%s, %d개초과", itm[i].name, val-itm[i].amount)
					return false
				}
			}
		}
	}
	return true
}

func bucketBuying(itm []item, byr *buyer, num *int, temp map[string]int, d chan bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n", r, "\n")

		}
	}()

	if len(byr.shoppingbucket) == 0 {
		panic("주문 가능한 목록이 없습니다.")
	} else { // 장바구니 물품 구입하기
		if *num < 5 {
			for index, val := range byr.shoppingbucket {
				temp[index] = val

				for i := range itm {
					if itm[i].name == index {
						itm[i].amount -= val            // 수량 차감
						byr.point -= itm[i].price * val // 포인트 차감
					}
				}
			}
			d <- true // 배송 시작

			byr.shoppingbucket = map[string]int{} // 장바구니 초기화
			*num++

			fmt.Println("주문 접수 되었습니다.")
		} else {
			fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
		}
	}
}

func deliveryStatus(d chan bool, i int, deliverylist []delivery, num *int, temp *map[string]int) {
	for {
		if <-d {
			for index, val := range *temp {
				deliverylist[i].onedelivery[index] = val // 임시 저장한 데이터를 배송 상품에 저장함
			}

			*temp = map[string]int{} // 임시 데이터 초기화

			deliverylist[i].status = "주문접수"
			time.Sleep(time.Second * 10)

			deliverylist[i].status = "배송중"
			time.Sleep(time.Second * 30)

			deliverylist[i].status = "배송완료"
			time.Sleep(time.Second * 10)

			deliverylist[i].status = ""
			*num--
			deliverylist[i].onedelivery = map[string]int{} // 배송 리스트에서 물품 지우기
		}
	}
}
