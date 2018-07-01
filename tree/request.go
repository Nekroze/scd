package tree

type DirectoryRequest struct {
	RawPath   string // The exact user input.
	Directory string // The final element of the RawPath.
	Path      string // Everthing except the final element of RawPath.
	Schema    string // If RawPath is a URI such as https://github.com/Nekroze/scd this would be https
}

func NewDirectoryRequest(userInput string) (new DirectoryRequest) {
	new.RawPath = userInput
	panic("Not Implemented")
	return new
}
