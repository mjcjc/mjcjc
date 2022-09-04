package main

import (
	"discord-bot/config"
	"discord-bot/mapleprofile"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"math"
	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	err = config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	s, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "도움",
			Description: "도움말 출력",
			Type:        discordgo.ChatApplicationCommand,
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "주간-보스-결정석-가격확인",
			Description: "결정석 가격출력",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "보스-이름",
					Description:  "보스-이름",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		{
			Name:        "전적검색",
			Description: "메이플 전적검색",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "아이디",
					Description:  "아이디",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		{
			Name:        "환산-주스탯-사이트-검색",
			Description: "환산 엑셀시트 나옴",
			Type:        discordgo.ChatApplicationCommand,
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "작-계산기",
			Description: "스타포스 입력후 방향키 한번더 입력",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "스타포스",
					Description:  "스타포스 몇성?",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "장비류",
					Description:  "자동입력에 없을시 그냥 입력하세요. 무기는 추옵빼고 전부 다 입력해주세요.",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
				{
					Name:         "힘",
					Description:  "힘수치",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "덱스",
					Description:  "덱스입력",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "럭",
					Description:  "럭입력",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "인트",
					Description:  "인트입력",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "공격력",
					Description:  "공격력입력",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "마력",
					Description:  "마력입력",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "템레벨",
					Description:  "템레벨입력",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		{
			Name:        "분배금결정",
			Description: "분배금 n값 결정",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "분배금",
					Description:  "총 금액을 입력하세요.",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: false,
				},
				{
					Name:         "인원",
					Description:  "인원수를 입력학세요.",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"도움": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			helpMsg := `명령어 사용 ex) /도움
1. 메이플 전적검색
2. 환산 주스탯 사이트 검색
3. 각종 주간보스 결정석 가격 / 수에큐 확인
4. 작-계산기(데벤X)
5. 분배금 결정
`
			// 입력하실때 스타포스-장비-힘-덱-럭-인-공-마 순으로 하셔아합니다!!
			// 잘못 입력하면 원하는 값과 다르게 나와요!!!!
			helpMsg = "```" + helpMsg + "```"

			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: helpMsg,
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},

		"주간-보스-결정석-가격확인": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				data := i.ApplicationCommandData()

				var content string
				if price, ok := bossPrice[data.Options[0].StringValue()]; ok {
					content = fmt.Sprintf(
						"```%s 가격 : %s\n22.09.03 기준```",
						data.Options[0].StringValue(),
						price,
					)
				} else {
					content = fmt.Sprintf(
						"%q 는 존재하지 않는 보스 이름입니다.",
						data.Options[0].StringValue(),
					)
				}

				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: content,
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			// Autocomplete options introduce a new interaction type (8) for returning custom autocomplete results.
			case discordgo.InteractionApplicationCommandAutocomplete:

				choices := []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "카오스 자쿰",
						Value: "카오스 자쿰",
					},
					{
						Name:  "하드 매그너스",
						Value: "하드 매그너스",
					},
					{
						Name:  "하드 힐라",
						Value: "하드 힐라",
					},
					{
						Name:  "카오스 파풀라투스",
						Value: "카오스 파풀라투스",
					},
					{
						Name:  "피에르",
						Value: "피에르",
					},
					{
						Name:  "반반",
						Value: "반반",
					},
					{
						Name:  "퀸",
						Value: "퀸",
					},
					{
						Name:  "벨룸",
						Value: "벨룸",
					},
					{
						Name:  "카오스 핑크빈",
						Value: "카오스 핑크빈",
					},
					{
						Name:  "시그너스",
						Value: "시그너스",
					},
					{
						Name:  "스우",
						Value: "스우",
					},
					{
						Name:  "데미안",
						Value: "데미안",
					},
					{
						Name:  "가엔슬",
						Value: "가엔슬",
					},
					{
						Name:  "루시드",
						Value: "루시드",
					},
					{
						Name:  "윌",
						Value: "윌",
					},
					{
						Name:  "더스크",
						Value: "더스크",
					},
					{
						Name:  "진 힐라",
						Value: "진 힐라",
					},
					{
						Name:  "듄켈",
						Value: "듄켈",
					},
					{
						Name:  "검마",
						Value: "검마",
					},
					{
						Name:  "세렌",
						Value: "세렌",
					},
					{
						Name:  "칼로스",
						Value: "칼로스",
					},
				}
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionApplicationCommandAutocompleteResult,
					Data: &discordgo.InteractionResponseData{
						Choices: choices, // This is basically the whole purpose of autocomplete interaction - return custom options to the user.
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},

		"전적검색": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				data := i.ApplicationCommandData()
				ign := data.Options[0].StringValue()

				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					// Type: discordgo.InteractionResponseChannelMessageWithSource,
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "fetching data...",
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}

				go func() {
					// download image from maple.gg
					mapleprofile.DownloadMapleProfile(ign)

					fileName := ign + ".png"
					f, err := os.Open("./" + fileName)
					if err != nil {
						errMsg := err.Error()
						_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
							Content: &errMsg,
						})
						if err != nil {
							fmt.Println(err.Error())
						}
					} else {
						_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
							Files: []*discordgo.File{
								{
									Name:   fileName,
									Reader: f,
								},
							},
						})
						if err != nil {
							fmt.Println(err.Error())
						}
					}

					f.Close()
					err = os.Remove("./" + fileName)
					if err != nil {
						fmt.Println(err.Error())
					}
				}()
			}
		},

		"환산-주스탯-사이트-검색": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			siteDoc := "구글 아이디 있어야함\n" +
				"https://docs.google.com/spreadsheets/d/1hD-P07knD8FlzWvUOGm_3f-NMZxTTQ14IivqddUMXOY/copy#gid=1397551479"
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "```" + siteDoc + "```",
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},
		"작-계산기": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.ApplicationCommandData()
			switch i.Type {	
			case discordgo.InteractionApplicationCommand:
			stat := 0
			EnAtt := 0
			j :=0
			
			Star, err := strconv.Atoi(data.Options[0].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			
			STR,err := strconv.Atoi(data.Options[2].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			DEX,err := strconv.Atoi(data.Options[3].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			LUK,err := strconv.Atoi(data.Options[4].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			INT,err := strconv.Atoi(data.Options[5].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			Att,err := strconv.Atoi(data.Options[6].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			Mtt,err := strconv.Atoi(data.Options[7].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			lev,err := strconv.Atoi(data.Options[8].StringValue())
			if err != nil {
				fmt.Println(err.Error())
			}
			equip := data.Options[1].StringValue()
				if 0 <= Star &&Star <=5{
					for i := 0 ; i<=Star; i++{
						stat +=2
					}
				}else if 6<= Star&& Star <=15{
					stat = 13
					for i := 6 ; i<=Star; i++{
						stat += 3
					}
				}else{
					stat=40
				}
				if lev <= 130{
					EnAtt = 0
					j = 7
						for i := 16; i<=Star; i++{
							stat +=7
							EnAtt += j
							j += 1
						}
				}else if lev <= 140{
					EnAtt = 0
					j = 8
						for i := 16 ; i<=Star; i++{
							stat +=9
							EnAtt += j
							j += 1
						}
				}else if lev <=150{
					 if 16 < Star {
						 EnAtt = 0
						 j = 9
						for i := 16 ; i<=Star; i++{
							stat +=11
							 EnAtt += j 
							j += 1
						}
					}
				}else if lev <=160{
					if 16 <= Star{
						EnAtt = 0
						j = 10
						for i := 16 ; i<=Star; i++{
							stat +=13
							EnAtt +=j
							j += 1
						}	
					}
				}else if lev <=200{
					if 16 <= Star {
					EnAtt = 0
						 j = 12
						for i := 16 ; i<=Star; i++{
							stat +=15
							EnAtt += j
							j += 1
						}
					}
				}else{
					fmt.Println(err.Error())
				}
				if equip == "장갑"{
					EnAtt =0
					if Star<= 5{
						EnAtt = 1
					}else if Star <=7{
						EnAtt = 2
					}else if Star <=9{
						EnAtt = 3
					}else if Star <=11{
						EnAtt = 4
					}else if Star ==13{
						EnAtt = 5
					}else if Star ==14{
						EnAtt = 6
					}else if Star == 15{
						EnAtt = 7
					}
					if lev == 160{
						EnAtt =7
						j = 10
						for i := 16 ; i<=Star; i++{
							EnAtt +=j
							j += 1
						}	
					}else if lev ==200{
						EnAtt =7
						j = 12
						if Star <=21{
							for i := 16 ; i<=Star; i++{
								EnAtt +=j
								j += 1
							}
						}else if Star == 22{
							EnAtt =109
						}	
					}
				}
				if equip =="무기"{
					EnAtt=0
					var tn float64 =float64(Att)
					for i :=Star; i>0; i--{  
						if i == 22{
							if lev <=150{
								tn=tn-72.0
								i=16
							}else if lev <=160{
								tn = tn-78.0
								i=16
							}else if lev <=200{
								tn=tn-102.0
								i=16
							}
						}else if i == 21{
							if lev <=150{
								tn=tn-59.0
								i=16
							}else if lev <=160{
								tn=tn-64.0
								i=16
							}else if lev <=200{
								tn=tn-85.0
								i=16
							}
						}else if i == 20{
							if lev <=150{
								tn=tn-47.0
								i=16
							}else if lev <=160{
								tn=tn-51.0
								i=16
							}else if lev <=200{
								tn=tn-69.0
								i=16
							}							
						}else if i == 19{
							if lev <=150{
								tn=tn-36.0
								i=16
							}else if lev <=160{
								tn=tn-39.0
								i=16
							}else if lev <=200{
								tn=tn-54.0
								i=16
							}	
						}else if i == 18{
							if lev <=150{
								tn=tn-26.0
								i=16
							}else if lev <=160{
								tn=tn-28.0
								i=16
							}else if lev <=200{
								tn=tn-40.0
								i=16
							}
						}else if i ==17||i==16{
							if lev <=150{
								if i ==17{
								tn=tn-9.0
							}else{
								tn=tn-8.0
							}
							}else if lev <=160{
								tn=tn-9.0
							}else if lev <=200{
								tn=tn-13.0
							}							
						}else{
							tn = math.Ceil((50*tn-50)/51)
						}
					}
					Att = int(tn)
				}
				STR = STR-stat
				DEX = DEX-stat
				LUK = LUK-stat
				INT = INT-stat
				Att = Att-EnAtt
				Mtt = Mtt-EnAtt
				if STR <0 {
					STR = 0
				}					
				if DEX <0 {
					DEX = 0
				}
				if LUK < 0{
					LUK = 0
				}
				if  INT <0{
					INT = 0
				}
				if Att <0{
					Att = 0
				}
				if Mtt <0{
					Mtt = 0
				}
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{

						Content: fmt.Sprintf(
							"```장비: %s  스타포스:%d 힘:%d 덱:%d 럭:%d 인트:%d 공격력:%d 마력:%d 레벨:%d\n ***주흔작은 4작때 공마1 붙습니다.***\n 무기는 자체 공/마 수치 빼주세요.```",
							equip, Star, STR , DEX, LUK, INT, Att, Mtt, lev,
						),
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			case discordgo.InteractionApplicationCommandAutocomplete:
				var choices []*discordgo.ApplicationCommandOptionChoice
				switch {
				case data.Options[1].Focused:
					choices = []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "장갑",
							Value: "장갑",
						},
						{
							Name:  "무기",
							Value: "무기",
						},
					}
					if data.Options[1].StringValue() != "" {
						choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
							Name:  data.Options[1].StringValue(),
							Value: data.Options[1].StringValue(),
						})
					}
				case data.Options[8].Focused:
					choices = []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "130",
							Value: "130",
						},
						{
							Name:  "140",
							Value: "140",
						},
						{
							Name:  "150",
							Value: "150",
						},
						{
							Name:  "160",
							Value: "160",
						},
						{
							Name:  "200",
							Value: "200",
						},
					}
					if data.Options[8].StringValue() != "" {
						choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
							Name:  data.Options[8].StringValue(), // To get user input you just get value of the autocomplete option.
							Value: data.Options[8].StringValue(),
						})
					}
					
			
				}
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionApplicationCommandAutocompleteResult,
					Data: &discordgo.InteractionResponseData{
						Choices: choices, // This is basically the whole purpose of autocomplete interaction - return custom options to the user.
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			// Autocomplete options introduce a new interaction type (8) for returning custom autocomplete results.
		},
		"분배금결정": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				data := i.ApplicationCommandData()
				pri := data.Options[0].StringValue()
				per := data.Options[1].StringValue()
				errMsg := ""

				intper, err := strconv.Atoi(per)
				if err != nil {
					errMsg = err.Error()
				}

				intPri, err := strconv.Atoi(pri)
				if err != nil {
					if errMsg != "" {
						errMsg += "\n"
					}
					errMsg += err.Error()
				}

				if errMsg == "" {
					pern := intPri / intper
					err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: fmt.Sprintf(
								"```분배금 금액: %s\n인당 분배금: %d```",
								data.Options[0].StringValue(),
								pern,
							),
						},
					})
					if err != nil {
						fmt.Println(err.Error())
					}
				} else {
					err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: fmt.Sprintf("```잘못 입력하셨습니다. \n입력된 값: %s```",per),
						},
					})
					if err != nil {
						fmt.Println(err.Error())
					}
				}

			case discordgo.InteractionApplicationCommandAutocomplete:
				data := i.ApplicationCommandData()
				var choices []*discordgo.ApplicationCommandOptionChoice
				switch {
				case data.Options[1].Focused:
					choices = []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "2",
							Value: "2",
						},
						{
							Name:  "3",
							Value: "3",
						},
						{
							Name:  "4",
							Value: "4",
						},
						{
							Name:  "5",
							Value: "5",
						},
						{
							Name:  "6",
							Value: "6",
						},
					}
				}
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionApplicationCommandAutocompleteResult,
					Data: &discordgo.InteractionResponseData{
						Choices: choices,
					},
				})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},
	}
)
func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { log.Println("봇 켜짐") })
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	createdCommands, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		log.Fatalf("Cannot register commands: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("종료")

	if config.RemoveCommands {
		for _, cmd := range createdCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
			if err != nil {
				log.Fatalf("Cannot delete %q command: %v", cmd.Name, err)
			}
		}
	}
}
