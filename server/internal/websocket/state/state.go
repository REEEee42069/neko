package state

type State struct {
	banned map[string]string // IP -> session ID (that banned it)
	locked map[string]string // resource name -> session ID (that locked it)

	fileTransferEnabled       bool   // admins can transfer files
	fileTransferUnprivEnabled bool   // all users can transfer files
	fileTransferPath          string // path where files are located
}

func New(fileTransferEnabled bool, fileTransferUnprivEnabled bool, fileTransferPath string) *State {
	return &State{
		banned: make(map[string]string),
		locked: make(map[string]string),

		fileTransferEnabled:       fileTransferEnabled,
		fileTransferUnprivEnabled: fileTransferUnprivEnabled,
		fileTransferPath:          fileTransferPath,
	}
}

// Ban

func (s *State) Ban(ip, id string) {
	s.banned[ip] = id
}

func (s *State) Unban(ip string) {
	delete(s.banned, ip)
}

func (s *State) IsBanned(ip string) bool {
	_, ok := s.banned[ip]
	return ok
}

func (s *State) GetBanned(ip string) (string, bool) {
	id, ok := s.banned[ip]
	return id, ok
}

func (s *State) AllBanned() map[string]string {
	return s.banned
}

// Lock

func (s *State) Lock(resource, id string) {
	s.locked[resource] = id
}

func (s *State) Unlock(resource string) {
	delete(s.locked, resource)
}

func (s *State) IsLocked(resource string) bool {
	_, ok := s.locked[resource]
	return ok
}

func (s *State) GetLocked(resource string) (string, bool) {
	id, ok := s.locked[resource]
	return id, ok
}

func (s *State) AllLocked() map[string]string {
	return s.locked
}

// File Transfer

func (s *State) FileTransferEnabled() bool {
	return s.fileTransferEnabled
}

func (s *State) UnprivFileTransferEnabled() bool {
	return s.fileTransferUnprivEnabled
}

func (s *State) FileTransferPath() string {
	return s.fileTransferPath
}
