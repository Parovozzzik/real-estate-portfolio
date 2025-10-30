package utils

import "time"

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) > 0 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}
