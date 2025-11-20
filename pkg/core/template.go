package core

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var appCodeTemplates = []string{
	"android/ui/theme/Type.kt.tmpl",
	"android/ui/theme/Color.kt.tmpl",
	"android/ui/theme/Theme.kt.tmpl",
	"android/CustomWebViewClient.kt.tmpl",
	"android/provider/OkHttpProvider.kt.tmpl",
	"android/screens/WebViewScreen.kt.tmpl",
	"android/MainActivity.kt.tmpl",
	"android/components/InternalWebView.kt.tmpl",
	"android/viewmodels/WebViewModel.kt.tmpl",
}

var projectFileTemplates = []string{
	"settings.gradle.kts.tmpl",
	".idea/.name.tmpl",
	"app/src/main/res/values/strings.xml.tmpl",
	"app/build.gradle.kts.tmpl",
}

func RenderTemplate(config ProjectConfiguration, source fs.FS, appCodeFolder fs.FS) error {
	templateConfig, err := config.ToTemplateConfig()
	if err != nil {
		return err
	}

	tempDir, err := os.MkdirTemp("", "appify-template-*")
	if err != nil {
		return err
	}

	err = os.CopyFS(tempDir, source)

	if err != nil {
		return err
	}

	packagePath := strings.ReplaceAll(config.PackageName, ".", "/")
	fullPackagePath := filepath.Join(tempDir, "template/app/src/main/java", packagePath)

	for _, projectFile := range projectFileTemplates {
		templatefilePath := filepath.Join(tempDir, "template", projectFile)
		outfilePath := filepath.Join(tempDir, "template", strings.Replace(projectFile, ".tmpl", "", 1))

		ExecuteTemplateWithCleanup(templatefilePath, outfilePath, *templateConfig)
	}

	err = os.CopyFS(tempDir, appCodeFolder)
	if err != nil {
		return err
	}

	for _, codeFile := range appCodeTemplates {
		codefilePath := filepath.Join(tempDir, "app", codeFile)
		outfilePath := filepath.Join(tempDir, "app", strings.Replace(codeFile, ".tmpl", "", 1))

		ExecuteTemplateWithCleanup(codefilePath, outfilePath, *templateConfig)
	}

	err = os.MkdirAll(filepath.Dir(fullPackagePath), 0777)
	if err != nil {
		return err
	}

	err = os.Mkdir(filepath.Join(tempDir, "template/app/src/main/assets"), 0777)
	if err != nil {
		return err
	}

	scriptFiles, err := os.ReadDir(config.OnloadScripts)
	if err != nil {
		return err
	}

	for _, scriptFile := range scriptFiles {
		ext := filepath.Ext(scriptFile.Name())
		if ext != "js" {
			continue
		}

		err = CopyFile(
			filepath.Join(config.OnloadScripts, scriptFile.Name()),
			filepath.Join(tempDir, "template/app/src/main/assets", scriptFile.Name()),
		)
		if err != nil {
			return err
		}
	}

	err = os.Rename(filepath.Join(tempDir, "app/android"), fullPackagePath)
	if err != nil {
		return err
	}

	os.Rename(filepath.Join(tempDir, "template"), config.ProjectDirectory)
	os.RemoveAll(tempDir)

	return nil
}

func CopyFile(inputfilePath string, outfilePath string) error {
	inputFile, err := os.Open(inputfilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outfile, err := os.Create(outfilePath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, inputFile)
	return err
}

func ExecuteTemplateWithCleanup(inputfilePath string, outfilePath string, templateConfig TemplateProjectConfiguration) error {
	outfile, err := os.Create(outfilePath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	templatefileContent, err := os.ReadFile(inputfilePath)
	if err != nil {
		return err
	}

	template, err := template.New(inputfilePath).Parse(string(templatefileContent))
	if err != nil {
		return err
	}

	err = template.Execute(outfile, templateConfig)
	if err != nil {
		return err
	}

	err = os.Remove(inputfilePath)
	if err != nil {
		return err
	}

	return nil
}
