// The file is different as it is a helper and it should be in a seperate file
// For clean coding principals
package handlers

// Adding basic API Key here as a global variable
// Could have been done via api_key table as well
const projectAPIKey string = "3251744f-125c-458b-b80d-6f623d2a34bc"

func authentication(apiKey string) int {
	if apiKey == projectAPIKey {
		return 200
	}

	return 403
}
