package cqcode

func At(qq int) string {
	data := map[string]any{
		"qq": qq,
	}
	cq := CQCode{Type: "at", Data: data}
	return cq.toString()
}

func Reply(msgId int) string {
	data := map[string]any{
		"id": msgId,
	}
	cq := CQCode{Type: "reply", Data: data}
	return cq.toString()
}

func Poke(qq int) string {
	data := map[string]any{
		"qq": qq,
	}
	cq := CQCode{Type: "poke", Data: data}
	return cq.toString()
}