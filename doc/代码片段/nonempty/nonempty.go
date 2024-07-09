package nonempty

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	nonempty([]string{"1", "1", "1", "1", "", "1", "1", "1"})
}
