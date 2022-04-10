package restapi

func CheckIfFileShouldBeHere(nodeName string, nodes []string) bool {
	for _, node := range nodes {
		if nodeName == node {
			return true
		}
	}
	return false
}
