package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	strconv "strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

var (
	Token          string
	TComplete      int
	GuildCount     int
	NitroMax       int
	Cooldown       int
	GiveawaySniper bool
	SnipeOnMain    bool
	DMHost         bool
	DMMsg          string
	NitroSniped    int
	SniperRunning  bool

	re                = regexp.MustCompile("(discord.com/gifts/|discordapp.com/gifts/|discord.gift/)([a-zA-Z0-9]+)")
	_                 = regexp.MustCompile("https://privnote.com/.*")
	reGiveaway        = regexp.MustCompile("You won the \\*\\*(.*)\\*\\*")
	reGiveawayMessage = regexp.MustCompile("<https://discordapp.com/channels/(.*)/(.*)/(.*)>")
	magenta           = color.New(color.FgMagenta)
	himagenta         = color.New(color.FgHiMagenta)
	green             = color.New(color.FgGreen)
	higreen           = color.New(color.FgHiGreen)
	yellow            = color.New(color.FgYellow)
	hiyellow          = color.New(color.FgHiYellow)
	red               = color.New(color.FgRed)
	hired             = color.New(color.FgHiRed)
	cyan              = color.New(color.FgCyan)
	hicyan            = color.New(color.FgHiCyan)
	strPost           = []byte("POST")
	strGet            = []byte("GET")
	_                 = []byte("GET")
	Tokens            []string
	lCnt              int
	triedC            []string
	didLoadT          bool
	intCnt            int
	appversion        string
	UserID            string
	UserN             []string
	startT            time.Time
	endT              time.Duration
	wg                sync.WaitGroup
)

type Thread struct {
	i int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lCnt = 0
	for scanner.Scan() {
		Tokens = append(Tokens, scanner.Text())
		lCnt++
	}
	return Tokens, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
func ClearCLI() {
	if isWindows() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}
func init() {
	appversion = "v3.1.5"
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if isWindows() {
		path = path + "\\tokens.txt"
	} else {
		path = path + "/tokens.txt"
	}
	Tokens, err = readLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	file, err := ioutil.ReadFile("settings.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed read file: %s\n", err)
		os.Exit(1)
	}

	var f interface{}
	err = json.Unmarshal(file, &f)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to parse JSON: %s\n", err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})

	str := fmt.Sprintf("%v", m["token"])
	flag.StringVar(&Token, "t", str, "Token")

	if Token == "put your token here" {
		hired.Print("You haven't properly configured the 'settings.json' file. Please put your Discord authorization token in settings.json using the correct JSON syntax and then run the program again.")
		didLoadT = false
	} else {
		didLoadT = true
	}

	str2 := fmt.Sprintf("%f", m["nitro_max"])
	value, _ := strconv.ParseFloat(str2, 64)
	flag.IntVar(&NitroMax, "m", int(value), "NitroMax")

	str3 := fmt.Sprintf("%t", m["giveaway_sniper"])
	value2, _ := strconv.ParseBool(str3)
	flag.BoolVar(&GiveawaySniper, "g", value2, "GiveawaySniper")

	str4 := fmt.Sprintf("%f", m["cooldown"])
	value3, _ := strconv.ParseFloat(str4, 64)
	flag.IntVar(&Cooldown, "c", int(value3), "cooldown")

	str5 := fmt.Sprintf("%t", m["snipe_on_main"])
	value4, _ := strconv.ParseBool(str5)
	flag.BoolVar(&SnipeOnMain, "s", value4, "SnipeOnMain")

	str6 := fmt.Sprintf("%t", m["dm_host"])
	value5, _ := strconv.ParseBool(str6)
	flag.BoolVar(&DMHost, "d", value5, "DMHost")

	DMMsg = fmt.Sprintf("%s", m["dm_message"])

	flag.Parse()

	NitroSniped = 0
	SniperRunning = true
}

func timerEnd() {
	SniperRunning = true
	NitroSniped = 0
	_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
	_, _ = higreen.Print("[+] Starting Nitro sniping")
}
func loadSniper(wg *sync.WaitGroup, str string, id int) {
	dg, err := discordgo.New(str)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	e := Thread{id}
	dg.AddHandler(e.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	ClearCLI()
	t := time.Now()
	GuildCount += len(dg.State.Guilds)
	TComplete++
	if TComplete == intCnt {
		color.HiGreen(`	
 ███▄    █  ██▓▄▄▄█████▓ ██▀███   ▒█████         ██████ ▓█████  ██▓      █████▒
 ██ ▀█   █ ▓██▒▓  ██▒ ▓▒▓██ ▒ ██▒▒██▒  ██▒     ▒██    ▒ ▓█   ▀ ▓██▒    ▓██   ▒ 
▓██  ▀█ ██▒▒██▒▒ ▓██░ ▒░▓██ ░▄█ ▒▒██░  ██▒     ░ ▓██▄   ▒███   ▒██░    ▒████ ░ 
▓██▒  ▐▌██▒░██░░ ▓██▓ ░ ▒██▀▀█▄  ▒██   ██░       ▒   ██▒▒▓█  ▄ ▒██░    ░▓█▒  ░ 
▒██░   ▓██░░██░  ▒██▒ ░ ░██▓ ▒██▒░ ████▓▒░ ██▓ ▒██████▒▒░▒████▒░██████▒░▒█░    
░ ▒░   ▒ ▒ ░▓    ▒ ░░   ░ ▒▓ ░▒▓░░ ▒░▒░▒░  ▒▓▒ ▒ ▒▓▒ ▒ ░░░ ▒░ ░░ ▒░▓v3.1.5░    
░ ░░   ░ ▒░ ▒ ░    ░      ░▒ ░ ▒░  ░ ▒ ▒░  ░▒  ░ ░▒  ░ ░ ░ ░  ░░ ░ ▒  ░ ░      
   ░   ░ ░  ▒ ░  ░        ░░   ░ ░ ░ ░ ▒   ░   ░  ░  ░     ░     ░ ░    ░ ░    
         ░  ░              ░         ░ ░    ░        ░     ░  ░    ░  ░        
                                            ░                                  
	`)
		checkUpdate()
		himagenta.Print(t.Format("15:04:05 "))
		hicyan.Print("Sniping Discord Nitro Codes and Giveaways on ")
		hiyellow.Print(strconv.Itoa(GuildCount))
		hicyan.Print(" Servers with ")
		hiyellow.Print(strconv.Itoa(intCnt))
		hicyan.Println(" Accounts 🔫")
		//_, _ = himagenta.Print(t.Format("15:04:05 "))
		//higreen.Println("[+] If we're lucky you'll get Nitro on " + ) need to setup a way to detect main user here.
		UserID = dg.State.User.ID

		//UserN[id] = dg.State.User.String()
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()
	defer wg.Done()
}

func checkUpdate() {
	var strRequestURI = []byte("https://raw.githubusercontent.com/noto-rious/Nitro.Self-V3/master/version.txt")
	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes(strGet)
	req.SetRequestURIBytes(strRequestURI)
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err != nil {
		panic("handle error")
	}

	fasthttp.ReleaseRequest(req)

	body := res.Body()

	bodyString := string(body)
	if appversion != bodyString {
		hired.Println("Looks like you may not be running the most current version. Check https://noto.cf/ for an update!\n")
	}
}
func redeemCode(theCode string, theChannel string, theAuthor string, theUser string) {
	ch := make(chan int)

	go func() {

		ch <- 1
	}()
	<-ch
}
func (e *Thread) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//do stuff with e.i

	ch := make(chan int)
	//fmt.Println(s.State.User.ID)
	go func() {
		if re.Match([]byte(m.Content)) && SniperRunning {

			code := re.FindStringSubmatch(m.Content)

			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			if err != nil {
				log.Fatal(err)
			}

			code[2] = reg.ReplaceAllString(code[2], "")

			if len(code[2]) < 2 {
				return
			}

			if len(code[2]) < 16 {
				//_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
				//_, _ = hired.Print("[=] Auto-detected a fake code: ")
				//_, _ = hired.Print(code[2])
				//_, _ = fmt.Println(" from " + m.Author.String())
				return
			}

			_, found := Find(triedC, code[2])
			if found != true {
				triedC = append(triedC, code[2])
			} else if found == true {
				//_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
				//_, _ = hired.Print("[=] Auto-detected a dupe code: ")
				//_, _ = hired.Print(code[2])
				//_, _ = fmt.Println(" from " + m.Author.String())
				return
			}

			println()
			_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
			_, _ = hicyan.Print(s.State.User.String() + " -> ")
			//return
			guild, err := s.State.Guild(m.GuildID)
			if err != nil || guild == nil {
				guild, err = s.Guild(m.GuildID)
				if err != nil {
					_, _ = hiyellow.Println("[DM with " + m.Author.String() + " > " + s.State.User.String() + "]")
				}
			}
			channel, err := s.State.Channel(m.ChannelID)
			if err != nil || guild == nil {
				channel, err = s.Channel(m.ChannelID)
				if err != nil {
				}
			} else if guild != nil {
				_, _ = hiyellow.Println("[" + guild.Name + " > " + channel.Name + " > " + m.Author.String() + "]")
			}

			_, _ = higreen.Print("[-] Checking code: ")
			_, _ = higreen.Print(code[2])
			_, _ = higreen.Println("...")

			startT = time.Now()
			var strRequestURI = []byte("https://discordapp.com/api/v8/entitlements/gift-codes/" + code[2] + "/redeem")
			req := fasthttp.AcquireRequest()
			req.Header.SetContentType("application/json")
			req.Header.Set("Authorization", Token)
			req.SetBody([]byte(`{"channel_id":` + m.ChannelID + "}"))
			req.Header.SetMethodBytes(strPost)
			req.SetRequestURIBytes(strRequestURI)
			res := fasthttp.AcquireResponse()
			//endT = time.Since(startT)

			if err := fasthttp.Do(req, res); err != nil {
				panic("handle error")
			}
			fasthttp.ReleaseRequest(req)
			endT = time.Since(startT)

			body := res.Body()

			bodyString := string(body)
			fasthttp.ReleaseResponse(res)

			if strings.Contains(bodyString, "This gift has been redeemed already.") || strings.Contains(bodyString, "Already purchased") || strings.Contains(bodyString, "Missing Access") {
				_, _ = hiyellow.Print("[-] Code has already been redeemed.")
				_, _ = fmt.Print(" - ")
				_, _ = hiyellow.Print("Delay: ")
				_, _ = hiyellow.Print(endT)
				_, _ = hiyellow.Print("\n")

			} else if strings.Contains(bodyString, "nitro") {
				_, _ = higreen.Print("[+] Code successfully redeemed!!!")
				_, _ = fmt.Print(" - ")
				_, _ = hiyellow.Print("Delay: ")
				_, _ = hiyellow.Print(endT)
				_, _ = hiyellow.Print("\n")
				NitroSniped++
				if NitroSniped == NitroMax {
					SniperRunning = false
					time.AfterFunc(time.Hour*time.Duration(Cooldown), timerEnd)
					_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
					_, _ = hiyellow.Println("[+] Stopping Nitro sniping for now")
				}
			} else if strings.Contains(bodyString, "You are being rate limited") {
				_, _ = hired.Print("[x] You are rate limited.")
				_, _ = fmt.Print(" - ")
				_, _ = hiyellow.Print("Delay: ")
				_, _ = hiyellow.Print(endT)
				_, _ = hiyellow.Print("\n")
			} else if strings.Contains(bodyString, "Unknown Gift Code") {
				_, _ = hired.Print("[x] Code was fake or expired.")
				_, _ = fmt.Print(" - ")
				_, _ = hiyellow.Print("Delay: ")
				_, _ = hiyellow.Print(endT)
				_, _ = hiyellow.Print("\n")
			} else {
				_, _ = hiyellow.Print("[?] Could not validate this code.")
				_, _ = fmt.Print(" - ")
				_, _ = hiyellow.Print("Delay: ")
				_, _ = hiyellow.Print(endT)
				_, _ = hiyellow.Print("\n")
				fmt.Println(bodyString)
			}
			//println()
		} else if GiveawaySniper && (strings.Contains(strings.ToLower(m.Content), "**giveaway**") || (strings.Contains(strings.ToLower(m.Content), "react with") && strings.Contains(strings.ToLower(m.Content), "giveaway"))) {
			if len(m.Embeds) > 0 && m.Embeds[0].Author != nil {
				if !strings.Contains(strings.ToLower(m.Embeds[0].Author.Name), "nitro") {
					return
				}
			} else {
				return
			}
			time.Sleep(time.Second)

			guild, err := s.State.Guild(m.GuildID)
			if err != nil || guild == nil {
				guild, err = s.Guild(m.GuildID)
				if err != nil {

				}
			}
			channel, err := s.State.Channel(m.ChannelID)
			if err != nil || guild == nil {
				channel, err = s.Channel(m.ChannelID)
				if err != nil {
				}
			} else if guild != nil {
			}
			err = s.MessageReactionAdd(m.ChannelID, m.ID, "🎉")
			time.Sleep(500 * time.Millisecond)
			if err != nil {
				println()
				_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
				_, _ = hicyan.Print(s.State.User.String() + " -> ")
				_, _ = hiyellow.Println("[" + guild.Name + " > " + channel.Name + " > " + m.Author.String() + "]")
				_, _ = hired.Println("[x] Failed to enter a giveaway :(")
			} else {
				println()
				_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
				_, _ = hicyan.Print(s.State.User.String() + " -> ")
				_, _ = hiyellow.Println("[" + guild.Name + " > " + channel.Name + " > " + m.Author.String() + "]")
				_, _ = higreen.Println("[+] Entered a Giveaway!")
			}
		} else if (strings.Contains(strings.ToLower(m.Content), "giveaway") || strings.Contains(strings.ToLower(m.Content), "win") || strings.Contains(strings.ToLower(m.Content), "won")) && strings.Contains(m.Content, s.State.User.ID) {
			reGiveawayHost := regexp.MustCompile("Hosted by: <@(.*)>")
			won := reGiveaway.FindStringSubmatch(m.Content)
			giveawayID := reGiveawayMessage.FindStringSubmatch(m.Content)
			guild, err := s.State.Guild(m.GuildID)
			if err != nil || guild == nil {
				guild, err = s.Guild(m.GuildID)
				if err != nil {
					return
				}
			}

			channel, err := s.State.Channel(m.ChannelID)
			if err != nil || guild == nil {
				channel, err = s.Channel(m.ChannelID)
				if err != nil {
					return
				}
			}
			if giveawayID == nil {
				if len(won) > 1 {
					println()
					_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
					_, _ = hicyan.Print(s.State.User.String() + " -> ")
					_, _ = hiyellow.Println("[" + guild.Name + " > " + channel.Name + " > " + m.Author.String() + "]")
					_, _ = higreen.Print("[+] Winner winner, chicken dinner, You won the ")
					_, _ = hicyan.Print(won[1])
					_, _ = higreen.Println(" giveaway!!!")
				}
				return
			}
			messages, _ := s.ChannelMessages(m.ChannelID, 1, "", "", giveawayID[3])

			if len(won) > 1 {
				println()
				_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
				_, _ = hicyan.Print(s.State.User.String() + " -> ")
				_, _ = hiyellow.Println("[" + guild.Name + " > " + channel.Name + " > " + m.Author.String() + "]")
				_, _ = higreen.Print("[+] Winner winner, Chicken dinner, You won the ")
				_, _ = hicyan.Print(won[1])
				_, _ = higreen.Println(" giveaway!!!")
			} else {
				return
			}

			giveawayHost := reGiveawayHost.FindStringSubmatch(messages[0].Embeds[0].Description)
			if len(giveawayHost) < 2 {
				return
			}
			hostChannel, err := s.UserChannelCreate(giveawayHost[1])

			if err != nil {
				return
			}
			time.Sleep(time.Second * 9)

			if DMHost == true {
				_, err = s.ChannelMessageSend(hostChannel.ID, DMMsg)
				if err != nil {
					return
				}

				host, _ := s.User(giveawayHost[1])
				println()
				_, _ = himagenta.Print(time.Now().Format("15:04:05 "))
				_, _ = higreen.Print("[+] ")
				_, _ = hicyan.Print(s.State.User.String())
				_, _ = higreen.Print(" sent DM to host: ")
				_, _ = hiyellow.Print(host.Username + "#" + host.Discriminator + "\n")
			}
		}
		ch <- e.i
	}()
	<-ch
}
func main() {
	if isWindows() == false {
		fmt.Printf("\033]0;Nitro.Self " + appversion + " - Developed By: Notorious\007")
	} else {
		cmd := exec.Command("cmd", "/c", "title", "Nitro.Self "+appversion+" - Developed By: Notorious")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if didLoadT == true {
		if SnipeOnMain {
			intCnt = len(Tokens) + 1
		} else {
			intCnt = len(Tokens)
		}

		// Calling Goroutine
		wg.Add(intCnt)

		intID := 1
		if SnipeOnMain {
			go loadSniper(&wg, Token, intID)
			ClearCLI()
		}

		for _, line := range Tokens {
			intID++
			go loadSniper(&wg, line, intID)
			ClearCLI()
		}

		// Calling normal function
		wg.Wait()
	} else {
		duration := time.Duration(30) * time.Second
		time.Sleep(duration)
	}
}
