package serviceaccounts

func fixClientIDAndClientSecret(items []map[string]interface{}) []map[string]interface{} {
	// Fix the client id and client secret
	answer := make([]map[string]interface{}, 0)
	for _, entry := range items {
		entry["client_id"] = entry["clientID"]
		delete(entry, "clientID")
		if entry["clientSecret"] != nil {
			entry["client_secret"] = entry["clientSecret"]
			delete(entry, "clientSecret")
		}
		answer = append(answer, entry)
	}
	return answer
}