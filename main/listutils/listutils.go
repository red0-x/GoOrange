package listutils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ReadList(path string) ([]map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var cols []string
	var rows []map[string]string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if cols == nil {
			cols = splitLine(line)
			continue
		}
		fields := splitLine(line)
		m := make(map[string]string)
		for i, col := range cols {
			var val string
			if i < len(fields) {
				val = fields[i]
			} else {
				val = ""
			}
			m[col] = val
		}
		rows = append(rows, m)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, errors.New("no header found in list file")
	}
	return rows, nil
}

func splitLine(s string) []string {
	parts := strings.Split(s, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
