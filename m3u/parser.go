package m3u

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var tvgNameRegex = regexp.MustCompile("tvg-name=\"(.*?)\"")
var tvgIDRegex = regexp.MustCompile("tvg-id=\"(.*?)\"")
var logoRegex = regexp.MustCompile("tvg-logo=\"(.*?)\"")
var groupTitleRegex = regexp.MustCompile("group-title=\"(.*?)\"")
var titleRegex = regexp.MustCompile(`[,](.*?)$`)

type M3uParser struct {
	logger *log.Logger
}

func NewParser() *M3uParser {
	return &M3uParser{
		log.New(os.Stdout, "m3uParser: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (p *M3uParser) ParseM3U(filename string) ([]*M3UEntry, error) {
	if info, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("The file %s does not exist", filename)
	} else if info.Size() == 0 {
		log.Fatalf("The file %s is empty", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entries := make([]*M3UEntry, 0)
	var extinf *ExtinfInfo

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#EXTINF:") {
			extinf = p.parseExtinfLine(line)
		} else if !strings.HasPrefix(line, "#") && extinf != nil {
			entry := &M3UEntry{
				Extinf: extinf,
				URL:    line,
			}
			entries = append(entries, entry)
			extinf = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (p *M3uParser) parseExtinfLine(line string) *ExtinfInfo {
	tvgName, err := p.getByRegex(tvgNameRegex, line)
	if err != nil {
		p.logger.Printf("failed to get tvgName:\n%v", err)
		return nil
	}
	tvgID, err := p.getByRegex(tvgIDRegex, line)
	if err != nil {
		p.logger.Printf("failed to get tvgID:\n%v", err)
		return nil
	}
	tvgLogo, err := p.getByRegex(logoRegex, line)
	if err != nil {
		p.logger.Printf("failed to get tvgLogo:\n%v", err)
		return nil
	}
	tvgGroupTitle, err := p.getByRegex(groupTitleRegex, line)
	if err != nil {
		p.logger.Printf("failed to get tvgCategory:\n%v", err)
		return nil
	}
	title, err := p.getByRegex(titleRegex, line)
	if err != nil {
		p.logger.Printf("failed to get title but will continue:\n%v", err)
	}

	return &ExtinfInfo{
		tvgID,
		tvgName,
		tvgLogo,
		tvgGroupTitle,
		title,
	}
}

func (p *M3uParser) getByRegex(re *regexp.Regexp, content string) (string, error) {
	matches := re.FindStringSubmatch(content)
	if len(matches) > 0 {
		return matches[1], nil
	}
	return "", fmt.Errorf("could not find %v in string %q", re, content)
}
