// Generated by goversion v0.1.0-SNAPSHOT
package goversion

// The following variables should be filled with goversion ldflags
var (
	buildBy           string
	buildDate         string
	buildPlatformArch string
	buildPlatformOS   string
	gitCommit         string
    gitTreeState      string
	goVersion         string
	version           string
)

func init() {
	Set(Info{
		BuildBy:           buildBy,
		BuildDate:         buildDate,
		BuildPlatformArch: buildPlatformArch,
		BuildPlatformOS:   buildPlatformOS,
		GitCommit:         gitCommit,
        GitTreeState:      gitTreeState,
		GoVersion:         goVersion,
		Version:           version,
	})
}
