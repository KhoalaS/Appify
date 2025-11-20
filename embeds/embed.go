package embeds

import "embed"

//go:embed template/*
//go:embed template/.idea/*
//go:embed template/app/*
var TemplateFolder embed.FS

//go:embed app
var AppCodeFolder embed.FS

//go:embed userscripts
//go:embed config.json
//go:embed config.schema.json
var ExampleFiles embed.FS

//go:embed env.d.ts
//go:embed package.json
//go:embed tsconfig.json
var TypescriptFiles embed.FS
