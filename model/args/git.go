package args

type GitArgs struct {
	SubmodulePaths []string
	TargetDir      string
	GitRoot        *string
}

func (g *GitArgs) BuildPrefixCmd() []string {
	args := []string{"git"}
	if g.GitRoot != nil {
		args = append(args, "-C", *g.GitRoot)
	}

	return args
}
