package sockets

type gameManager struct {
}

// makeGameManager -- Return a pointer to a new gameManager instance, and set it in motion
func makeGameManager() *gameManager {
	gm := gameManager{}
	return &gm
}
