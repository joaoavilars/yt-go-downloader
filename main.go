package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Format struct {
	FormatID   string  `json:"format_id"`
	Ext        string  `json:"ext"`
	FormatNote string  `json:"format_note"`
	VCodec     string  `json:"vcodec"`
	ACodec     string  `json:"acodec"`
	Abr        float64 `json:"abr"`
}

type VideoInfo struct {
	Title   string   `json:"title"`
	Formats []Format `json:"formats"`
}

func getExecPath(filename string) string {
	exePath, err := os.Executable()
	if err != nil {
		return filename
	}
	dir := filepath.Dir(exePath)
	return filepath.Join(dir, filename)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: ytbaixar.exe <URL>")
		os.Exit(1)
	}
	url := os.Args[1]

	ytDlp := getExecPath("yt-dlp.exe")

	cmd := exec.Command(ytDlp, "-j", "--no-playlist", url)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("Erro ao obter informações do vídeo:")
		fmt.Println(string(stdout)) // Exibe erro que veio junto
		os.Exit(1)
	}

	var info VideoInfo
	err = json.Unmarshal(stdout, &info)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:")
		fmt.Println(string(stdout)) // Mostra o que veio do yt-dlp
		os.Exit(1)
	}

	type Option struct {
		Label    string
		FormatID string
	}
	var options []Option

	fmt.Println("\nFormatos disponíveis:")
	for _, f := range info.Formats {
		if f.VCodec != "none" && f.ACodec != "none" && f.Ext == "mp4" {
			label := fmt.Sprintf("mp4 - %s (%s)", f.FormatNote, f.FormatID)
			options = append(options, Option{Label: label, FormatID: f.FormatID})
		} else if f.VCodec == "none" && f.Ext == "m4a" {
			label := fmt.Sprintf("mp3 - %.0fkbps (%s)", f.Abr, f.FormatID)
			options = append(options, Option{Label: label, FormatID: f.FormatID})
		}
	}

	for i, opt := range options {
		fmt.Printf("%d. %s\n", i+1, opt.Label)
	}

	fmt.Print("\nDigite o número da opção desejada: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	if choice < 1 || choice > len(options) {
		fmt.Println("Opção inválida.")
		os.Exit(1)
	}

	selected := options[choice-1]
	fmt.Println("\nBaixando:", selected.Label)

	//ffmpeg := getExecPath("ffmpeg.exe")
	env := append(os.Environ(), fmt.Sprintf("PATH=%s;%s", filepath.Dir(ytDlp), os.Getenv("PATH")))

	var downloadCmd *exec.Cmd
	if strings.HasPrefix(selected.Label, "mp3") {
		downloadCmd = exec.Command(ytDlp, "-f", selected.FormatID, "--extract-audio", "--audio-format", "mp3", "--audio-quality", "0", url)
	} else {
		downloadCmd = exec.Command(ytDlp, "-f", selected.FormatID, url)
	}

	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	downloadCmd.Env = env
	err = downloadCmd.Run()
	if err != nil {
		fmt.Println("Erro ao baixar o vídeo:", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Download concluído.")
}
