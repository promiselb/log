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

ā = errors

š š = represent values

ā = non-zero value

ā = zero value

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

// PrintError Print the error and its threat level. The message contains ā to attract attention.
func (logger *Logger) PrintError(err error, th int) {
	if th < 1 {
		th = 1
	}
	x := fmt.Sprintf(":\nā\t{ ā: %v ā }", fmt.Sprintf("%-15s", err.Error()))
	logger.l.Printf("%-50s T.H. %s", x, nbEmojiMap[th])
}

// Fatal print the error and set th to 5 then exit current program with the given status code. The message contains ā¢ to attract attention.
func (logger *Logger) Fatal(err error) {
	logger.l.Printf(":\nā¢\t { ā¢ : %v ā¢ } \t\tT.H. %s\n", fmt.Sprintf("%-15s", err.Error()), nbEmojiMap[5])
	os.Exit(1)
}

// PrintValue print the variable as name-value pair as the passed to the parametre.
// The message contains ā to indicate it's non-zero value and ā to indicate it is.
func (logger *Logger) PrintValue(varName string, v interface{}) {
	if IsNil(v) {
		x := fmt.Sprintf(":\nā\t{ š %s: %v š }", fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15v", v))
		logger.l.Printf("%-50s T.H. %s\n", x, nbEmojiMap[1])
		return
	}
	x := fmt.Sprintf(":\nā\t{ š %s: %v š }", fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15v", v))
	logger.l.Printf("%-50s T.H. %s\n", x, nbEmojiMap[0])
}

// PrintArrayPadding print a slice in a form of | i | v | table for any type t of v.
// Padding is used to determine how much is the space between each i| and v for all i and v.
func PrintArrayPadding[t VALUESType](logger *Logger, arrName string, arr []t, padding int) {
	s := ": " + arrName + ":\n\t" + strings.Repeat("-", padding) + "\n\t "
	for i, v := range arr {
		s += fmt.Sprintf("| %-4d|%-20v|"+"\n", i, v)
		s += "\t" + strings.Repeat("-", padding)
		s += "\n\t "
	}
	logger.l.Println(s)
}

// PrintArrayPadding print a slice in a form of | i | v | table for any type t of v.
func PrintArray[t VALUESType](logger *Logger, arrName string, arr []t) {
	PrintArrayPadding(logger, arrName, arr, 30)
}

// PrintMapPadding print a map in a form of | key | v | table for any type t of v.
// Padding is used to determine how much is the space between each key| and v for all i and v.
func PrintMapPadding[t1 KEYSType, t2 VALUESType](logger *Logger, mapName string, mapp map[t1]t2, padding int) {
	s := ": " + mapName + ":\n\t" + strings.Repeat("-", padding) + "\n\t "
	for i, v := range mapp {
		s += fmt.Sprintf("| %-4v|%-20v|"+"\n", i, v)
		s += "\t" + strings.Repeat("-", padding)
		s += "\n\t "
	}
	logger.l.Println(s)
}

// PrintMapPadding print a map in a form of | key | v | table for any type t of v.
func PrintMap[t1 KEYSType, t2 VALUESType](logger *Logger, mapName string, mapp map[t1]t2) {
	PrintMapPadding(logger, mapName, mapp, 30)
}

// MakePrinter return a printer p which take the function name as passed and prefix everything p logs it.
func (logger *Logger) MakePrinter(funcName string) Printer {
	return Printer{funcName, logger}
}

// PrintError Print the error and its threat level. The message contains ā to attract attention.
func (p Printer) PrintError(err error, th int) {
	if th < 1 {
		th = 1
	}
	x := fmt.Sprintf("%s%s :\nā\t{ ā: %v ā }", p.logger.prefix, p.name, fmt.Sprintf("%-15s", err.Error()))
	noFlagsLogger.Printf("%-60s T.H. %s", x, nbEmojiMap[th])
}

// Fatal print the error and set th to 5 then exit current program with the given status code. The message contains ā¢ to attract attention.
func (p Printer) Fatal(err error) {
	noFlagsLogger.Printf("%s%s :\nā¢\t { ā¢ : %v ā¢ }\t\tT.H. %s\n", p.logger.prefix, p.name, fmt.Sprintf("%-15s", err.Error()), nbEmojiMap[5])
	os.Exit(1)
}

// PrintValue print the variable as name-value pair as the passed to the parametre.
// The message contains ā to indicate it's non-zero value and ā to indicate it is.
func (p Printer) PrintValue(varName string, v interface{}) {
	if IsNil(v) {
		x := fmt.Sprintf("%s%s :\nā\t{ š %s: %v š }\t", p.logger.prefix, p.name, fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15v", v))
		noFlagsLogger.Printf("%-50v T.H. %s\n", x, nbEmojiMap[1])
		return
	}
	x := fmt.Sprintf("%s%s :\nā\t{ š %s: %v š }\t", p.logger.prefix, p.name, fmt.Sprintf("%-15s", varName), fmt.Sprintf("%-15v", v))
	noFlagsLogger.Printf("%-50v T.H. %s\n", x, nbEmojiMap[0])
}

func IsNil(v interface{}) bool {
	if v == nil || v == 0 || v == "" {
		return true
	}
	return false
}

var nbEmojiMap = map[int]string{
	0:  "0ļøā£",
	1:  "1ļøā£",
	2:  "2ļøā£",
	3:  "3ļøā£",
	4:  "4ļøā£",
	5:  "5ļøā£",
	6:  "6ļøā£",
	7:  "7ļøā£",
	8:  "8ļøā£",
	9:  "9ļøā£",
	10: "š",
}

type KEYSType interface {
	~string | ~int | ~float64
}

type VALUESType interface {
	KEYSType | ~[]interface{} | interface{}
}
