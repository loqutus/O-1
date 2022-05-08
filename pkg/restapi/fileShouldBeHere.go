package restapi

// CheckIfFileShouldBeHere checks if file should be at this node
func CheckIfFileShouldBeHere(nodeName string, nodes []string) bool {
	for _, node := range nodes {
		if nodeName == node {
			return true
		}
	}
	return false
}
