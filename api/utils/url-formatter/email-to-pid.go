package urlformatter

func EmailToPid(pid string) string {
	var email []byte
	for _, K := range pid {
		if pid[K] == '@' {
			break
		}

		email[K] = pid[K]
	}

	return string(email)
}
