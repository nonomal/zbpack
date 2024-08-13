package nodejs_test

import (
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/zeabur/zbpack/internal/nodejs"
)

func TestMain(m *testing.M) {
	v := m.Run()

	// After all tests have run `go-snaps` will sort snapshots
	snaps.Clean(m, snaps.CleanOpts{Sort: true})

	os.Exit(v)
}

func TestTemplate_NBuildCmd_NOutputDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "RUN yarn install",
		BuildCmd:   "",
		StartCmd:   "yarn start",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_NBuildCmd_OutputDir_NSPA(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "RUN yarn install",
		BuildCmd:   "",
		StartCmd:   "yarn start",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_NBuildCmd_OutputDir_SPA(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "RUN yarn install",
		BuildCmd:   "",
		StartCmd:   "yarn start",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_BuildCmd_NOutputDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "RUN yarn install",
		BuildCmd:   "yarn build",
		StartCmd:   "yarn start",

		Serverless: true,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_BuildCmd_OutputDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",

		InstallCmd: "RUN yarn install",
		BuildCmd:   "yarn build",
		StartCmd:   "yarn start",

		OutputDir: "/app/dist",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_BuildCmd_Bun(t *testing.T) {
	ctx := nodejs.TemplateContext{
		Bun:         true,
		NodeVersion: "18",
		InstallCmd:  "RUN bun install",
		StartCmd:    "bun start main.ts",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_Monorepo(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",
		AppDir:      "myservice",
		InstallCmd:  "WORKDIR /src/myservice\nRUN yarn install",
		StartCmd:    "yarn start",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_MonorepoServerless(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",
		AppDir:      "myservice",
		InstallCmd:  "WORKDIR /src/myservice\nRUN yarn install",
		StartCmd:    "yarn start",
		Serverless:  true,
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_MonorepoServerlessOutDir(t *testing.T) {
	ctx := nodejs.TemplateContext{
		NodeVersion: "18",
		AppDir:      "myservice",
		InstallCmd:  "WORKDIR /src/myservice\nRUN yarn install",
		StartCmd:    "yarn start",
		OutputDir:   "/app/dist",
	}

	result, err := ctx.Execute()
	assert.NoError(t, err)
	snaps.MatchSnapshot(t, result)
}

func TestTemplate_NitroPreset(t *testing.T) {
	t.Parallel()

	nitroBasedFrameworks := []string{
		"nuxt.js",
		"nitropack",
	}

	for _, framework := range nitroBasedFrameworks {
		t.Run(framework, func(t *testing.T) {
			t.Parallel()

			t.Run("node-server", func(t *testing.T) {
				t.Parallel()

				ctx := nodejs.TemplateContext{
					NodeVersion: "18",
					InstallCmd:  "RUN yarn install",
					StartCmd:    "node .output/server/index.mjs",
					Framework:   framework,
					Serverless:  false,
				}

				result, err := ctx.Execute()
				assert.NoError(t, err)
				assert.Contains(t, result, "ENV NITRO_PRESET=node-server")
			})

			t.Run("bun", func(t *testing.T) {
				t.Parallel()

				ctx := nodejs.TemplateContext{
					NodeVersion: "18",
					InstallCmd:  "RUN bun install",
					StartCmd:    "bun .output/server/index.mjs",
					Framework:   framework,
					Serverless:  false,
					Bun:         true,
				}

				result, err := ctx.Execute()
				assert.NoError(t, err)
				assert.Contains(t, result, "ENV NITRO_PRESET=bun")
			})

			t.Run("node", func(t *testing.T) {
				t.Parallel()

				ctx := nodejs.TemplateContext{
					NodeVersion: "18",
					InstallCmd:  "RUN yarn install",
					StartCmd:    "",
					Framework:   framework,
					Serverless:  true,
				}

				result, err := ctx.Execute()
				assert.NoError(t, err)
				assert.Contains(t, result, "ENV NITRO_PRESET=node")
			})
		})
	}

	t.Run("empty if not nitro-based framework", func(t *testing.T) {
		t.Parallel()

		ctx := nodejs.TemplateContext{
			NodeVersion: "18",
			InstallCmd:  "RUN yarn install",
			StartCmd:    "yarn start",
			Framework:   "next.js",
			Serverless:  false,
		}

		result, err := ctx.Execute()
		assert.NoError(t, err)
		assert.NotContains(t, result, "ENV NITRO_PRESET")
	})
}
