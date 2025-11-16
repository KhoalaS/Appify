package appify

import "embed"

//go:embed template/*
//go:embed template/.idea/*
//go:embed template/app/*
var TemplateFolder embed.FS

//go:embed app
var AppCodeFolder embed.FS