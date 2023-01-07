package log

import (
	"fmt"
	"io"
	stdLog "log"
	"os"
	"strings"
)

/*
Logger struct is used to print errors and values in beautifuls texts.

❗ = errors

👉 👈 = represent values

✅ = non-zero value

❌ = zero value

T.H. = Threat Level

T.H is in range of [0:5], 0 indicate fine and 5 indicate fatal errors (only used by (*Logger).Fatal).
*/
type Logger struct {
	prefix string
	l      *stdLog.Logger
}

type Printer struct {
	name   string
	logger *Logger
}

var noFlagsLogger = stdLog.New(os.Stdout, "", 0)

// NewLogger returns a new logger l which logs to writer w.
// If w is nil then l logs to os.Stdout. If flag is 0, then flag = log.Ldate.
// And add space to prefix if it doesn't end with it.
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	if prefix[len(prefix)-1] != ' ' && prefix != "" {
		prefix += " "
	}

	if flag == 0 {
		flag = stdLog.Ldate
	}

	if w == nil {
		w = os.Stdout
	}
	return &Logger{prefix, stdLog.New(w, prefix, flag)}
}

// PrintError Print the error and its threat level. The message contains ❗ to attract attention.
func (logger *Logger) PrintError(err error, th int) {
	if th < 1 {
		th = 1
	}
	x := fmt.Sprintf(":\n❗\t{ ❗: %v ❗ }", fmt.Sprintf("%-15s", err.Error()))
	logger.l.Printf("%-50s T.H. %s", x, nbEmojiMap[th])
}

// Fatal print the error and set th to 5 then exit current program with the given status code. The message contains ☢ to attract attention.
func (logger *Logger) Fatal(err error) {
	logger.l.Printf(":\n☢\t { ☢ : %v ☢ } \t\tT.H. %s\n", fmt.Sprintf("%-15s", err.Error()), nbEmojiMap[5])
	os.Exit(1)
}

// PrintValue print the variable as name-value pair as the passed to the parametre.
// The message contains ✅ to indicate it's non-zero value and ❌ to indicate it is.
func (logger *Logger) PrintValue(varName string, v interface{}) {
	if IsNil(v) {
		x := fmt.Sprintf(":\n❌\t{ 👉 %s: %v 👈 }", fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15v", v))
		logger.l.Printf("%-50s T.H. %s\n", x, nbEmojiMap[1])
		return
	}
	x := fmt.Sprintf(":\n✅\t{ 👉 %s: %v 👈 }", fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15v", v))
	logger.l.Printf("%-50s T.H. %s\n", x, nbEmojiMap[0])
}

// PrintArrayPadding print a slice in a form of | i | v | table for any type t of v.
// Padding is used to determine how much is the space between each i| and v for all i and v.
func PrintArrayPadding[t vALUESType](logger *Logger, arrName string, arr []t, padding int) {
	s := ": " + arrName + ":\n\t" + strings.Repeat("-", padding) + "\n\t "
	for i, v := range arr {
		s += fmt.Sprintf("| %-4d|%-20v|"+"\n", i, v)
		s += "\t" + strings.Repeat("-", padding)
		s += "\n\t "
	}
	logger.l.Println(s)
}

// PrintArrayPadding print a slice in a form of | i | v | table for any type t of v.
func PrintArray[t vALUESType](logger *Logger, arrName string, arr []t) {
	PrintArrayPadding(logger, arrName, arr, 30)
}

// PrintMapPadding print a map in a form of | key | v | table for any type t of v.
// Padding is used to determine how much is the space between each key| and v for all i and v.
func PrintMapPadding[t1 kEYSType, t2 vALUESType](logger *Logger, mapName string, mapp map[t1]t2, padding int) {
	s := ": " + mapName + ":\n\t" + strings.Repeat("-", padding) + "\n\t "
	for i, v := range mapp {
		s += fmt.Sprintf("| %-4v|%-20v|"+"\n", i, v)
		s += "\t" + strings.Repeat("-", padding)
		s += "\n\t "
	}
	logger.l.Println(s)
}

// PrintMapPadding print a map in a form of | key | v | table for any type t of v.
func PrintMap[t1 kEYSType, t2 vALUESType](logger *Logger, mapName string, mapp map[t1]t2) {
	PrintMapPadding(logger, mapName, mapp, 30)
}

// MakePrinter return a printer p which take the function name as passed and prefix everything p logs it.
func (logger *Logger) MakePrinter(funcName string) Printer {
	return Printer{funcName, logger}
}

// PrintError Print the error and its threat level. The message contains ❗ to attract attention.
func (p Printer) PrintError(err error, th int) {
	if th < 1 {
		th = 1
	}
	x := fmt.Sprintf("%s%s :\n❗\t{ ❗: %v ❗ }", p.logger.prefix, p.name, fmt.Sprintf("%-15s", err.Error()))
	noFlagsLogger.Printf("%-60s T.H. %s", x, nbEmojiMap[th])
}

// Fatal print the error and set th to 5 then exit current program with the given status code. The message contains ☢ to attract attention.
func (p Printer) Fatal(err error) {
	noFlagsLogger.Printf("%s%s :\n☢\t { ☢ : %v ☢ }\t\tT.H. %s\n", p.logger.prefix, p.name, fmt.Sprintf("%-15s", err.Error()), nbEmojiMap[5])
	os.Exit(1)
}

// PrintValue print the variable as name-value pair as the passed to the parametre.
// The message contains ✅ to indicate it's non-zero value and ❌ to indicate it is.
func (p Printer) PrintValue(varName string, v interface{}) {
	if IsNil(v) {
		x := fmt.Sprintf("%s%s :\n❌\t{ 👉 %s: %v 👈 }\t", p.logger.prefix, p.name, fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15s", v))
		noFlagsLogger.Printf("%-50s T.H. %s\n", x, nbEmojiMap[1])
		return
	}
	x := fmt.Sprintf("%s%s :\n✅\t{ 👉 %s: %v 👈 }\t", p.logger.prefix, p.name, fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15s", v))
	noFlagsLogger.Printf("%-50s T.H. %s\n", x, nbEmojiMap[0])
}

func IsNil(v interface{}) bool {
	if v == nil || v == 0 || v == "" {
		return true
	}
	return false
}

var nbEmojiMap = map[int]string{
	0:  "0️⃣",
	1:  "1️⃣",
	2:  "2️⃣",
	3:  "3️⃣",
	4:  "4️⃣",
	5:  "5️⃣",
	6:  "6️⃣",
	7:  "7️⃣",
	8:  "8️⃣",
	9:  "9️⃣",
	10: "🔟",
}

type kEYSType interface {
	~string | ~int | ~float64
}

type vALUESType interface {
	kEYSType | ~[]interface{}
}
